package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
)

func indexMux(w http.ResponseWriter, r *http.Request) {
	//var i http.ResponseWriter
	//i.Header() //panic

	io.WriteString(w, "hello world, mux")
}

func outGoing(w http.ResponseWriter, r *http.Request) {
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

func main() {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("GoHttpMuxTest"),
		pinpoint.WithAgentId("GoHttpMuxAgent"),
		pinpoint.WithAgentName("GoHttpMuxTest_GoHttpMuxAgent"),
		pinpoint.WithConfigFile(os.Getenv("HOME") + "/tmp/pinpoint-config.yaml"),
		pphttp.WithHttpServerStatusCodeError([]string{"500", "400"}),
		pphttp.WithHttpServerRecordRequestHeader([]string{"user-agent", "connection", "foo"}),
	}

	//os.Setenv("PINPOINT_GO_ACTIVEPROFILE", "real")

	cfg, err := pinpoint.NewConfig(opts...)
	if err != nil {
		log.Fatalf("pinpoint configuration fail: %v", err)
	}

	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		log.Fatalf("pinpoint agent start fail: %v", err)
	}
	defer agent.Shutdown()

	mux := pphttp.NewServeMux()
	mux.Handle("/foo", http.HandlerFunc(indexMux))
	mux.HandleFunc("/bar", outGoing)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
