package request

import (
	"strings"
)

// BasicRequest struct
type BasicRequest struct {
	Time      int    `json:"time"`
	Signature string `json:"signature"`
}

// BasicRequester interface
type BasicRequester interface {
	GetTime() int
	GetSignature() string
}

// GetTime impl
func (req *BasicRequest) GetTime() int {
	return req.Time
}

// GetSignature impl
func (req *BasicRequest) GetSignature() string {
	return req.Signature
}

// UserBasicRequester interface
type UserBasicRequester interface {
	BasicRequester
	GetUser() string
}

// UserBasicRequest struct
type UserBasicRequest struct {
	BasicRequest
	UserClient string `json:"user"`
}

// GetUser impl
func (req *UserBasicRequest) GetUser() string {
	return req.UserClient
}

// User impl
func (req *UserBasicRequest) User() string {
	return strings.ToLower(req.UserClient)
}

// BaseIDRequest struct
type BaseIDRequest struct {
	ID int `json:"id"`
}
