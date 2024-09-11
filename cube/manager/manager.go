package manager

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hieutdle/Himmel/cube/structures"
	"github.com/hieutdle/Himmel/cube/task"
)

type Manager struct {
	Pending       structures.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
