package transport

import (
	"crypto"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/finnan444/utils/math/ints"
	"github.com/finnan444/utils/pool"
	"github.com/finnan444/utils/transport/request"
	"github.com/finnan444/utils/transport/response"
	"github.com/valyala/fasthttp"
)

var (
	hashPool = sync.Pool{
		New: func() interface{} {
			return crypto.MD5.New()
		},
	}
)

// GetResponse returns base response
func GetResponse() *response.BasicResponse {
	return response.GetResponse()
}

// Decode decodes request to object
func Decode(ctx *fasthttp.RequestCtx, to interface{}) bool {
	if err := json.Unmarshal(ctx.PostBody(), to); err != nil {
		log.Printf("[%s] has decode error: %v", ctx.Path(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return false
	}
	return true
}

// Authenticate do smth
func Authenticate(request request.BasicRequester, response response.BasicResponser, secret string, server PathesLogger) bool {
	hash := hashPool.Get().(hash.Hash)
	io.WriteString(hash, strconv.Itoa(request.GetTime()))
	io.WriteString(hash, secret)
	var sign string = fmt.Sprintf("%x", hash.Sum(nil))
	hash.Reset()
	hashPool.Put(hash)
	if sign != request.GetSignature() {
		response.SetError(SignatureMismatch, "Signature mismatched")
		return false
	}
	return true
}

// AuthenticateUser do smth
func AuthenticateUser(request request.UserBasicRequester, response response.BasicResponser, secret string, server PathesLogger) bool {
	hash := hashPool.Get().(hash.Hash)
	io.WriteString(hash, request.GetUser())
	io.WriteString(hash, secret)
	io.WriteString(hash, strconv.Itoa(request.GetTime()))
	sign := fmt.Sprintf("%x", hash.Sum(nil))
	hash.Reset()
	hashPool.Put(hash)
	if sign != request.GetSignature() {
		response.SetCode(SignatureMismatch)
		response.SetMessage("Signature mismatched")
		return false
	}
	return true
}

// SendResponse do smth
func SendResponse(ctx *fasthttp.RequestCtx, response pool.Reusable, startTime time.Time, server PathesLogger) {
	js, err := json.Marshal(response)
	response.Reuse()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType(ApplicationJSONUTF8)
	ctx.SetBody(js)
	path := string(ctx.Path())
	reqID := ctx.ID()
	if logFlag := server.GetLogFlag(path); (logFlag & ToLog) != 0 {
		if (logFlag & FullLog) != 0 {
			logger.Printf("[%s %s %d][Response %s] %s\n", ctx.Method(), path, reqID, time.Since(startTime), js)
		} else {
			logger.Printf("[%s %s %d][Response %s] %s\n", ctx.Method(), path, reqID, time.Since(startTime), js[:ints.MinInt(len(js), 255)])
		}
	}
}

// GenerateRandom generates random string
func GenerateRandom(salt string) string {
	hash := hashPool.Get().(hash.Hash)
	io.WriteString(hash, strconv.FormatInt(time.Now().UnixNano(), 10))
	io.WriteString(hash, salt)
	result := fmt.Sprintf("%x", hash.Sum(nil))
	hash.Reset()
	hashPool.Put(hash)
	return result
}
