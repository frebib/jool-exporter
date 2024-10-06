package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/exporter-toolkit/web"
)

var opts struct {
	MetricsPath   string `short:"p" long:"metrics-path" default:"/metrics" description:"http path at which to serve metrics"`
	ListenAddr    string `long:"web.listen-address" default:":9441" description:"http listen address to serve metrics"`
	WebConfigFile string `long:"web.config.file" required:"false" description:"path to web-config file"`
}

func main() {
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(log)

	parser := flags.NewParser(&opts, 0)
	pos, err := parser.Parse()
	if err != nil || len(pos) > 0 {
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(collectors.NewBuildInfoCollector())
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(NewJoolCollector(context.Background(), "jool"))

	router := http.NewServeMux()
	router.Handle(opts.MetricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><head><title>Jool Exporter</title></head><body><h1>Jool Exporter</h1><p><a href='/metrics'>Metrics</a></p></body></html>`))
	})

	flagConfig := &web.FlagConfig{
		WebListenAddresses: &[]string{opts.ListenAddr},
		WebConfigFile:      &opts.WebConfigFile,
	}
	server := &http.Server{
		Handler: router,
	}

	err = web.ListenAndServe(server, flagConfig, log)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error(err.Error())
	}
}
