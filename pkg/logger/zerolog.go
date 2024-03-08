package logger

import (
	"article-3-how-use-zerolog/config"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog zerolog.Logger
}

func NewLoggerFromConfig(lcfg config.LoggerConfig) (*Logger, error) {
	level, err := zerolog.ParseLevel(lcfg.Level)
	if err != nil {
		return nil, err
	}
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	zerolog := zerolog.New(output).Level(level).With().Timestamp().Logger()
	log := &Logger{zerolog}
	if lcfg.Development {
		log.SetDevelopmentContext()
	}
	return log, nil
}

func (l *Logger) SetDevelopmentContext() {
	l.zerolog.With().Caller().Stack()
}

// Println sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) {
	l.zerolog.Println(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.zerolog.Printf(format, v...)
}

func (l *Logger) Trace(msg string) {
	l.zerolog.Trace().Msg(msg)
}
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.zerolog.Trace().Msgf(format, v...)
}

func (l *Logger) Debug(msg string) {
	l.zerolog.Debug().Msg(msg)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.zerolog.Debug().Msgf(format, v...)
}

func (l *Logger) Info(msg string) {
	l.zerolog.Info().Msg(msg)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.zerolog.Info().Msgf(format, v...)
}

func (l *Logger) Warn(msg string) {
	l.zerolog.Warn().Msg(msg)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.zerolog.Warn().Msgf(format, v...)
}

func (l *Logger) Error(err error) {
	l.zerolog.Error().Err(err).Send()
}

func (l *Logger) Fatal(err error) {
	l.zerolog.Fatal().Err(err).Send()
}
