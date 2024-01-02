package server

import (
	"crypto/tls"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
)

func WithKeepAliveTimeout(value time.Duration) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.KeepAliveTimeout = value
		},
	}
}

func WithReadTime(value time.Duration) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.ReadTimeout = value
		},
	}
}

func WithWriteTimeout(value time.Duration) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.WriteTimeout = value
		},
	}
}

func WithIdleTimeout(value time.Duration) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.IdleTimeout = value
		},
	}
}

func WithRedirectTrailingSlash() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.RedirectTrailingSlash = true
		},
	}
}

func WithMaxRequestBodySize(value int) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.MaxRequestBodySize = value
		},
	}
}

func WithMaxKeepBodySize(value int) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.MaxKeepBodySize = value
		},
	}
}

func WithGetOnly() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.GetOnly = true
		},
	}
}

func WithDisableKeepAlive() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.DisableKeepalive = true
		},
	}
}

func WithRedirectFixedPath() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.RedirectFixedPath = true
		},
	}
}

func WithHandleMethodNotAllowed() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.HandleMethodNotAllowed = true
		},
	}
}

func WithUseRawPath() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.UseRawPath = true
		},
	}
}

func WithRemoveExtraSlash() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.RemoveExtraSlash = true
		},
	}
}

func WithUnescapedPathValues() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.UnescapePathValues = true
		},
	}
}

func WithDisablePreParseMultipartForm() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.DisablePreParseMultipartForm = true
		},
	}
}

func WithNoDefaultDate() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.NoDefaultDate = true
		},
	}
}

func WithNoDefaultContentType() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.NoDefaultContentType = true
		},
	}
}

func WithStreamRequestBody() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.StreamRequestBody = true
		},
	}
}

func WithDisablePrintRoute() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.DisablePrintRoute = true
		},
	}
}

func WithNetwork(value string) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.Network = value
		},
	}
}

func WithAddress(value string) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.Addr = value
		},
	}
}

func WithBasePath(value string) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.BasePath = value
		},
	}
}

func WithExitWaitTimeout(value time.Duration) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.ExitWaitTimeout = value
		},
	}
}

func WithTLS(value *tls.Config) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.TLS = value
		},
	}
}

func WithH2C() config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.H2C = true
		},
	}
}

func WithReadBufferSize(value int) config.Option {
	return config.Option{
		F: func(o *config.Options) {
			o.ReadBufferSize = value
		},
	}
}
