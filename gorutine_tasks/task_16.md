# Реализация функции `inc` в Go

Напишите функцию `inc`, которая принимает на вход канал, читает из него значения и записывает эти значения, увеличенные на единицу, в возвращаемый канал. Дополните функцию `main` созданием цепочки каналов, используя `inc`, так, чтобы программа завершалась без паники.

## Функция `inc`
- Принимает на вход канал `in`.
- Читает значения из канала `in` и записывает их, увеличенные на единицу, в возвращаемый канал.

```go
package main

func main() {
	first := make(chan int)
	last := make(<-chan int)
	n := 10
	//
	// напишите ваш код здесь
	//
	first <- 0
	close(first)
	if n != <-last {
		panic("wrong code")
	}
}

func inc(in <-chan int) <-chan int {
	// напишите ваш код здесь
}
```

# Решение
```go
package main

func main() {
	first := make(chan int)
	last := make(<-chan int)
	n := 10

	last = inc(first)
	for i := 1; i < n; i++ {
		last = inc(last)
	}

	first <- 0
	close(first)
	if n != <-last {
		panic("wrong code")
	}
}

func inc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- (i + 1)
		}
	}()

	return out
}
```
