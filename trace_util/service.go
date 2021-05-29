package trace_util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

type Service struct {
	conf   Config
	tracer opentracing.Tracer
}

func GetService(conf Config) *Service {
	return &Service{conf: conf}
}

func (d *Service) GetTracer() opentracing.Tracer {
	return d.tracer
}

func (d *Service) Init() {
	cfg := jaegercfg.Configuration{
		ServiceName: d.conf.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1, // 1-全采样 0-不采样
		},
		Reporter: &jaegercfg.ReporterConfig{
			LocalAgentHostPort: fmt.Sprintf("%s:%s", d.conf.AgentHost, d.conf.AgentPort),
			LogSpans:           true,
		},
	}

	var err error

	jLogger := jaegerlog.NullLogger
	jMetricsFactory := metrics.NullFactory
	d.tracer, _, err = cfg.NewTracer(jaegercfg.Logger(jLogger), jaegercfg.Metrics(jMetricsFactory))

	if err != nil {
		fmt.Println("create tracer failed!", err)
	}

	opentracing.SetGlobalTracer(d.tracer)
}

func (d *Service) Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		if d.conf.Enable && d.tracer != nil {
			var span opentracing.Span

			spanC, _ := d.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if spanC == nil {
				// root span
				span = d.tracer.StartSpan(c.Request.URL.RequestURI())
			} else {
				span = d.tracer.StartSpan(c.Request.URL.RequestURI(), opentracing.ChildOf(spanC))
			}

			_ = d.tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			span.SetTag("method", c.Request.Method)

			c.Next()

			span.Finish()
		}
	}
}
