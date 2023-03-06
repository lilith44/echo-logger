package echo_logger

import (
	"encoding/json"
	"io"
	"unsafe"

	"github.com/labstack/gommon/log"
	"github.com/mattn/go-colorable"
	"github.com/valyala/fasttemplate"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger   *zap.SugaredLogger
	output   io.Writer
	prefix   string
	level    log.Lvl
	template *fasttemplate.Template
}

func NewZapLogger(logger *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{
		logger: logger,
		output: colorable.NewColorableStdout(),
		level:  log.INFO,
	}
}

func (zl *ZapLogger) Output() io.Writer {
	return zl.output
}

func (zl *ZapLogger) SetOutput(w io.Writer) {
	zl.output = w
}

func (zl *ZapLogger) Prefix() string {
	return zl.prefix
}

func (zl *ZapLogger) SetPrefix(p string) {
	zl.prefix = p
}

func (zl *ZapLogger) Level() log.Lvl {
	return zl.level
}

func (zl *ZapLogger) SetLevel(v log.Lvl) {
	zl.level = v
}

func (zl *ZapLogger) SetHeader(h string) {
	zl.template = fasttemplate.New(h, "${", "}")
}

func (zl *ZapLogger) Print(i ...any) {
	zl.logger.Info(i...)
}

func (zl *ZapLogger) Printf(format string, args ...any) {
	zl.logger.Infof(format, args...)
}

func (zl *ZapLogger) Printj(j log.JSON) {
	data, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	zl.logger.Info(*(*string)(unsafe.Pointer(&data)))
}

func (zl *ZapLogger) Debug(i ...any) {
	if zl.level <= log.DEBUG {
		zl.logger.Debug(i...)
	}
}

func (zl *ZapLogger) Debugf(format string, args ...any) {
	if zl.level <= log.DEBUG {
		zl.logger.Debugf(format, args...)
	}
}

func (zl *ZapLogger) Debugj(j log.JSON) {
	if zl.level <= log.DEBUG {
		data, err := json.Marshal(j)
		if err != nil {
			panic(err)
		}
		zl.logger.Debug(*(*string)(unsafe.Pointer(&data)))
	}
}

func (zl *ZapLogger) Info(i ...any) {
	if zl.level <= log.INFO {
		zl.logger.Info(i...)
	}
}

func (zl *ZapLogger) Infof(format string, args ...any) {
	if zl.level <= log.INFO {
		zl.logger.Infof(format, args...)
	}
}

func (zl *ZapLogger) Infoj(j log.JSON) {
	if zl.level <= log.INFO {
		data, err := json.Marshal(j)
		if err != nil {
			panic(err)
		}
		zl.logger.Info(*(*string)(unsafe.Pointer(&data)))
	}
}

func (zl *ZapLogger) Warn(i ...any) {
	if zl.level <= log.WARN {
		zl.logger.Warn(i...)
	}
}

func (zl *ZapLogger) Warnf(format string, args ...any) {
	if zl.level <= log.WARN {
		zl.logger.Warnf(format, args...)
	}
}

func (zl *ZapLogger) Warnj(j log.JSON) {
	if zl.level <= log.WARN {
		data, err := json.Marshal(j)
		if err != nil {
			panic(err)
		}
		zl.logger.Warn(*(*string)(unsafe.Pointer(&data)))
	}
}

func (zl *ZapLogger) Error(i ...any) {
	if zl.level <= log.ERROR {
		zl.logger.Error(i...)
	}
}

func (zl *ZapLogger) Errorf(format string, args ...any) {
	if zl.level <= log.ERROR {
		zl.logger.Errorf(format, args...)
	}
}

func (zl *ZapLogger) Errorj(j log.JSON) {
	if zl.level <= log.ERROR {
		data, err := json.Marshal(j)
		if err != nil {
			panic(err)
		}
		zl.logger.Error(*(*string)(unsafe.Pointer(&data)))
	}
}

func (zl *ZapLogger) Fatal(i ...any) {
	zl.logger.Fatal(i...)
}

func (zl *ZapLogger) Fatalj(j log.JSON) {
	data, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	zl.logger.Fatal(*(*string)(unsafe.Pointer(&data)))
}

func (zl *ZapLogger) Fatalf(format string, args ...any) {
	zl.logger.Fatalf(format, args...)
}

func (zl *ZapLogger) Panic(i ...any) {
	zl.logger.Panic(i...)
}

func (zl *ZapLogger) Panicj(j log.JSON) {
	data, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	zl.logger.Panic(*(*string)(unsafe.Pointer(&data)))
}

func (zl *ZapLogger) Panicf(format string, args ...any) {
	zl.logger.Panicf(format, args...)
}
