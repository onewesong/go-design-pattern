package builder

import "time"

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10 * time.Second
	defaultCaching = false
)

type options struct {
	timeout time.Duration
	caching bool
}

type optionFunc func(*options) error

func WithTimeout(timeout time.Duration) optionFunc {
	return func(o *options) error {
		// 此处可做一些参数校验, 抛出error
		o.timeout = timeout
		return nil
	}
}

func WithCaching(caching bool) optionFunc {
	return func(o *options) error {
		o.caching = caching
		return nil
	}
}

func NewConnection(addr string, opts ...optionFunc) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		// 通过传入指针改变options的值
		err := o(&options)
		if err != nil {
			return nil, err
		}
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}
