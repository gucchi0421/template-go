//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
)

type AppContainer struct{}

func NewAppContainer() (*AppContainer, error) {
	wire.Build(
		wire.Struct(new(AppContainer), "*"),
	)
	return &AppContainer{}, nil
}
