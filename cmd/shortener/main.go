package main

import (
	"flag"
	"github.com/Neeeooshka/alice-skill.git/internal/compress"
	"github.com/Neeeooshka/alice-skill.git/internal/config"
	"github.com/Neeeooshka/alice-skill.git/internal/handlers"
	"github.com/Neeeooshka/alice-skill.git/internal/logger"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	opt, cfg := getOptions()

	logrusLogger := newLogrusLogger(logrus.InfoLevel)

	sh := newShortener(opt, cfg)
	compressor := newGzipCompressor()

	// create router
	router := chi.NewRouter()
	router.Post("/", logger.IncludeLogger(compress.IncludeCompressor(handlers.GetShortenerHandler(&sh), compressor), logrusLogger))
	router.Get("/{id}", logger.IncludeLogger(compress.IncludeCompressor(handlers.GetExpanderHandler(&sh), compressor), logrusLogger))
	router.Post("/api/shorten", logger.IncludeLogger(compress.IncludeCompressor(handlers.GetAPIShortenHandler(&sh), compressor), logrusLogger))

	// create HTTP Server
	server := opt.GetServer()
	if cfg.ServerAddress != "" {
		server = cfg.ServerAddress
	}
	http.ListenAndServe(server, router)
}

// init options && config
func getOptions() (config.Options, config.Config) {
	opt := config.NewOptions()
	cfg := config.NewConfig()

	flag.Var(&opt.ServerAddress, "a", "Server address - host:port")
	flag.Var(&opt.BaseURL, "b", "Server ShortLink Base address - protocol://host:port")

	flag.Parse()
	env.Parse(&cfg)
	return opt, cfg
}
