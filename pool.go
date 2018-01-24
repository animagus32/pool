package pool

import (
	"sync"

	"golang.org/x/net/context"
)

type Pool struct {
	Queue chan func()
	Limit int
	Ctx   context.Context
	wg    sync.WaitGroup
}

func GetPool(ctx context.Context, limit int) *Pool {

	p := Pool{
		Limit: limit,
		Ctx:   ctx,
	}
	p.Queue = make(chan func(), limit)

	return &p
}

func (p *Pool) Add(f func()) {
	p.Queue <- f
}

// 加入日志
func (p *Pool) Run() {
	for i := 0; i < p.Limit; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case <-p.Ctx.Done():
					return
				case f := <-p.Queue:
					f()
				}
			}
		}()
	}
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
