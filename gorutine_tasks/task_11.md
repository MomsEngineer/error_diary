# Реализация функции `bridge` в Go

Напишите функцию `bridge`, которая принимает на вход канал каналов и возвращает канал, который направляет данные из всех входных каналов, пока они открыты и контекст не отменен.

## Функция `bridge`
- Принимает на вход `context.Context` и канал каналов `ins`.
- Возвращает канал, который направляет данные из всех входных каналов, пока они открыты и контекст не отменен.

```go
package main

import (
	"context"
	"reflect"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}

	var res []interface{}
	for v := range bridge(context.Background(), genVals()) {
		res = append(res, v)
	}
	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}

func bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{} {
	// напишите ваш код здесь
}

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return out
}
```

# Решение

```go
package main

import (
	"context"
	"reflect"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}

	var res []interface{}
	for v := range bridge(context.Background(), genVals()) {
		res = append(res, v)
	}
	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}

func bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			var stream <-chan interface{}
			select {
			case <-ctx.Done():
				return
			case s, ok := <-ins:
				if !ok {
					return
				}
				stream = s
			}

			for val := range orDone(ctx, stream) {
				select {
				case out <- val:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return out
}
```