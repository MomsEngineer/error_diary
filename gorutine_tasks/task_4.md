# Реализация функций `repeatFn` и `take` в Go

## Задание

Напишите функции `repeatFn` и `take`:

1. **Функция `repeatFn`**:
   - Принимает на вход `context.Context` и функцию `fn`.
   - Бесконечно вызывает функцию `fn` и записывает результат ее работы в возвращаемый канал.
   - Должна завершаться при отмене контекста.

2. **Функция `take`**:
   - Принимает на вход `context.Context`, канал `in` и число `num`.
   - Читает не более чем `num` значений из канала `in`, пока канал открыт.
   - Записывает значения в возвращаемый канал.
   - Должна завершаться при отмене контекста.

## Код

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
)

func main() {
	// Создание контекста
	ctx := context.Background()

	// Построение конвейера: бесконечный вызов функции -> ограничение количества значений
	randFn := func() interface{} { return rand.Int() }
	pipeline := take(ctx, repeatFn(ctx, randFn), 3)

	// Чтение и вывод данных из конвейера
	var res []interface{}
	for num := range pipeline {
		res = append(res, num)
	}

	if len(res) != 3 {
		panic("wrong code")
	}
	fmt.Println(res)
}

// Функция repeatFn
func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	// напишите ваш код здесь
}

// Функция take
func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	// напишите ваш код здесь
}
```

# Решение

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
)

func main() {
	// Создание контекста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Построение конвейера: бесконечный вызов функции -> ограничение количества значений
	randFn := func() interface{} { return rand.Int() }
	pipeline := take(ctx, repeatFn(ctx, randFn), 3)

	// Чтение и вывод данных из конвейера
	var res []interface{}
	for num := range pipeline {
		res = append(res, num)
	}

	if len(res) != 3 {
		panic("wrong code")
	}
	fmt.Println(res)
}

// Функция repeatFn
func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
				case <- ctx.Done():
					return
				case out <- fn():
			}
		}
	}()

	return out
}

// Функция take
func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < numl; i++ {
			select {
				case <-ctx.Done():
					return
				case v, ok := <- in:
					if !ok {
						return
					}
					out <- v
			}
		}
	}()

	return out
}
```
