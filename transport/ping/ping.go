package ping

import (
	"time"

	"github.com/imbuba/utils/transport"
	"github.com/valyala/fasthttp"
)

var (
	pingResponse = []byte("OK")
)

func init() {
	transport.AddGetRoute("/ping", ping)
	transport.AddGetRoute("/blind-ping", ping)
}

func ping(ctx *fasthttp.RequestCtx, now time.Time, adds ...string) {
	ctx.SetContentType("text/plain; charset=utf-8")
	ctx.SetBody(pingResponse)
}
