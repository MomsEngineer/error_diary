# Реализация методов структуры `ringBuffer` в Go

Реализуйте методы структуры `ringBuffer` так, чтобы `main` отрабатывал без паники.

## Структура `ringBuffer`
- `newRingBuffer(size int) *ringBuffer`: Конструктор для создания нового кольцевого буфера заданного размера.
- `write(v int)`: Метод для записи значения в буфер.
- `close()`: Метод для закрытия буфера.
- `read() (v int, ok bool)`: Метод для чтения значения из буфера.

```go
package main

import (
	"fmt"
	"reflect"
)

type ringBuffer struct {
	c chan int
}

func newRingBuffer(size int) *ringBuffer {
	// напишите ваш код
}

func (b *ringBuffer) write(v int) {
	// напишите ваш код
}

func (b *ringBuffer) close() {
	// напишите ваш код
}

func (b *ringBuffer) read() (v int, ok bool) {
	// напишите ваш код
}

func main() {
	buff := newRingBuffer(3)
	for i := 1; i <= 6; i++ {
		buff.write(i)
	}
	buff.close()
	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}
	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
```

# Решение
```go
package main

import (
	"fmt"
	"reflect"
)

type ringBuffer struct {
	c chan int
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		c: make(chan int size),
	}
}

func (b *ringBuffer) write(v int) {
	select {
	case b.c <- v:
	default:
		<-b.c
		b.write(v)
	}
}

func (b *ringBuffer) close() {
	close(b.c)
}

func (b *ringBuffer) read() (v int, ok bool) {
	ok, v = <- b.c
}

func main() {
	buff := newRingBuffer(3)
	for i := 1; i <= 6; i++ {
		buff.write(i)
	}
	buff.close()
	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}
	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
```