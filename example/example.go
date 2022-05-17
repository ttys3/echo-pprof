package main

import (
	"github.com/labstack/echo/v4"
	echopprof "github.com/ttys3/echo-pprof/v4"
	"log"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	echopprof.Wrap(e)

	// echopprof also plays well with *echo.Group
	// prefix := "/debug/pprof"
	// group := e.Group(prefix)
	// echopprof.WrapGroup(prefix, group)

	log.Println("pprof URL: http://127.0.0.1:3000/debug/pprof/")
	log.Fatal(e.Start(":3000"))
}
