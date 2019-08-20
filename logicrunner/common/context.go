package common

import (
	"context"

	"go.opencensus.io/trace"

	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/instracer"
	"github.com/insolar/insolar/log"
)

func FreshContextFromContextAndRequest(ctx context.Context, req record.IncomingRequest) context.Context {
	res := inslogger.ContextWithTrace(
		context.Background(),
		req.APIRequestID, // this is HACK based on awareness, we just know how trace id is formed
	)
	// FIXME: need way to get level out of context
	// res = inslogger.WithLoggerLevel(res, data.LogLevel)
	parentSpan, ok := instracer.ParentSpan(ctx)
	if ok {
		res = instracer.WithParentSpan(res, parentSpan)
	}
	if pctx := trace.FromContext(ctx); pctx != nil {
		res = trace.NewContext(res, pctx)
	}
	return res
}

func ContextWithServiceData(ctx context.Context, data payload.ServiceData) context.Context {
	ctx = inslogger.ContextWithTrace(ctx, data.LogTraceID)
	ctx = inslogger.WithLoggerLevel(ctx, data.LogLevel)
	if data.TraceSpanData != nil {
		parentSpan := instracer.MustDeserialize(data.TraceSpanData)
		return instracer.WithParentSpan(ctx, parentSpan)
	}
	return ctx
}

func FreshContextFromContext(ctx context.Context) context.Context {
	res := inslogger.ContextWithTrace(
		context.Background(),
		inslogger.TraceID(ctx),
	)
	parentSpan, ok := instracer.ParentSpan(ctx)
	if ok {
		res = instracer.WithParentSpan(res, parentSpan)
	}

	if pctx := trace.FromContext(ctx); pctx != nil {
		res = trace.NewContext(res, pctx)
	}

	return res
}

func ServiceDataFromContext(ctx context.Context) *payload.ServiceData {
	if ctx == nil {
		log.Error("nil context, can't create correct ServiceData")
		return &payload.ServiceData{}
	}
	return &payload.ServiceData{
		LogTraceID:    inslogger.TraceID(ctx),
		LogLevel:      inslogger.GetLoggerLevel(ctx),
		TraceSpanData: instracer.MustSerialize(ctx),
	}
}
