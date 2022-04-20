package lesson

import "cli-tutor/src/task"


type Lesson struct {
    Name string
    Vocabulary []string
    Description string
    Tasks []task.Task
}
