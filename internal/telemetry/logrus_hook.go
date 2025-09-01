package telemetry

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
)

type OTelHook struct {
	logger log.Logger
}

func NewOTelHook() *OTelHook {
	lp := global.GetLoggerProvider()
	return &OTelHook{
		logger: lp.Logger("arkd"),
	}
}

func (h *OTelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func mapLevel(l logrus.Level) log.Severity {
	switch l {
	case logrus.PanicLevel:
		return log.SeverityFatal
	case logrus.FatalLevel:
		return log.SeverityFatal
	case logrus.ErrorLevel:
		return log.SeverityError
	case logrus.WarnLevel:
		return log.SeverityWarn
	case logrus.InfoLevel:
		return log.SeverityInfo
	case logrus.DebugLevel:
		return log.SeverityDebug
	case logrus.TraceLevel:
		return log.SeverityTrace
	default:
		return log.SeverityInfo
	}
}

func (h *OTelHook) Fire(e *logrus.Entry) error {
	var pcs [1]uintptr
	runtime.Callers(6, pcs[:])
	rec := log.Record{}
	rec.SetTimestamp(e.Time)
	rec.SetSeverity(mapLevel(e.Level))
	rec.SetBody(log.StringValue(e.Message))

	rec.AddAttributes(
		log.String("log.kind", "app"),
		log.String("logger", "ark.daemon"),
		log.String("level", e.Level.String()),
	)

	// include fields as attributes
	for k, v := range e.Data {
		rec.AddAttributes(log.String(k, toString(v)))
	}

	// observed ts (optional)
	rec.SetObservedTimestamp(time.Now())

	h.logger.Emit(context.Background(), rec)
	return nil
}

func toString(v any) string {
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprintf("%v", v)
	}
}
