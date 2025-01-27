# Реализация методов у структуры Group (аналог sync.WaitGroup) в Go

Реализуйте методы у структуры Group (аналог sync.WaitGroup), чтобы код не приводил к панике.

## Структура Group
- `New(size int) *Group`: Конструктор для создания новой группы заданного размера.
- `Done()`: Метод для завершения задачи.
- `Wait()`: Метод для ожидания завершения всех задач.

```go
package main

import (
	"reflect"
	"sort"
	"sync"
)

type Group struct {
	c    chan struct{}
	size int
}

func New(size int) *Group {
	// напишите ваш код здесь
}

func (s *Group) Done() {
	// напишите ваш код здесь
}

func (s *Group) Wait() {
	// напишите ваш код здесь
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)
	var res []int
	var mu sync.Mutex
	group := New(n)
	for _, num := range numbers {
		go func(num int) {
			defer group.Done()
			mu.Lock()
			res = append(res, num)
			mu.Unlock()
		}(num)
	}
	group.Wait()
	sort.IntSlice(res).Sort()
	if !reflect.DeepEqual(res, numbers) {
		panic("wrong code")
	}
}
```

# Решение
```go
package main

import (
	"reflect"
	"sort"
	"sync"
)

type Group struct {
	c    chan struct{}
	size int
}

func New(size int) *Group {
	return &Group{
		c: make(chan struct{}, size),
		size: size,
	}
}

func (s *Group) Done() {
	s.c <- struct{}{}
}

func (s *Group) Wait() {
	for i := 0; i < s.size; i++ {
		<- s.c
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)
	var res []int
	var mu sync.Mutex
	group := New(n)
	for _, num := range numbers {
		go func(num int) {
			defer group.Done()
			mu.Lock()
			res = append(res, num)
			mu.Unlock()
		}(num)
	}
	group.Wait()
	sort.IntSlice(res).Sort()
	if !reflect.DeepEqual(res, numbers) {
		panic("wrong code")
	}
}
```