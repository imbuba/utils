package ping

import (
	"time"

	tr "github.com/imbuba/utils/transport"
	"github.com/valyala/fasthttp"
)

var (
	pingResponse = []byte("OK")
)

func init() {
	tr.AddGetRoute("/ping", ping)
	tr.AddGetRoute("/blind-ping", ping)
}

func ping(ctx *fasthttp.RequestCtx, now time.Time, adds ...string) {
	ctx.SetContentType("text/plain; charset=utf-8")
	ctx.SetBody(pingResponse)
}
