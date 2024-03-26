module sbg_client

go 1.22.1

replace sbg_server => ../sbg_server

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/sirupsen/logrus v1.9.3
	sbg_server v0.0.0-00010101000000-000000000000
)

require (
	github.com/alexflint/go-scalar v1.1.0 // indirect
	github.com/fabienm/go-logrus-formatters v1.0.0 // indirect
	github.com/gemnasium/logrus-graylog-hook/v3 v3.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
