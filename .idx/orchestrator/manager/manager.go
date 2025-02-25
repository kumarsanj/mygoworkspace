package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"cube/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDB        map[string][]*task.Task
	EventDB       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update Tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
