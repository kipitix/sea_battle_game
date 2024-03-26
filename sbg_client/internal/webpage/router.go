package webpage

import (
	"net/http"
	"os"
	"path/filepath"
	"sbg_server/pkg/logtools"
	"text/template"
)

type Router interface {
	Run() error
}

type RouterCfg struct {
	ServerAddress   string `arg:"--server-address,env:SERVER_ADDRESS" default:":8080" help:"Address and port of this service"`
	WebTemplatePath string `arg:"--web-template-path,env:WEB_TEMPLATE_PATH" default:"./web/template" help:"Path to directory with template pages"`
}

type routerImpl struct {
	cfg RouterCfg
}

func NewRouter(cfg RouterCfg) (Router, error) {
	// Check directory existence
	if _, err := os.Stat(cfg.WebTemplatePath); err != nil {
		return nil, logtools.WithStackErrorf("bad web template path: %w", err)
	}

	return &routerImpl{
		cfg: cfg,
	}, nil
}

var _ Router = (*routerImpl)(nil)

func (rt *routerImpl) Run() error {
	mux := http.NewServeMux()

	// fs := http.FileServer(http.Dir(rt.cfg.WebTemplatePath))
	fs := http.FileServer(http.Dir("../web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// mux.Handle("/", fs)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lp := filepath.Join("../web/template", "layout.html")
		fp := filepath.Join("../web/template", filepath.Clean(r.URL.Path))

		page, err := template.ParseFiles(lp, fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		page.ExecuteTemplate(w, "layout", nil)
	})

	err := http.ListenAndServe(rt.cfg.ServerAddress, mux)
	if err != nil {
		return logtools.WithStackErrorf("failed to start server: %w", err)
	}

	return nil
}
