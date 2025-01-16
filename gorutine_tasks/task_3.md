# Реализация функций `generator` и `squarer` в Go

## Задание

Напишите функции `generator` и `squarer`:

1. **Функция `generator`**:
   - Принимает на вход `context.Context` и слайс целых чисел.
   - Последовательно записывает элементы слайса в возвращаемый канал.
   - Должна завершаться при отмене контекста.

2. **Функция `squarer`**:
   - Принимает на вход `context.Context` и канал целых чисел.
   - Последовательно читает числа из канала, возводит их в квадрат и записывает в возвращаемый канал.
   - Также должна завершаться при отмене контекста.

## Код

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// Создание контекста
	ctx := context.Background()

	// Построение конвейера: генерация чисел -> возведение в квадрат
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	// Чтение и вывод данных из конвейера
	for x := range pipeline {
		fmt.Println(x)
	}
}

// Функция generator
func generator(ctx context.Context, in ...int) <-chan int {
	// напишите ваш код здесь
}

// Функция squarer
func squarer(ctx context.Context, in <-chan int) <-chan int {
	// напишите ваш код здесь
}
```

# Решение

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// Создание контекста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Построение конвейера: генерация чисел -> возведение в квадрат
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	// Чтение и вывод данных из конвейера
	for x := range pipeline {
		fmt.Println(x)
	}
}

// Функция generator
func generator(ctx context.Context, in ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range in {
			select {
				case <- ctx.Done():
					return
				default:
					out <- i
			}
		}
	}()

	return out
}

// Функция squarer
func squarer(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			select {
				case <- ctx.Done():
					return
				default:
					out <- i * i
			}
		}
	}()

	return out
}
```