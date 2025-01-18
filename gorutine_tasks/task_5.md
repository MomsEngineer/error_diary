# Реализация worker pool в Go

Необходимо написать worker pool: нужно выполнить параллельно `numJobs` заданий, используя `numWorkers` горутин, которые запущены единожды за время выполнения программы. Для этого напишите функции `worker` и `main`.

## Функция `worker`
- На вход получает функцию для выполнения `f`, канал для получения аргументов `jobs` и канал для записи результатов `results`.
- Читает из `jobs` и записывает результат выполнения `f(job)` в `results`.

## Функция `main`
- Запускает функцию `worker` в `numWorkers` горутинах.
- В качестве первого аргумента `worker` использует функцию `multiplier`.
- Пишет числа от 1 до `numJobs` в канал `jobs`.
- Читает и выводит полученные значения из канала `results`, параллельно работе воркеров.

```go
package main

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	// напишите ваш код здесь
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}

	// напишите ваш код здесь
}
```

# Решение
```go
package main

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- f(j)
	}
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}

	multiplier := func(x int) int {
		return x * 10
	}

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			worker(multiplier, jobs, results)
		}()
	}

	done := make(chan struct{})
	go func() {
		for i := range results {
			fmt.Println(i)
		}
		close(done)
	}()

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	close(results)
	<- done
}
```