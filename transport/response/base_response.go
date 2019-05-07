package response

import (
	"sync"

	"github.com/imbuba/utils/pool"
)

var (
	responses = sync.Pool{
		New: func() interface{} {
			return &BasicResponse{}
		},
	}
)

// BasicResponse struct
type BasicResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

// BasicResponser interface
type BasicResponser interface {
	SetCode(code int)
	SetMessage(message string)
	SetError(code int, message string)
}

// SetCode impl
func (resp *BasicResponse) SetCode(code int) {
	resp.Code = code
}

// SetMessage impl
func (resp *BasicResponse) SetMessage(message string) {
	resp.Msg = message
}

// SetError impl
func (resp *BasicResponse) SetError(code int, message string) {
	resp.Code = code
	resp.Msg = message
}

// Reuse impl
func (resp *BasicResponse) Reuse() {
	resp.Code = 0
	resp.Msg = ""
	if reuse, ok := resp.Payload.(pool.Reusable); ok {
		reuse.Reuse()
	}
	resp.Payload = nil
	responses.Put(resp)
}

// GetResponse returns basic response
func GetResponse() *BasicResponse {
	return responses.Get().(*BasicResponse)
}
