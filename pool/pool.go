package pool

import (
	"fmt"
	"sync"
)

//Executor interface to force implementing Run() method
type Executor interface {
	Run()
}

//GoroutinePool struct to manage goroutines
type GoroutinePool struct {
	queue chan work
	wg    sync.WaitGroup
}

type work struct {
	fn Executor
}

//NewGoroutinePool set number of workers
func NewGoroutinePool(workerSize int) *GoroutinePool{
	gp := &GoroutinePool{
		queue: make(chan work),
	}

	gp.AddWorkers(workerSize)
	return gp
}

//Close to close the queue
func (gp *GoroutinePool) Close(){
	close(gp.queue)
	gp.wg.Wait()
}

//ScheduleWorks method to schedule functions
func (gp *GoroutinePool) ScheduleWorks(fn Executor){
	gp.queue <- work{fn}
}

//AddWorkers execute queue functions
func (gp *GoroutinePool) AddWorkers(numWorkers int){
	gp.wg.Add(numWorkers)
	for i:= 0; i < numWorkers; i++{
		go func(workerID int){
			count := 0
			for job := range gp.queue{
				job.fn.Run()
				count++
			}
			fmt.Println(fmt.Sprintf("Executor %d executed %d requests", workerID, count))
			gp.wg.Done()
		}(i)
	}
}