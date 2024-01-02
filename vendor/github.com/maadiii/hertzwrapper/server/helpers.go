package server

import (
	"net/http"
	"reflect"
)

func requestType[T any](v any) T {
	p := reflect.TypeOf(v).In(1).Elem()

	return reflect.New(p).Interface().(T)
}

var methods = map[string]string{
	"[GET]":     http.MethodGet,
	"[HEAD]":    http.MethodHead,
	"[POST]":    http.MethodPost,
	"[PUT]":     http.MethodPut,
	"[PATCH]":   http.MethodPatch,
	"[DELETE]":  http.MethodDelete,
	"[CONNECT]": http.MethodConnect,
	"[OPTIONS]": http.MethodOptions,
	"[TRACE]":   http.MethodTrace,
}
