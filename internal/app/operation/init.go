package operation

import (
	"DatabaseLab/internal/app"
	"DatabaseLab/internal/app/operation/dao"
	"DatabaseLab/internal/app/operation/router"
	"DatabaseLab/kernel"
	"context"
	"sync"
)

type (
	Operation struct {
		Name string
		app.UnimplementedModule
	}
)

func (p *Operation) Info() string {
	return p.Name
}

func (p *Operation) PreInit(engine *kernel.Engine) error {
	dao.InitOP(engine.OpenGauss.DB)
	return nil
}

func (p *Operation) Init(*kernel.Engine) error {
	return nil
}

func (p *Operation) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Operation) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppOperationInit(engine.GIN)
	return nil
}

func (p *Operation) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Operation) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Operation) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}
