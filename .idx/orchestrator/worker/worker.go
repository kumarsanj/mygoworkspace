package worker

import (
	"errors"
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

func (w *Worker) StartTask(t *task.Task) task.DockerResult {
	fmt.Println("Starting Task")

	//create new Config using Task data
	c := t.CreateNewDockerConfigFromTask()

	//instantiate Docker and underlying Config object to start task as a Docker container
	d := t.NewDocker(c)

	//run docker container
	dr := d.Run()
	if dr.Error != nil {
		fmt.Println("error while running docker container:", dr.Error)
		t.State = task.Failed
	}

	fmt.Println("Task Started")
	w.Db[t.ID] = t
	w.TaskCount++
	t.State = task.Running
	t.ContainerId = dr.ContainerId

	return dr
}

func (w *Worker) StopTask(t *task.Task) task.DockerResult {
	fmt.Println("Stopping Task", t)

	dr := task.DockerResult{}

	//TODO: should we add panic exit for below code?
	if len(t.ContainerId) == 0 {
		dr.Error = errors.New("ContainerID is empty. Cannot stop the task")
		fmt.Println("Cannot stop the task. ContainerID is empty")
	} else {
		dr = t.Docker.Stop(t.ContainerId)
	}

	if dr.Error != nil {
		fmt.Println("Stopping task resulted in an error. ", dr.Error)
		t.State = task.Failed
	}

	fmt.Println("Task Stopped")
	w.TaskCount--
	t.State = task.Completed
	return dr
}

func (w *Worker) CollectStats() (int, int) {
	fmt.Println("Collecting Stats")
	return w.TaskCount, len(w.Db)
}
