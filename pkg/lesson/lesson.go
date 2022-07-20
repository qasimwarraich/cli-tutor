package lesson

type Lesson struct {
	Name        string
	Vocabulary  []string
	Description string
	Tasks       []Task
}

type Task struct {
	Title       string
	Description string
	Expected    string
}
