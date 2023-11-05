package log

import (
	"flag"
	"os"
	"sync"

	"github.com/go-logr/logr"
	klogv2 "k8s.io/klog/v2"
	klogr "k8s.io/klog/v2/klogr"
)

var (
	_globalMu sync.RWMutex
	_globalL  *Log = nil
)

type Log struct {
	logr logr.Logger
}

func L() logr.Logger {
	_globalMu.Lock()
	defer _globalMu.Unlock()
	if _globalL != nil {
		return _globalL.logr
	}

	_globalL = &Log{logr: InitLog(getLogCfgFromEnv())}
	return _globalL.logr
}

func InTestSetLogger(logger logr.Logger) {
	_globalMu.Lock()
	defer _globalMu.Unlock()
	_globalL = &Log{logr: logger}
}

func Flush() {
	klogv2.Flush()
}

type LogCfg struct {
	ToFile string
	Level  string
}

func getLogCfgFromEnv() LogCfg {
	cfg := LogCfg{
		ToFile: "",
		Level:  "",
	}

	if os.Getenv("LOG_EXT") == "true" {
		logFile := os.Getenv("LOG_FILE")
		cfg.ToFile = logFile
		logLevel := os.Getenv("LOG_LEVEL")
		if logLevel != "" {
			cfg.Level = logLevel
		}
	}
	return cfg
}

func InitLog(cfg LogCfg) logr.Logger {
	flags := flag.NewFlagSet("klog", flag.ExitOnError)
	klogv2.InitFlags(flags)
	if cfg.ToFile != "" {
		flags.Set("log_file", cfg.ToFile)
		flags.Set("alsologtostderr", "true")
		flags.Set("logtostderr", "false")
	}
	if cfg.Level != "" {
		flags.Set("v", cfg.Level)
	}
	defer Flush()
	logger := klogr.NewWithOptions(klogr.WithFormat("Klog"))
	return logger
}
