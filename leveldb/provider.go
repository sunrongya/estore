package leveldb

import (
	"github.com/sunrongya/estore"
)

type Provider struct {
	_path             string
	_snapshotInterval int
	_coder            Coder
}

func (provider *Provider) GetState() estore.ProviderState {
	return newProviderState(provider)
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
