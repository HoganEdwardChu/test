package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/echo"
)

func hello(c echo.Context) error {
	return c.String(200, "Hello World!!")
}

func myError(c echo.Context) error {
	tracer := pinpoint.TracerFromRequestContext(c.Request())
	tracer.SpanEvent().SetError(errors.New("my error message"))

	return c.String(500, "Error!!")
}

func sleep() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed).Intn(10000)
	time.Sleep(time.Duration(random+1) * time.Millisecond)
}

func main() {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("GoEchoTest"),
		pinpoint.WithAgentId("GoEchoTestAgent"),
		pinpoint.WithHttpUrlStatEnable(true),
		pinpoint.WithConfigFile(os.Getenv("HOME") + "/tmp/pinpoint-config.yaml"),
	}
	cfg, _ := pinpoint.NewConfig(opts...)
	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		log.Fatalf("pinpoint agent start fail: %v", err)
	}
	defer agent.Shutdown()

	e := echo.New()
	e.Use(ppecho.Middleware())

	//e.GET("/hello", ppecho.WrapHandler(hello))
	//e.GET("/error", ppecho.WrapHandler(myError))

	e.GET("/users/:id", func(c echo.Context) error {
		sleep()
		return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users/:name/age/:old", f)

	e.GET("/users/files/*", func(c echo.Context) error {
		sleep()
		return c.String(http.StatusOK, "/users/files/*")
	})

	e.Start(":9000")
}

func f(c echo.Context) error {
	sleep()
	return c.String(http.StatusOK, "/user/:name/age/:old")
}
