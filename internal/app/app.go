package app

import (
	"context"
	"github.com/mistandok/chat-client/pkg/auth_v1"

	"github.com/mistandok/chat-client/internal/config"
	"github.com/mistandok/platform_common/pkg/closer"
)

// App ..
type App struct {
	serviceProvider *serviceProvider
	configPath      string
}

// NewApp ..
func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run ..
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	logger := a.serviceProvider.Logger()
	logger.Info().Msg(a.serviceProvider.AuthConfig().Address())
	logger.Info().Msg(a.serviceProvider.ChatConfig().Address())
	client := a.serviceProvider.AuthV1Client(context.TODO())
	_, err := client.Login(context.TODO(), &auth_v1.LoginRequest{
		Email:    "anton",
		Password: "anton@mail.ru",
	})
	if err != nil {
		logger.Err(err).Msg("error")
	}

	a.serviceProvider.ChatV1Client(context.TODO())
	a.serviceProvider.UserV1Client(context.TODO())

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	initDepFunctions := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, f := range initDepFunctions {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
