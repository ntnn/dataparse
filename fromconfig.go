package dataparse

import (
	"errors"
	"io"
	"slices"
)

type fromConfig struct {
	channelSize int

	separator string
	trimSpace bool
	headers   []string

	reader  io.Reader
	closers []func() error
}

func newFromConfig(opts ...FromOption) *fromConfig {
	cfg := &fromConfig{
		channelSize: 100,
		separator:   ",",
		trimSpace:   true,
		headers:     []string{},
		closers:     []func() error{},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func (cfg fromConfig) Close() error {
	var retErr error
	slices.Reverse(cfg.closers)
	for _, closer := range cfg.closers {
		if err := closer(); err != nil {
			retErr = errors.Join(retErr, err)
		}
	}
	return retErr
}

type FromOption func(*fromConfig)

// WithChannelSize defines the buffer size of channels for functions
// returning channels.
// Defaults to 100.
func WithChannelSize(i int) FromOption {
	return func(opt *fromConfig) {
		opt.channelSize = i
	}
}

// WithSeparator defines the separator to use when splitting strings or
// when reading formats with delimiters.
// Defaults to ",".
// This does not apply to unmarshalled values like JSON.
func WithSeparator(sep string) FromOption {
	return func(opt *fromConfig) {
		opt.separator = sep
	}
}

// WithTrimSpace defines whether values are trimmed when parsing input.
// Defaults to true.
// This does not apply to unmarshalled values like JSON.
func WithTrimSpace(trim bool) FromOption {
	return func(opt *fromConfig) {
		opt.trimSpace = trim
	}
}

// WithHeaders defines which headers are expected when reading delimited
// formats like csv. If no headers are set the input is expected to have
// headers.
// Defaults to []string.
func WithHeaders(headers ...string) FromOption {
	return func(opt *fromConfig) {
		opt.headers = headers
	}
}
