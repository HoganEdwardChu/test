package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
)

type handler struct {
}

func doRequest(tracer pinpoint.Tracer) (string, error) {
	req, err := http.NewRequest("GET", "http://localhost:9000/hello", nil)
	if nil != err {
		return "", err
	}
	client := &http.Client{}
	tracer = pphttp.NewHttpClientTracer(tracer, "doRequest", req)
	resp, err := client.Do(req)
	pphttp.EndHttpClientTracer(tracer, resp, err)

	if nil != err {
		return "", err
	}

	fmt.Println("response code is", resp.StatusCode)

	ret, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	func() {
		defer tracer.NewSpanEvent("f1").EndSpanEvent()

		func() {
			defer tracer.NewSpanEvent("f2").EndSpanEvent()
			time.Sleep(2 * time.Millisecond)
		}()
		time.Sleep(2 * time.Millisecond)
	}()

	return string(ret), nil
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	tracer := pphttp.NewHttpServerTracer(req, "Go Http Server")
	defer tracer.EndSpan()

	status := http.StatusOK
	if req.URL.String() == "/hello" {
		ret, _ := doRequest(tracer)

		io.WriteString(writer, ret)
		time.Sleep(10 * time.Millisecond)
	} else {
		status = http.StatusNotFound
		writer.WriteHeader(status)
	}

	pphttp.RecordHttpServerResponse(tracer, status, writer.Header())
}

func main() {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("testGoAgent"),
		pinpoint.WithAgentId("testGoAgentId"),
		pinpoint.WithCollectorHost("localhost"),
	}
	c, _ := pinpoint.NewConfig(opts...)
	agent, _ := pinpoint.NewAgent(c)
	defer agent.Shutdown()

	server := http.Server{
		Addr:    ":8000",
		Handler: &handler{},
	}

	server.ListenAndServe()
}
