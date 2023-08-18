package dataparse

import (
	"errors"
	"io"
	"slices"
)

type FromConfig struct {
	channelSize    int
	errChannelSize int

	reader  io.Reader
	closers []func() error
}

func newFromConfig(opts ...FromOption) *FromConfig {
	cfg := &FromConfig{
		channelSize:    100,
		errChannelSize: 1,
		closers:        []func() error{},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func (cfg FromConfig) channels() (chan Map, chan error) {
	return make(chan Map, cfg.channelSize), make(chan error, 1)
}

func (cfg FromConfig) Close() error {
	var retErr error
	slices.Reverse(cfg.closers)
	for _, closer := range cfg.closers {
		if err := closer(); err != nil {
			retErr = errors.Join(retErr, err)
		}
	}
	return retErr
}

type FromOption func(*FromConfig)

func WithChannelSize(i int) FromOption {
	return func(opt *FromConfig) {
		opt.channelSize = i
	}
}

func WithErrChannelSize(i int) FromOption {
	return func(opt *FromConfig) {
		opt.errChannelSize = i
	}
}
