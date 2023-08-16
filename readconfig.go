package dataparse

type ReadConfig struct {
	channelSize   int
	finishChannel chan struct{}
}

func newReadConfig(opts ...ReadOption) *ReadConfig {
	cfg := &ReadConfig{
		channelSize: 100,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func (cfg *ReadConfig) closeFinishChannel() {
	if cfg.finishChannel != nil {
		close(cfg.finishChannel)
	}
}

func (cfg ReadConfig) channels() (chan Map, chan error) {
	return make(chan Map, cfg.channelSize), make(chan error, 1)
}

type ReadOption func(*ReadConfig)

func WithChannelSize(i int) ReadOption {
	return func(opt *ReadConfig) {
		opt.channelSize = i
	}
}

func withFinishChannel(ch chan struct{}) ReadOption {
	return func(opt *ReadConfig) {
		opt.finishChannel = ch
	}
}
