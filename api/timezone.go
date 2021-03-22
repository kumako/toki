package api

import (
	"fmt"
	"log"
	"time"
	"toki/config"

	"github.com/valyala/fasthttp"
)

// TZHandler API handler that returns default timezone
func TZHandler(ctx *fasthttp.RequestCtx) {
	// Log execution time
	start := time.Now()
	defer log.Printf("HTTP - API Call on %s - %s", ctx.URI(), time.Since(start))

	// Build http response
	response := fmt.Sprintf(`{
	"timezone": "%s"
}`, config.DefaultTZ)

	// Send http response
	fmt.Fprint(ctx, response)
	ctx.SetStatusCode(200)
	ctx.SetContentType("application/json")
}
