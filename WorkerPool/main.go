package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker  int
	wg         *sync.WaitGroup
	queuedTask chan func()
}

func main() {
	wp := workerPool{
		maxWorker:  3,
		wg:         &sync.WaitGroup{},
		queuedTask: make(chan func()),
	}

	go func() {
		for {
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()
	wp.Run()
	type result struct {
		id    int
		value int
	}

	totalTask := 5
	wp.wg.Add(5)
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", id)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d", res.id, res.value)
	}

	wp.wg.Wait()
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTask {
				task()
				wp.wg.Done()
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTask <- task
}
