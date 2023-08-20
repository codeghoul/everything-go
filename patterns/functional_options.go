package patterns

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOptions() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

type Server struct {
	opts Opts
}

func WithTLS(opts *Opts) {
	opts.tls = true
}

func WithMaxConn(n int) OptFunc {
	return func(o *Opts) {
		o.maxConn = n
	}
}

func WithID(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func NewServer(opts ...OptFunc) *Server {

	o := defaultOptions()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		opts: o,
	}
}
