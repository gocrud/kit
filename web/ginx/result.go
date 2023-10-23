package ginx

import (
	"net/http"
)

type Result interface {
	getStatus() int
	getBody() any
}

type result struct {
	status int
	body   any
	err    error
}

func (r *result) getStatus() int {
	return r.status
}

func (r *result) getBody() any {
	return r.body
}

func (r *result) Body() any {
	return r.body
}

func (r *result) Status() int {
	return r.status
}

func (r *result) SetBody(value any) {
	r.body = value
}

func (r *result) SetStatus(status int) {
	r.status = status
}

func (r *result) GetError() error {
	return r.err
}

type KV struct {
	Key   string
	Value interface{}
}

// Success 函数返回一个成功的 Result 实例，其中包含 HTTP 状态码为 200 和指定的数据。
// 参数 data 是任意类型的数据。
func Success(data any) Result {
	return &result{
		status: http.StatusOK,
		body:   data,
	}
}

func wrapError(code int, err error, values ...KV) Result {
	var val = make(map[string]interface{})
	val["msg"] = err.Error()
	for _, v := range values {
		val[v.Key] = v.Value
	}
	return &result{
		err:    err,
		status: code,
		body:   val,
	}
}

// Bad 函数返回一个失败的 Result 实例，其中包含 HTTP 状态码为 400 和指定的错误信息。
func Bad(err error, values ...KV) Result {
	return wrapError(http.StatusBadRequest, err, values...)
}

// Unauthorized 函数返回一个失败的 Result 实例，其中包含 HTTP 状态码为 401 和指定的错误信息。
func Unauthorized(err error, values ...KV) Result {
	return wrapError(http.StatusUnauthorized, err, values...)
}

// Forbidden 函数返回一个失败的 Result 实例，其中包含 HTTP 状态码为 403 和指定的错误信息。
func Forbidden(err error, values ...KV) Result {
	return wrapError(http.StatusForbidden, err, values...)
}

// NotFound 函数返回一个失败的 Result 实例，其中包含 HTTP 状态码为 404 和指定的错误信息。
func NotFound(err error, values ...KV) Result {
	return wrapError(http.StatusNotFound, err, values...)
}

// Internal 函数返回一个失败的 Result 实例，其中包含 HTTP 状态码为 500 和指定的错误信息。
func Internal(err error, values ...KV) Result {
	return wrapError(http.StatusInternalServerError, err, values...)
}
