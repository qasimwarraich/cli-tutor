package lesson

import "cli-tutor/src/task"


type Lesson struct {
    name string
    vocabulary []string
    description string
    tasks []task.Task
}
