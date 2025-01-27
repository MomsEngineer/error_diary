# Реализация функций `produce` и `main` в Go

Напишите функции `produce` и `main`.

## Функция `produce`
- На вход получает канал `pipe`.
- Бесконечно пишет целые числа в канал `pipe`, начиная с 0.
- По сигналу от `main` должна завершать работу.
- При завершении должна заснуть на 3 секунды, после чего напечатать "produce finished".

## Функция `main`
- Должна создать канал `pipe`.
- Запустить `produceCount` функций `produce` и начать читать из канала `pipe`, печатая каждое число.
- При получении числа `produceStop` из `pipe` должна перестать читать новые числа из канала и должна отправить сигнал в функции `produce`, завершающий их работу.
- Должна дождаться всех сообщений "produce finished" и напечатать "main finished".

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	produceCount = 3
	produceStop  = 10
)

func produce(pipe chan<- int) { // допускается добавить доп. аргументы
	// напишите свой код здесь
}

func main() {
	// напишите свой код здесь
}
```

## Последние 4 строки вывода программы:
produce finished
produce finished
produce finished
main finished

# Решение
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	produceCount = 3
	produceStop  = 10
)

func produce(ctx context.Context, wg *sync.WaitGroup, pipe chan<- int) {
	defer wg.Done()
	i := 0
	for {
		select {
			case <- ctx.Done():
				time.Sleep(3*time.Second)
				fmt.Println("produce finished")
				return
			case pipe <- i:
				i++
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pipe := make(chan int)

	produceCount := 3
	produceStop := 32

	wg := &sync.WaitGroup{}
	for i := 0; i < produceCount; i++ {
		wg.Add(1)
		go produce(ctx, wg, pipe)
	}

	for i := range pipe {
		fmt.Println(i)
		if i == produceStop {
			cancel()
			break
		}
	}

	wg.Wait()
	close(pipe)

	fmt.Println("main finished")
}
```