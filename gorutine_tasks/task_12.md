# Реализация интерфейса `waiter` в Go

Есть интерфейс `waiter`, который должен:
1. Параллельно запускать переданные в `run` функции с указанным контекстом.
2. Количество параллельных запусков определяется параметром `maxParallel` при создании `waiter` через `newGroupWait`.
3. Возвращать ошибку из `wait`, если хотя бы одна функция из `run` вернула ее.
4. Возвращать комбинацию ошибок от вызовов `run`, если несколько задач завершились с ошибками (можно использовать `errors.Join`).

Реализуйте поля и методы структуры `waitGroup` для интерфейса `waiter`.

## Интерфейс `waiter`
- Параллельно запускает переданные в `run` функции с указанным контекстом.
- Количество параллельных запусков определяется параметром `maxParallel`.
- Возвращает ошибку из `wait`, если хотя бы одна функция из `run` вернула ее.
- Возвращает комбинацию ошибок от вызовов `run`, если несколько задач завершились с ошибками.

```go
package main

import (
	"context"
	"errors"
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

type waitGroup struct {
	// напишите ваш код здесь
}

func (g *waitGroup) wait() error {
	// напишите ваш код здесь
}

func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	go func() {
		defer wg.Done()
		err fn(ctx)
	}()
}

func newGroupWait(maxParallel int) waiter {
	// напишите ваш код здесь
}

func main() {
	g := newGroupWait(2)
	ctx := context.Background()
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")
	g.run(ctx, func(ctx context.Context) error {
		return nil
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})
	err := g.wait()
	if err != nil {
		fmt.Println("Errors:", err)
	}
}
```

# Решение
```go
package main

import (
	"context"
	"errors"
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

type pair struct {
	ctx context.Context
	fn func(ctx context.Context) error
}

type waitGroup struct {
	workersWg sync.WaitGroup
	queue chan pair
	queueWg sync.WaitGroup
	mu sync.Mutex
	err error
}

func (g *waitGroup) wait() error {
	g.queueWg.Wait()
	close(g.queue)
	g.workersWg.Wait()

	return g.err
}

func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	g.queueWg.Add(1)
	go func() {
		defer g.queueWg.Done()
		g.queue <- pair{ctx: ctx, fn: fn}
	}()
}

func newGroupWait(maxParallel int) waiter {
	g := &waitGroup {
		workersWg: sync.WaitGroup{},
		queue: make(chan pair),
		queueWg: sync.WaitGroup{},
	}

	g.workersWg.Add(maxParallel)
	for i := 0; i < maxParallel; i++ {
		go func() {
			defer g.workersWg.Done()
			for p := range g.queue {
				err := p.fn(p.ctx)
				if err != nil {
					g.mu.Lock()
					g.err = errors.Join(g.err, err)
					g.mu.Unlock()
				}
			}
		}()
	}

	return g
}

func main() {
	g := newGroupWait(2)
	ctx := context.Background()
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")
	g.run(ctx, func(ctx context.Context) error {
		return nil
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})
	err := g.wait()
	if err != nil {
		fmt.Println("Errors:", err)
	}
}
```