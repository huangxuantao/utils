package log_util

import (
	"github.com/kjk/dailyrotate"
	"github.com/op/go-logging"
	"io"
	"os"
)

var (
	ApiLogger *logging.Logger
	Logger    *logging.Logger
	logLevel  = map[string]logging.Level{
		"DEBUG":    logging.DEBUG,
		"INFO":     logging.INFO,
		"WARNING":  logging.WARNING,
		"ERROR":    logging.ERROR,
		"CRITICAL": logging.CRITICAL,
		"NOTICE":   logging.NOTICE,
	}
)

type LoggerConfig struct {
	ErrorLogPath  string
	AccessLogPath string
	Level         string
	DevMode       bool
}

func Init(cfg LoggerConfig) {
	ApiLogger = cfg.getLogger("access")
	Logger = cfg.getLogger("error")
}

func (cfg *LoggerConfig) getLevelBackend(logFile io.Writer, level string, format logging.Formatter) logging.LeveledBackend {
	logBackend := logging.NewLogBackend(logFile, "", 0)
	levelBackend := logging.AddModuleLevel(logging.NewBackendFormatter(logBackend, format))
	levelBackend.SetLevel(logLevel[level], "")
	return levelBackend
}

func (cfg *LoggerConfig) getLoggerBackend(logPath, level string, devMode bool) logging.LeveledBackend {
	loggerFmtOut := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05} %{shortfile} %{shortfunc} [%{level}%{color:reset}] > %{message}`)
	loggerFmtFile := logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05} %{shortfile} %{shortfunc} [%{level:.4s}] > %{message}`)
	logF, err := dailyrotate.NewFile(logPath, nil)

	loggerBackend := logging.SetBackend()
	fileLevelBackend := cfg.getLevelBackend(logF, level, loggerFmtFile)
	stdErrLevelBackend := cfg.getLevelBackend(os.Stderr, level, loggerFmtOut)

	if err == nil && devMode {
		loggerBackend = logging.SetBackend(fileLevelBackend, stdErrLevelBackend)
	} else if err == nil && !devMode {
		loggerBackend = logging.SetBackend(fileLevelBackend)
	} else {
		loggerBackend = logging.SetBackend(stdErrLevelBackend)
	}
	return loggerBackend
}

func (cfg *LoggerConfig) getLogger(logger string) *logging.Logger {
	switch logger {
	case "access":
		var logger = logging.MustGetLogger("access")
		loggerBackend := cfg.getLoggerBackend(cfg.AccessLogPath, cfg.Level, cfg.DevMode)
		logger.SetBackend(loggerBackend)
		return logger
	default:
		var logger = logging.MustGetLogger("error")
		loggerBackend := cfg.getLoggerBackend(cfg.AccessLogPath, cfg.Level, cfg.DevMode)
		logger.SetBackend(loggerBackend)
		return logger
	}
}
