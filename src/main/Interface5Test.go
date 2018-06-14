package main

import "log"

type Task struct {
	Command string
	*log.Logger
}

func NewTask(command string, logger *log.Logger) *Task {
	return &Task{command, logger}
}

func main() {

}
