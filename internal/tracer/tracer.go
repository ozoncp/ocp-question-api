package tracer

import (
	"fmt"
	"io"

	"github.com/ozoncp/ocp-question-api/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

// InitTracer - init tracer
func InitTracer(serviceName string) io.Closer {
	conf := config.NewConfig()

	metricsConfig := &jaegerConfig.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%v:%v", conf.Jagger.Host, conf.Jagger.Port),
		},
	}

	tracer, closer, err := metricsConfig.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		log.Error().Err(err).Msg("failed init jaeger")
	}

	opentracing.SetGlobalTracer(tracer)
	log.Info().Msg("Traces started")

	return closer
}
