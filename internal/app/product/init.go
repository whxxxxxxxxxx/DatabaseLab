package product

import (
	"DatabaseLab/internal/app"
	"DatabaseLab/internal/app/product/dao"
	"DatabaseLab/internal/app/product/router"
	"DatabaseLab/kernel"
	"context"
	"sync"
)

type (
	Product struct {
		Name string
		app.UnimplementedModule
	}
)

func (p *Product) Info() string {
	return p.Name
}

func (p *Product) PreInit(engine *kernel.Engine) error {
	dao.InitOP(engine.OpenGauss.DB)
	return nil
}

func (p *Product) Init(*kernel.Engine) error {
	return nil
}

func (p *Product) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Product) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppProductInit(engine.GIN)
	return nil
}

func (p *Product) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Product) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Product) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}
