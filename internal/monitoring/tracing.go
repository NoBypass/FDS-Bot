package monitoring

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/nobypass/fds-bot/internal/pkg/version"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"os"
)

func CreateTracer() (opentracing.Tracer, io.Closer) {
	endpoint := os.Getenv("JAEGER_ENDPOINT")
	cfg := config.Configuration{
		ServiceName: fmt.Sprintf("FDS Discord Bot %s", version.VERSION),
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: endpoint,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatal(err)
	}
	return tracer, closer
}
