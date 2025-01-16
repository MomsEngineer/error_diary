# Реализация функций `merge` и `fillChan` в Go

Данная задача требует реализации двух функций: `fillChan` и `merge`.

## Функция `fillChan`
- На вход принимает целое число `n`.
- Возвращает канал, в который записываются числа от `0` до `n-1`.

## Функция `merge`
- На вход принимает массив каналов `cs`.
- Возвращает канал, который:
  - Параллельно читает данные из каждого канала из `cs`.
  - Записывает все прочитанные значения в возвращаемый канал.

```go
package main

import (
	"fmt"
	"sync"
)

// merge - соединяет каналы в один
func merge(cs ...<-chan int) <-chan int {
//напишите ваш код здесь
}

// fillChan - заполняет канал числами от 0 до n-1
func fillChan(n int) <-chan int {
//напишите ваш код здесь
}

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
```
## Пример
Функция `main` должна объединить результаты работы `fillChan` для каналов `a`, `b` и `c` и вывести их значения:

### Ожидаемый результат:
```text
0
1
0
1
2
0
1
2
3
```

# Решение

```go
package main

import (
	"fmt"
	"sync"
)

// merge - соединяет каналы в один
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, c := range cs {
		wg.Add(1)
		go func(r <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// fillChan - заполняет канал числами от 0 до n-1
func fillChan(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
```