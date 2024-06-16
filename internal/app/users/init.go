package users

import (
	"DatabaseLab/internal/app"
	"DatabaseLab/internal/app/users/dao"
	"DatabaseLab/internal/app/users/router"
	"DatabaseLab/kernel"
	"context"
	"sync"
)

type (
	Users struct {
		Name string
		app.UnimplementedModule
	}
)

func (p *Users) Info() string {
	return p.Name
}

func (p *Users) PreInit(engine *kernel.Engine) error {
	err := dao.InitOp(engine.OpenGauss.DB)
	return err
}

func (p *Users) Init(*kernel.Engine) error {
	return nil
}

func (p *Users) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Users) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppUsersInit(engine.GIN)
	return nil
}

func (p *Users) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Users) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Users) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}
