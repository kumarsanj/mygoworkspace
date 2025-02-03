package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"cube/task"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) RunTask() {
	fmt.Println("Running Task")
}

func (w *Worker) StartTask() {
	fmt.Println("Starting Task")
}

func (w *Worker) EndTask() {
	fmt.Println("Ending Task")
}

func (w *Worker) CollectStats() {
	fmt.Println("Collecting Stats")
}
