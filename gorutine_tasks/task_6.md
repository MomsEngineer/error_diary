# Реализация функции `mergeSorted` в Go

Необходимо написать функцию `mergeSorted`, которая принимает на вход два "отсортированных" канала и возвращает результирующий отсортированный канал.

## Функция `mergeSorted`
- Принимает на вход два отсортированных канала `a` и `b`.
- Возвращает результирующий отсортированный канал.

```go
package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	// напишите ваш код здесь
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)

	c := mergeSorted(a, b)

	for val := range c {
		fmt.Printf("%d ", val)
	}
}
```

### Ожидаемый результат:
```text
-1 1 2 4 4 5
```

# Решение

```go
package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		aVal, aOK := <-a
		bVal, bOK := <-b

		for aOK && bOK {
			if aVal < bVal {
				out <- aVal
				aVal, aOK = <-a
			} else {
				out <- bVal
				bVal, bOK = <-b
			}
		}

		for aOK {
			out <- aVal
			aVal, aOK = <-a
		}

		for bOK {
			out <- bVal
			bVal, bOK = <-b
		}
	}()

	return out
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)
	for val := range c {
		fmt.Printf("%d ", val)
	}
}
```