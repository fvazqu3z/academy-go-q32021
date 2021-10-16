package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testTask struct {
	Name string
	TaskProcessor func(...interface{})
}

func (t *testTask) Run(){
	t.TaskProcessor(t.Name)
}

func TestDispatcher(t *testing.T){
	pool := NewGoroutinePool(3)
	taskSize := 10
	taskCounter := 0

	//wait
	wg := &sync.WaitGroup{}
	wg.Add(taskSize)

	sampleStringTaskFn := func(dm ...interface{}) {
		if myInput, ok := dm[0].(string); ok{
			time.Sleep(time.Second)
			if myInput != ""{
				fmt.Printf("Finished #{%v}\n" , myInput)
			}
			taskCounter++
			wg.Done()
		}
	}

	var tasks []*testTask
	for v:=0 ; v < taskSize; v++ {
		tasks = append(tasks , &testTask{
			Name: fmt.Sprintf("task %d" , v),
			TaskProcessor: sampleStringTaskFn,
		})
	}

	for _, task := range tasks{
		pool.ScheduleWorks(task)
	}
	pool.Close()

	wg.Wait()

	assert.NotNil(t, pool)
	assert.EqualValues(t, taskCounter , taskSize)
}