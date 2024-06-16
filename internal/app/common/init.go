package common

import (
	"DatabaseLab/internal/app"
	"DatabaseLab/internal/app/common/dao"
	"DatabaseLab/internal/app/common/router"
	"DatabaseLab/kernel"
	"context"
	"sync"
)

type (
	Common struct {
		Name string
		app.UnimplementedModule
	}
)

func (p *Common) Info() string {
	return p.Name
}

func (p *Common) PreInit(engine *kernel.Engine) error {
	dao.InitOP(engine.OpenGauss.DB)
	return nil
}

func (p *Common) Init(*kernel.Engine) error {
	return nil
}

func (p *Common) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Common) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppCommonInit(engine.GIN)
	return nil
}

func (p *Common) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Common) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Common) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}
