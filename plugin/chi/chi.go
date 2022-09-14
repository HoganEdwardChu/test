package chi

import (
	"net/http"

	pinpoint "github.com/pinpoint-apm/pinpoint-go-agent"
	phttp "github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !pinpoint.GetAgent().Enable() {
				next.ServeHTTP(w, r)
				return
			}

			tracer := phttp.NewHttpServerTracer(r, "Chi Server")
			defer tracer.EndSpan()

			routePath := r.URL.Path
			if r.URL.RawPath != "" {
				routePath = r.URL.RawPath
			}
			defer tracer.NewSpanEvent(routePath).EndSpanEvent()

			status := http.StatusOK
			w = phttp.WrapResponseWriter(w, &status)
			r = pinpoint.RequestWithTracerContext(r, tracer)

			next.ServeHTTP(w, r)
			phttp.RecordHttpServerResponse(tracer, status, w.Header())
		}
		return http.HandlerFunc(fn)
	}
}
