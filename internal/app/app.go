package app

import (
	"context"

	"github.com/mistandok/chat-client/internal/cli"
	"github.com/mistandok/chat-client/internal/config"
	"github.com/mistandok/platform_common/pkg/closer"
)

// App ..
type App struct {
	serviceProvider *serviceProvider
	configPath      string
	cliChat         *cli.Chat
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
func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.cliChat.Execute(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	initDepFunctions := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initCliChat,
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

func (a *App) initCliChat(ctx context.Context) error {
	a.cliChat = cli.NewChat(
		a.serviceProvider.Logger(),
		a.serviceProvider.ChatService(ctx),
		a.serviceProvider.ConsoleWriter(ctx),
	)

	return nil
}
