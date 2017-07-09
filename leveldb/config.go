package leveldb

type Coder interface {
	Decode([]byte, interface{}) error
	Encode(interface{}) ([]byte, error)
}

type Config struct {
	snapshotInterval int
	coder            Coder
}

type Option func(*Config)

func WithSnapshot(interval int) Option {
	return func(config *Config) {
		config.snapshotInterval = interval
	}
}

func WithCoder(coder Coder) Option {
	return func(config *Config) {
		config.coder = coder
	}
}

func (config *Config) Coder() Coder {
	if config.coder != nil {
		return config.coder
	}
	return new(transcoder)
}

func (config *Config) SnapshotInterval() int {
	if config.snapshotInterval > 0 {
		return config.snapshotInterval
	}
	return 100000
}
