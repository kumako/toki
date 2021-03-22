package api_test

import (
	"toki/api"
	"toki/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

var _ = Describe("TimeZone API", func() {
	ctx := fasthttp.RequestCtx{}

	Describe("Call TimeZone API Handler", func() {
		api.TZHandler(&ctx)
	})

	It("Should returns 200 status", func() {
		Expect(ctx.Response.StatusCode()).To(Equal(200))
	})

	It("Should returns default timezone", func() {
		Expect(ctx.Response.Body()).To(ContainSubstring(config.DefaultTZ))
	})
})
