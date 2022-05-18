package main 

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger provides contextual formated json logs for consumption by humans
// and machines. It is powered by zerolog.
type Logger struct {
	log zerolog.Logger
}

// logWriter uses the APP_ENV value to determine the io.Writer to use with
// zerolog.ConsoleWriter. By default it writes logs to os.Stdout, but will use
// io.Discard (/dev/null effectively) if env is set to "test".
func logWriter(env string) zerolog.ConsoleWriter {
	var output io.Writer = os.Stdout
	if env == "test" {
		output = io.Discard
	}

	return zerolog.ConsoleWriter{Out: output, TimeFormat: time.RFC3339Nano}
}

// NewLogger returns a new pointer to Logger.
func NewLogger(conf Config) *Logger {
	return &Logger{
		log: zerolog.New(logWriter(conf.AppEnv)).
			With().Timestamp().
			Str("APP_NAME", conf.AppName).
			Str("APP_ENV", conf.AppEnv).
			Logger(),
	}
}

// And adds another string key/value pair to the context of the logger.
func (logger *Logger) And(key, value string) *Logger {
	logger.log = logger.log.With().Str(key, value).Logger()
	return logger
}

// Dump complex values to the log under the given key.
func (logger *Logger) Dump(key string, value interface{}) *Logger {
	logger.log = logger.log.With().Interface(key, value).Logger()
	return logger
}

// Info writes to the error INFO log. This function can be chained.
func (logger *Logger) Info(msg string) *Logger {
	logger.log.Info().Msg(msg)
	return logger
}

// Debug writes to the error DBG log. This function can be chained.
func (logger *Logger) Debug(msg string) *Logger {
	logger.log.Debug().Msg(msg)
	return logger
}

// Error writes to the error log when err is not nil. It also returns the
// error, so it can be chained with other error handling code.
func (logger *Logger) Error(err error, msg string) error {
	if err != nil {
		logger.log.Error().Err(err).Msg(msg)
	}

	return err
}

// Fatal writes to the error log when err is not nil, and then exits.
func (logger *Logger) Fatal(err error, msg string) error {
	if err != nil {
		logger.log.Fatal().Err(err).Msg(msg)
	}

	return err
}
