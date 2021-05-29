package hook

import (
	"context"
	"github.com/huangxuantao/xorm-v1/contexts"
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"xorm.io/builder"
)

const ContextSpanKey = "ContextSpanKey"

type XormTracingHook struct {
	before func(c *contexts.ContextHook) (context.Context, error)
	after  func(c *contexts.ContextHook) error
}

var _ contexts.Hook = &XormTracingHook{}

func (h *XormTracingHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return h.before(c)
}

func (h *XormTracingHook) AfterProcess(c *contexts.ContextHook) error {
	return h.after(c)
}

func NewXormTracingHook() *XormTracingHook {
	return &XormTracingHook{
		before: xormBefore,
		after:  xormAfter,
	}
}

func xormBefore(c *contexts.ContextHook) (context.Context, error) {
	var span opentracing.Span
	sql, _ := builder.ConvertToBoundSQL(c.SQL, c.Args)

	if c.Ctx.Value(ContextSpanKey) != nil {
		span = opentracing.GlobalTracer().StartSpan(sql, opentracing.ChildOf(c.Ctx.Value(ContextSpanKey).(jaeger.SpanContext)))
	} else {
		span = opentracing.GlobalTracer().StartSpan(sql)
	}
	c.Ctx = context.WithValue(c.Ctx, ContextSpanKey, span)
	return c.Ctx, nil
}

func xormAfter(c *contexts.ContextHook) error {
	span, ok := c.Ctx.Value(ContextSpanKey).(opentracing.Span)
	if !ok {
		return nil
	}
	defer span.Finish()

	if c.Err != nil {
		span.LogFields(tracerLog.Object("err", c.Err))
	}

	if len(c.Args) != 0 {
		span.LogFields(tracerLog.Object("args", c.Args))
	}
	span.SetTag("execute time", c.ExecuteTime)

	return nil
}
