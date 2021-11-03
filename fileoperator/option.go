package fileoperator

type Option func(o *Options)

type Options struct {
	IsZip bool
}

// WithZip 启用压缩
func WithZip() Option{
	return func(options *Options) {
		options.IsZip = true
	}
}

func newOption(opts ...Option)Options{
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	return o
}