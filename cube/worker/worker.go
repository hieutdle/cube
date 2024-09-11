package worker

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hieutdle/Himmel/cube/structures"
	"github.com/hieutdle/Himmel/cube/task"
)

type Worker struct {
	Name      string
	Queue     structures.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("I Will stop a task")
}
