package dataparse

import (
	"errors"
	"io"
	"slices"
)

type ReadConfig struct {
	channelSize    int
	errChannelSize int

	reader  io.Reader
	closers []func() error
}

func newReadConfig(opts ...ReadOption) *ReadConfig {
	cfg := &ReadConfig{
		channelSize:    100,
		errChannelSize: 1,
		closers:        []func() error{},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func (cfg ReadConfig) channels() (chan Map, chan error) {
	return make(chan Map, cfg.channelSize), make(chan error, 1)
}

func (cfg ReadConfig) Close() error {
	var retErr error
	slices.Reverse(cfg.closers)
	for _, closer := range cfg.closers {
		if err := closer(); err != nil {
			retErr = errors.Join(retErr, err)
		}
	}
	return retErr
}

type ReadOption func(*ReadConfig)

func WithChannelSize(i int) ReadOption {
	return func(opt *ReadConfig) {
		opt.channelSize = i
	}
}

func WithErrChannelSize(i int) ReadOption {
	return func(opt *ReadConfig) {
		opt.errChannelSize = i
	}
}
