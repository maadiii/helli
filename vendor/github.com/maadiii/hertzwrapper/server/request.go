package server

import (
	"context"
	"net"

	"github.com/cloudwego/hertz/pkg/app"
)

type Context struct {
	context.Context //nolint
	request         *app.RequestContext
}

func (ctx *Context) Host() []byte {
	return ctx.request.Host()
}

func (ctx *Context) RemoteAddr() net.Addr {
	return ctx.request.RemoteAddr()
}

func (ctx *Context) Path() []byte {
	return ctx.request.URI().Path()
}

func (ctx *Context) FullPath() string {
	return ctx.request.FullPath()
}

func (ctx *Context) Set(key string, value any) {
	ctx.request.Set(key, value)
}

func (ctx *Context) Value(key any) any {
	return ctx.request.Value(key)
}

func (ctx *Context) SetIdentity(identity Identity) {
	ctx.request.Set(identityKey, identity)
}

func (ctx *Context) Identity() Identity {
	return ctx.request.Value(identityKey).(Identity)
}

type Identity map[string]any

const identityKey = "identity"
