package lesson1

import (
    "cli-tutor/src/lesson"
    "cli-tutor/src/task"
)


func GetLesson() lesson.Lesson {

    lesson1 := new(lesson.Lesson)
    lesson1.Name = "What is textual interaction?"
    lesson1.Vocabulary = []string{"pwd", "ls", "cd", "whoami", "uname", "echo", "curl", "man", "clear", "less", "vim"}
    lesson1.Description = "This lesson will teach you about the basics of command line interaction"
    lesson1.Tasks = []task.Task{
        {
            Description: 
            `Welcome to the command line, this may look intimidating 
            but this tutorial will introduce you to the conceptd you need to be 
            familiar with to thrive at the command line. 

            This funny looking line below you is called the "prompt".
            `,
        },
        {
            Description: "Some spam",
        },
        {
            Description: "Some spam that will eventually be a task description.",
        },
        {
            Description: "Last spam",

        },
    }
    return *lesson1
}
