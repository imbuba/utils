package transport

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/imbuba/utils/math/ints"
	"github.com/valyala/fasthttp"
)

var (
	postRoutes    = make(map[string]RouterFunc)
	postRegRoutes = make(map[*regexp.Regexp]RouterFunc)
	getRoutes     = make(map[string]RouterFunc)
	getRegRoutes  = make(map[*regexp.Regexp]RouterFunc)
	clientsPool   = sync.Pool{
		New: func() interface{} {
			return &fasthttp.Client{}
		},
	}
	timings    = make(map[string]*median)
	timingsReg = make(map[*regexp.Regexp]*median)
	logger     = log.New(os.Stderr, "\n-----------------------------\n", log.LstdFlags)
)

func init() {
	AddGetRoute("/internal/stats", handlerInternalStats)
}

type median struct {
	sync.Mutex
	Min, Max, Total, Count time.Duration
}

func (m *median) Update(d time.Duration) {
	m.Lock()
	if m.Min == 0 || m.Min > d {
		m.Min = d
	}
	if m.Max < d {
		m.Max = d
	}
	m.Total += d
	m.Count++
	m.Unlock()
}

func (m *median) String() string {
	if m.Count > 0 {
		return fmt.Sprintf(": {\"min\":%v, \"max\":%v, \"med\":%v}\n", m.Min, m.Max, m.Total/m.Count)
	}
	return ": Not enough stats\n"
}

// SetLogger sets new logger
func SetLogger(lgr *log.Logger) {
	logger = lgr
}

// RouterFunc router function
type RouterFunc func(*fasthttp.RequestCtx, time.Time, ...string)

// AddGetRoute adds get route
func AddGetRoute(path string, handler RouterFunc) {
	getRoutes[path] = handler
	timings["[GET] "+path] = &median{}
}

// AddGetRegexpRoute adds get route. For example /accounts/([0-9]+)/suggest/.
//The result of regex will be passed as s third parameter in router.RouterFunc
func AddGetRegexpRoute(path string, handler RouterFunc) {
	if re, err := regexp.Compile(path); err == nil {
		getRegRoutes[re] = handler
		timingsReg[re] = &median{}
	}
}

// AddPostRoute adds post route
func AddPostRoute(path string, handler RouterFunc) {
	postRoutes[path] = handler
	timings["[POST] "+path] = &median{}
}

// AddPostRegexpRoute adds post route
func AddPostRegexpRoute(path string, handler RouterFunc) {
	if re, err := regexp.Compile(path); err == nil {
		postRegRoutes[re] = handler
		timingsReg[re] = &median{}
	}
}

// ProcessRouting returns router
func ProcessRouting(server PathesLogger) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		now := time.Now()
		path := string(ctx.Path())
		reqID := ctx.ID()
		if ctx.IsPost() {
			body := ctx.PostBody()
			if logFlag := server.GetLogFlag(path); (logFlag & ToLog) != 0 {
				if (logFlag & FullLog) != 0 {
					logger.Printf("[POST %s %d][Request] %s\n", path, reqID, body)
				} else {
					logger.Printf("[POST %s %d][Request] %s\n", path, reqID, body[:ints.MinInt(len(body), 255)])
				}
			}
			if handler, ok := postRoutes[path]; ok {
				handler(ctx, now)
				timings["[POST] "+path].Update(time.Since(now))
			} else {
				for k, v := range postRegRoutes {
					adds := k.FindStringSubmatch(path)
					if len(adds) > 1 {
						v(ctx, now, adds[1:]...)
						return
					}
				}
				ctx.Error("Not found", fasthttp.StatusNotFound)
			}
		} else if ctx.IsGet() {
			if logFlag := server.GetLogFlag(path); (logFlag & ToLog) != 0 {
				logger.Printf("[GET %s %d][Request] %s\n", path, reqID, ctx.QueryArgs().QueryString())
			}
			if handler, ok := getRoutes[path]; ok {
				handler(ctx, now)
				timings["[GET] "+path].Update(time.Since(now))
			} else {
				for k, v := range getRegRoutes {
					adds := k.FindStringSubmatch(path)
					if len(adds) > 1 {
						v(ctx, now, adds[1:]...)
						timingsReg[k].Update(time.Since(now))
						return
					}
				}
				ctx.Error("Not found", fasthttp.StatusNotFound)
			}
		} else {
			ctx.Error("Not found", fasthttp.StatusNotFound)
		}
	}
}

// ProcessSimpleRouting returns router
func ProcessSimpleRouting() func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		now := time.Now()
		path := string(ctx.Path())
		if ctx.IsPost() {
			if handler, ok := postRoutes[path]; ok {
				handler(ctx, now)
				timings["[POST] "+path].Update(time.Since(now))
			} else {
				for k, v := range postRegRoutes {
					adds := k.FindStringSubmatch(path)
					if len(adds) > 1 {
						v(ctx, now, adds[1:]...)
						return
					}
				}
				ctx.Error("Not found", fasthttp.StatusNotFound)
			}
		} else if ctx.IsGet() {
			if handler, ok := getRoutes[path]; ok {
				handler(ctx, now)
				timings["[GET] "+path].Update(time.Since(now))
			} else {
				for k, v := range getRegRoutes {
					adds := k.FindStringSubmatch(path)
					if len(adds) > 1 {
						v(ctx, now, adds[1:]...)
						timingsReg[k].Update(time.Since(now))
						return
					}
				}
				ctx.Error("Not found", fasthttp.StatusNotFound)
			}
		} else {
			ctx.Error("Not found", fasthttp.StatusNotFound)
		}
	}
}

// GetHTTPClient returns client from pool
func GetHTTPClient() *fasthttp.Client {
	return clientsPool.Get().(*fasthttp.Client)
}

// PutHTTPClient returns client to pool
func PutHTTPClient(client *fasthttp.Client) {
	clientsPool.Put(client)
}

func handlerInternalStats(ctx *fasthttp.RequestCtx, now time.Time, adds ...string) {
	for k, v := range timings {
		ctx.WriteString(k)
		ctx.WriteString(v.String())
	}
	for k, v := range timingsReg {
		ctx.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
}
