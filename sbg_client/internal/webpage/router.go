package webpage

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sbg_server/pkg/logtools"
)

type Router interface {
	Run() error
}

type RouterCfg struct {
	ServerAddress    string `arg:"--server-address,env:SERVER_ADDRESS" default:":8080" help:"Address and port of this service"`
	WebStaticPath    string `arg:"--web-template-path,env:WEB_TEMPLATE_PATH" default:"./web/static" help:"Path to directory with template pages"`
	WebIndexFilename string `arg:"--web-index-filename,env:WEB_INDEX_FILENAME" default:"index.html" help:"Name of index file"`
}

type routerImpl struct {
	cfg RouterCfg
}

func NewRouter(cfg RouterCfg) (Router, error) {
	// Check directory existence
	if _, err := os.Stat(cfg.WebStaticPath); err != nil {
		return nil, logtools.WithStackErrorf("bad web template path: %w", err)
	}
	res := &routerImpl{
		cfg: cfg,
	}
	return res, nil
}

var _ Router = (*routerImpl)(nil)

func (rt *routerImpl) Run() error {
	mux := http.NewServeMux()

	// fs := http.FileServer(http.Dir(rt.cfg.WebTemplatePath))
	// fs := http.FileServer(http.Dir("../web/static"))
	// mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// mux.Handle("/", fs)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		var path string
		if r.URL.Path == "/" {
			path = filepath.Join(rt.cfg.WebStaticPath, rt.cfg.WebIndexFilename)
		} else {
			path = filepath.Join(rt.cfg.WebStaticPath, r.URL.Path)
		}
		http.ServeFile(w, r, path)
	})

	err := http.ListenAndServe(rt.cfg.ServerAddress, mux)
	if err != nil {
		return logtools.WithStackErrorf("failed to start server: %w", err)
	}

	return nil
}
