# Реализация функции `executeTaskWithTimeout` в Go

Есть функция `executeTask`, которая во время исполнения может зависнуть на неопределенно долгое время. Реализуйте функцию-обертку `executeTaskWithTimeout`, которая:
- Исполняет `executeTask`.
- Принимает аргументом контекст.
- Завершается либо в результате исполнения `executeTask`, либо в результате отмены контекста. В последнем случае возвращает ошибку контекста.

## Функция `executeTaskWithTimeout`
- Принимает на вход `context.Context`.
- Исполняет `executeTask`.
- Завершается либо в результате исполнения `executeTask`, либо в результате отмены контекста.

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	// напишите ваш код здесь
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
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
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		defer close(done)
		executeTask()
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
```