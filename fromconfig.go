package dataparse

import (
	"errors"
	"io"
	"slices"
)

type FromConfig struct {
	channelSize    int
	errChannelSize int

	separator string
	trimSpace bool
	headers   []string

	reader  io.Reader
	closers []func() error

	// Not something that necessarily needs to be done but allows for
	// an easy cleanup with one call to .Close.
	mapCh chan Map
	errCh chan error
}

func newFromConfig(opts ...FromOption) *FromConfig {
	cfg := &FromConfig{
		channelSize:    100,
		errChannelSize: 1,
		separator:      ",",
		trimSpace:      true,
		headers:        []string{},
		closers:        []func() error{},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func (cfg *FromConfig) channels() (chan Map, chan error) {
	cfg.mapCh = make(chan Map, cfg.channelSize)
	cfg.errCh = make(chan error, cfg.errChannelSize)
	return cfg.mapCh, cfg.errCh
}

func (cfg FromConfig) Close() error {
	if cfg.mapCh != nil {
		close(cfg.mapCh)
	}
	if cfg.errCh != nil {
		close(cfg.errCh)
	}

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

// WithChannelSize defines the buffer size of channels for functions
// returning channels.
// Defaults to 100.
func WithChannelSize(i int) FromOption {
	return func(opt *FromConfig) {
		opt.channelSize = i
	}
}

// WithErrChannelSize defines the buffer size of error channels for
// functions returning error channels.
// Defaults to 1.
func WithErrChannelSize(i int) FromOption {
	return func(opt *FromConfig) {
		opt.errChannelSize = i
	}
}

// WithSeparator defines the separator to use when splitting strings or
// when reading formats with delimiters.
// Defaults to ",".
// This does not apply to unmarshalled values like JSON.
func WithSeparator(sep string) FromOption {
	return func(opt *FromConfig) {
		opt.separator = sep
	}
}

// WithTrimSpace defines whether values are trimmed when parsing input.
// Defaults to true.
// This does not apply to unmarshalled values like JSON.
func WithTrimSpace(trim bool) FromOption {
	return func(opt *FromConfig) {
		opt.trimSpace = trim
	}
}

// WithHeaders defines which headers are expected when reading delimited
// formats like csv. If no headers are set the input is expected to have
// headers.
// Defaults to []string.
func WithHeaders(headers []string) FromOption {
	return func(opt *FromConfig) {
		opt.headers = headers
	}
}
