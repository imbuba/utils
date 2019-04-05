package health

import (
	"encoding/json"
	"runtime"
	"time"

	tr "github.com/finnan444/utils/transport"
	"github.com/valyala/fasthttp"
)

type healthResult struct {
	HeapAlloc     uint64 `json:"heapAlloc"`
	HeapObjects   uint64 `json:"heapObjects"`
	LiveObjects   uint64 `json:"liveObjects"`
	NumGoroutines int    `json:"numGoroutines"`
}

var memStat = new(runtime.MemStats)
var result = new(healthResult)

func init() {
	tr.AddGetRoute("/health", health)
}

func health(ctx *fasthttp.RequestCtx, now time.Time, adds ...string) {
	runtime.ReadMemStats(memStat)
	result.HeapAlloc = memStat.HeapAlloc
	result.HeapObjects = memStat.HeapObjects
	result.LiveObjects = memStat.Mallocs - memStat.Frees
	result.NumGoroutines = runtime.NumGoroutine()
	bytes, _ := json.Marshal(result)
	ctx.SetContentType(tr.ApplicationJSON)
	ctx.SetBody(bytes)
}
