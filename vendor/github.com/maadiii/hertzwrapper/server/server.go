package server

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/maadiii/hertzwrapper/errors"
)

func Run(devMode bool, opts ...config.Option) {
	dev = devMode
	s = server.New(opts...)
	s.Use(recovery.Recovery())

	for key, handlers := range handlersMap {
		action := strings.Split(key, "::")

		s.Handle(action[0], action[1], handlers...)
	}

	s.Spin()
}

var (
	s           *server.Hertz
	dev         bool
	handlersMap map[string][]app.HandlerFunc = make(map[string][]app.HandlerFunc, 0)
)

func Handle[IN any, OUT any](handlers ...func(*Context, IN) (OUT, error)) {
	befores := make([]*Handler[IN, OUT], 0)
	afters := make([]*Handler[IN, OUT], 0)
	main := &Handler[IN, OUT]{}

	for _, h := range handlers {
		handler := &Handler[IN, OUT]{action: h}
		handler.fix()

		if len(handler.path) == 0 && len(handler.method) == 0 {
			befores = append(befores, handler)

			continue
		}

		if len(handler.path) != 0 && len(handler.method) != 0 {
			main = handler

			continue
		}

		afters = append(afters, handler)
	}

	key := fmt.Sprintf("%s::%s::%s::%s", main.method, main.path, main.status, main.contentType)

	for _, h := range befores {
		handlersMap[key] = append(handlersMap[key], handle(h.action, main.method, main.path, main.status, main.contentType))
	}

	handlersMap[key] = append(handlersMap[key], handle(main.action, main.method, main.path, main.status, main.contentType))

	for _, h := range afters {
		handlersMap[key] = append(handlersMap[key], handle(h.action, main.method, main.path, main.status, main.contentType))
	}
}

func handle[IN any, OUT any](handler func(*Context, IN) (OUT, error), method, path, status, contentType string) app.HandlerFunc {
	return func(c context.Context, rctx *app.RequestContext) {
		req := requestType[IN](handler)
		if err := rctx.Bind(req); err != nil {
			rctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				errors.New(fmt.Sprintf( //nolint
					"%s\n#Api=%s#Method=%s#Action=%s",
					err.Error(),
					path,
					method,
					funcPathAndName(handler),
				)),
			)

			return
		}

		ctx := &Context{
			Context: c,
			request: rctx,
		}

		res, err := handler(ctx, req)
		if err != nil {
			handleError(rctx, err)

			return
		}

		rctx.SetContentType(contentType)

		status, err := strconv.Atoi(status)
		if err != nil {
			panic(err)
		}

		if isNil(res) {
			rctx.Status(status)

			return
		}

		rctx.JSON(status, res)
	}
}

func isNil[T any](t T) bool {
	v := reflect.ValueOf(t)
	kind := v.Kind()
	// Must be one of these types to be nillable
	return (kind == reflect.Ptr ||
		kind == reflect.Interface ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Chan ||
		kind == reflect.Func) &&
		v.IsNil()
}

func handleError(ctx *app.RequestContext, err error) {
	switch t := err.(type) {
	case *errors.Error:
		if !dev {
			t.Message = strings.Split(t.Message, "\n")[0]
		}
		ctx.AbortWithStatusJSON(abort(t), t)
	default:
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}

func abort(err *errors.Error) (status int) { //nolint
	switch err {
	case errors.BadRequest:
		status = http.StatusBadRequest
	case errors.Unauthorized:
		status = http.StatusUnauthorized
	case errors.PaymentRequired:
		status = http.StatusPaymentRequired
	case errors.Forbidden:
		status = http.StatusForbidden
	case errors.NotFound:
		status = http.StatusNotFound
	case errors.MethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case errors.NotAcceptable:
		status = http.StatusNotAcceptable
	case errors.ProxyAuthRequired:
		status = http.StatusProxyAuthRequired
	case errors.RequestTimeout:
		status = http.StatusRequestTimeout
	case errors.Conflict:
		status = http.StatusConflict
	case errors.Gone:
		status = http.StatusGone
	case errors.LengthRequired:
		status = http.StatusLengthRequired
	case errors.PreconditionFailed:
		status = http.StatusPreconditionFailed
	case errors.RequestEntityTooLarge:
		status = http.StatusRequestEntityTooLarge
	case errors.RequestURITooLong:
		status = http.StatusRequestURITooLong
	case errors.UnsupportedMediaType:
		status = http.StatusUnsupportedMediaType
	case errors.RequestedRangeNotSatisfiable:
		status = http.StatusRequestedRangeNotSatisfiable
	case errors.ExpectationFailed:
		status = http.StatusExpectationFailed
	case errors.Teapot:
		status = http.StatusTeapot
	case errors.MisdirectedRequest:
		status = http.StatusUnauthorized
	case errors.UnprocessableEntity:
		status = http.StatusUnprocessableEntity
	case errors.Locked:
		status = http.StatusLocked
	case errors.FailedDependency:
		status = http.StatusFailedDependency
	case errors.TooEarly:
		status = http.StatusTooEarly
	case errors.UpgradeRequired:
		status = http.StatusUpgradeRequired
	case errors.PreconditionRequired:
		status = http.StatusPreconditionFailed
	case errors.TooManyRequests:
		status = http.StatusTooManyRequests
	case errors.RequestHeaderFieldsTooLarge:
		status = http.StatusRequestHeaderFieldsTooLarge
	case errors.UnavailableForLegalReasons:
		status = http.StatusUnavailableForLegalReasons
	case errors.NotImplemented:
		status = http.StatusNotImplemented
	case errors.BadGateway:
		status = http.StatusBadGateway
	case errors.ServiceUnavailable:
		status = http.StatusServiceUnavailable
	case errors.GatewayTimeout:
		status = http.StatusGatewayTimeout
	case errors.HTTPVersionNotSupported:
		status = http.StatusHTTPVersionNotSupported
	case errors.VariantAlsoNegotiates:
		status = http.StatusVariantAlsoNegotiates
	case errors.InsufficientStorage:
		status = http.StatusInsufficientStorage
	case errors.LoopDetected:
		status = http.StatusLoopDetected
	case errors.NotExtended:
		status = http.StatusNotExtended
	case errors.NetworkAuthenticationRequired:
		status = http.StatusNetworkAuthenticationRequired
	default:
		status = http.StatusInternalServerError
	}

	return //nolint
}
