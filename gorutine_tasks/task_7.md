# Реализация функции `getResults` в Go

Необходимо написать функцию `getResults`, которая запускает конкурентный поиск для каждого набора реплик из `replicaKinds`, используя `getFirstResult`, и возвращает результат для каждого набора реплик.

## Функция `getFirstResult`
- Принимает на вход `context.Context` и `replicas`.
- Возвращает первый доступный результат из реплик.

## Функция `getResults`
- Принимает на вход `context.Context` и массив `replicaKinds`.
- Запускает конкурентный поиск для каждого набора реплик из `replicaKinds`, используя `getFirstResult`.
- Возвращает результат для каждого набора реплик.

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	msg string
	err error
}

type search func() *result
type replicas []search

func fakeSearch(kind string) search {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	// напишите ваш код здесь
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	// напишите ваш код здесь
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	replicaKinds := []replicas{
		replicas{fakeSearch("web1"), fakeSearch("web2")},
		replicas{fakeSearch("image1"), fakeSearch("image2")},
		replicas{fakeSearch("video1"), fakeSearch("video2")},
	}
	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}
```

# Решение
```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type result struct {
	msg string
	err error
}

type search func() *result
type replicas []search

func fakeSearch(kind string) search {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	if len(replicas) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan *result)
	defer close(c)

	for _, replica := range replicas {
		go func(replica search) {
			select {
				case <-ctx.Done():
				case c <- replica():
			}
		}(replica)
	}

	select {
		case res := <- c:
			return res
		case <- ctx.Done():
			return &result{err: ctx.Err()}
	}
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	c := make(chan *result)
	wg := sync.WaitGroup{}

	for _, searchKind := range replicaKinds {
		wg.Add(1)
		go func(searchKind replicas) {
			defer wg.Done()
			c <- getFirstResult(ctx, searchKind)
		}(searchKind)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	res := []*result{}
	for {
		select {
			case ans, ok := <- c:
				if !ok {
					return res
				}
				res = append(res, ans)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	replicaKinds := []replicas{
		replicas{fakeSearch("web1"), fakeSearch("web2")},
		replicas{fakeSearch("image1"), fakeSearch("image2")},
		replicas{fakeSearch("video1"), fakeSearch("video2")},
	}
	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}