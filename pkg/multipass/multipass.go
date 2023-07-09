package multipass

import (
	"context"

	"github.com/loft-sh/log"
	"github.com/mrsimonemms/devpod-provider-multipass/pkg/options"
)

type Provider struct {
	Config *options.Options
	Log    log.Logger
}

// func execCmd() {}

func New(cfg *options.Options) *Provider {
	return &Provider{
		Config: cfg,
		Log:    log.Default,
	}
}

func (p *Provider) Create(ctx context.Context) error {
	return nil
}
