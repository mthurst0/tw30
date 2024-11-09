package appcore

import (
	"context"

	"github.com/mthurst0/tw30/internal/appenv"
)

type Clients interface {
	Settings() appenv.Settings
	Ctx() context.Context
}

type clients struct {
	settings appenv.Settings
	ctx      context.Context
}

func (c *clients) Settings() appenv.Settings {
	return c.settings
}

func (c *clients) Ctx() context.Context {
	return c.ctx
}

func newClients() (*clients, error) {
	ctx := context.Background()
	return &clients{
		settings: appenv.MustSettings(),
		ctx:      ctx,
	}, nil
}

var OverrideClient Clients

// Must returns concrete implementations of the core interfaces
// to external services used by tw30. Further, it ensures that
// our Dynamo DB table is created and available for use (or panics).
func Must() Clients {
	if OverrideClient != nil {
		return OverrideClient
	}
	c, err := newClients()
	if err != nil {
		panic(err)
	}
	return c
}
