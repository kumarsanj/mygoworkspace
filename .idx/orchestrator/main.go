package main

import (
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"cube/task"
	"cube/worker"
)

/*
*
TODO install docker
*/
func main() {

	fmt.Println("inside main")
	t := task.Task{
		ID:          uuid.New(),
		ContainerId: "Task-1",
		Image:       "Image-1",
		Disk:        1,
		Memory:      1024,
		State:       task.Pending,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Task:      t,
		TimeStamp: time.Now(),
	}

	w := worker.Worker{
		Name:  "Worker-1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	fmt.Println("inside cube main", t, te, w)
	w.CollectStats()
	w.StartTask()

	fmt.Println("creating a test container")
	docker, dockerResult := createContainer()

	if dockerResult.Error != nil {
		fmt.Println("error creating container. printing before os.exit")
		os.Exit(1)
	}

	fmt.Println("stopping the test container")
	time.Sleep(time.Second * 5)
	_ = stopContainer(docker, dockerResult.ContainerId)

}

func createContainer() (*task.Docker, *task.DockerResult) {
	c := task.Config{
		Name:  "test-container1",
		Image: "postgres:13",
		Env: []string{
			"POSTGRES_USER=cube",
			"POSTGRES_SECRET=secret",
		},
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := task.Docker{
		Client: dc,
		Config: c,
	}

	dr := d.Run()
	if dr.Error != nil {
		fmt.Printf("%v", dr.Error)
		return nil, nil
	}

	fmt.Printf("Container %s is running with config %v", dr.ContainerId, c)

	return &d, &dr
}

func stopContainer(d *task.Docker, id string) *task.DockerResult {
	dr := d.Stop(id)

	if dr.Error != nil {
		fmt.Printf("Error stopping container %s: %v", id, dr.Error)
		return nil
	}

	fmt.Printf("container has been stopped. %s", dr.ContainerId)
	return &dr
}
