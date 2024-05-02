// пакеты исполняемых приложений должны называться main
package main

import (
	"github.com/Neeeooshka/alice-skill.git/internal/gzip"
	"github.com/Neeeooshka/alice-skill.git/internal/handlers"
	"github.com/Neeeooshka/alice-skill.git/internal/zap"
	"github.com/Neeeooshka/alice-skill.git/pkg/compressor"
	"log"
	"net/http"
)

func main() {
	// обрабатываем аргументы командной строки
	parseFlags()

	logger, err := zap.NewZapLogger("info")
	if err != nil {
		panic(err)
	}

	logger.Info("Running server", logger.String("address", flagRunAddr))
	// оборачиваем хендлер в middleware с логированием и поддержкой gzip
	log.Fatal(http.ListenAndServe(flagRunAddr, zap.RequestLogger(compressor.IncludeCompressor(handlers.AliceSkill, gzip.NewGzipCompressor()))))
}
