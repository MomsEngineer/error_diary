# Реализация функции `Run` в Go

Напишите функцию `Run`, которая запускает конкурентное выполнение функций `fs` и дожидается их окончания. Если одна или несколько функций из `fs` завершились с ошибкой, `Run` возвращает любую из них.

## Функция `Run`
- Принимает на вход список функций `fs` (массив `fn`).
- Запускает конкурентное выполнение функций `fs`.
- Дожидается окончания выполнения всех функций.
- Возвращает любую ошибку, если одна или несколько функций завершились с ошибкой.

```go
package main

import (
	"errors"
)

type fn func() error

func main() {
	expErr := errors.New("error")
	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}
	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}

func Run(fs ...fn) error {
	// напишите ваш код здесь
}
```

# Решение
```go
package main

import (
	"errors"
)

type fn func() error

func main() {
	expErr := errors.New("error")
	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}
	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}

func Run(fs ...fn) error {
	out := make(chan error)
	wg := &sync.WaitGroup{}

	for _, f := range fs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out <- f()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for err := range out {
		if err != nil {
			go func() {
				for _ = range out {}
			}()
			return err
		}
	}

	return nil
}
```