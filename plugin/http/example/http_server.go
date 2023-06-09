package main

import (
	"context"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
	"io"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	tracer := pinpoint.FromContext(r.Context())
	defer tracer.NewSpanEvent("dummy").EndSpanEvent()

	io.WriteString(w, "hello world")
}

func wrapRequest(w http.ResponseWriter, r *http.Request) {
	ctx := pinpoint.NewContext(context.Background(), pinpoint.FromContext(r.Context()))
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:9000/hello", nil)

	resp, err := pphttp.DoClient(http.DefaultClient.Do, req)
	if nil != err {
		io.WriteString(w, err.Error())
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func wrapClient(w http.ResponseWriter, r *http.Request) {
	client := pphttp.WrapClientWithContext(r.Context(), &http.Client{})
	resp, err := client.Get("http://localhost:9000/async_wrapper?foo=bar&say=goodbye")

	if nil != err {
		w.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(w, err.Error())
		return
	}
	defer resp.Body.Close()

	w.Header().Set("foo", "bar")
	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, "wrapClient success")
}

func handlerNotTraced(w http.ResponseWriter, r *http.Request) {
	tracer := pinpoint.FromContext(r.Context())
	defer tracer.NewSpanEvent("NotTrace").EndSpanEvent()
	io.WriteString(w, "handler is not traced")
}

func trace(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return pphttp.WrapHandlerFunc(f)
}

func main() {
	//os.Setenv("PINPOINT_GO_HTTP_SERVER_STATUSCODEERRORS", "5xx,301,400")

	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("GoHttpTest"),
		pinpoint.WithAgentId("GoHttpAgent"),
		pinpoint.WithConfigFile(os.Getenv("HOME") + "/tmp/pinpoint-config.yaml"),

		pphttp.WithHttpServerStatusCodeError([]string{"5xx", "4xx"}),
		pphttp.WithHttpServerExcludeUrl([]string{"/wrapreq*", "/**/*.go", "/*/*.do", "/abc**"}),
		pphttp.WithHttpServerExcludeMethod([]string{"put", "POST"}),
		pphttp.WithHttpServerRecordRequestHeader([]string{"HEADERS-ALL"}),
		pphttp.WithHttpClientRecordRequestHeader([]string{"user-agent", "connection", "foo"}),
		pphttp.WithHttpServerRecordRespondHeader([]string{"content-length"}),
		pphttp.WithHttpClientRecordRespondHeader([]string{"HEADERS-ALL"}),
		pphttp.WithHttpServerRecordRequestCookie([]string{"_octo"}),
		pphttp.WithHttpClientRecordRequestCookie([]string{"HEADERS-ALL"}),
	}
	cfg, _ := pinpoint.NewConfig(opts...)
	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		log.Fatalf("pinpoint agent start fail: %v", err)
	}
	defer agent.Shutdown()

	http.HandleFunc("/", trace(index))
	http.HandleFunc("/wraprequest", trace(wrapRequest))
	http.HandleFunc("/wraprequest/a.zo", trace(wrapRequest))
	http.HandleFunc("/wraprequest/aa/b.zo", trace(wrapRequest))
	http.HandleFunc("/wrapclient", trace(wrapClient))
	http.HandleFunc("/wrapclient/aa/a.go", trace(wrapClient))
	http.HandleFunc("/wrapclient/aa/bb/a.go", trace(wrapClient))
	http.HandleFunc("/wrapclient/c.do", trace(wrapClient))
	http.HandleFunc("/wrapclient/dd/d.do", trace(wrapClient))
	http.HandleFunc("/wrapclient/c@do", trace(wrapClient))
	http.HandleFunc("/abcd", trace(wrapClient))
	http.HandleFunc("/abcd/e.go", trace(wrapClient))
	http.HandleFunc("/notrace", handlerNotTraced)

	http.ListenAndServe(":8000", nil)
}
