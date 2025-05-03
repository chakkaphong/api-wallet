package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	log  *zap.Logger
	once sync.Once
)

// Init initializes a Zap logger (singleton-safe)
func Init() *zap.Logger {
	once.Do(func() {
		var err error
		log, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	})
	return log
}

// L returns the logger instance
func L() *zap.Logger {
	if log == nil {
		panic("Logger not initialized. Call logger.Init() first.")
	}
	return log
}
