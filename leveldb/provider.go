package leveldb

import (
	"github.com/AsynkronIT/protoactor-go/persistence"
)

type Provider struct {
	_path             string
	_snapshotInterval int
	_coder            Coder
}

func (provider *Provider) GetState() persistence.ProviderState {
	return &ProviderState{
		Provider: provider,
	}
}

func New(path string, options ...Option) *Provider {
	config := &Config{}
	for _, option := range options {
		option(config)
	}

	provider := &Provider{
		_path:             path,
		_snapshotInterval: config.SnapshotInterval(),
		_coder:            config.Coder(),
	}

	return provider
}

func (provider *Provider) Path() string {
	return provider._path
}
