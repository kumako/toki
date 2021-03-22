package api

import (
	"log"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/markbates/pkger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func routes() func(*fasthttp.RequestCtx) {

	router := router.New()

	// Static Files
	www, err := pkger.Open("/www")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(www)
	router.NotFound = fasthttpadaptor.NewFastHTTPHandler(fs)

	// Timezone API
	router.GET("/api/timezone", TZHandler)

	return router.Handler
}
