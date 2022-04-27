package lesson1

import (
	"cli-tutor/src/lesson"
	"cli-tutor/src/styler"
	"cli-tutor/src/task"
	"fmt"
	"os"
	"os/user"
)


func GetLesson() lesson.Lesson {
    user, _ := user.Current()
    user_styled := styler.Style(user.Username, "tip")
    host, _ := os.Hostname()
    host_styled := styler.Style(host, "tip")
    cwd, _ := os.Getwd()
    cwd_styled := styler.Style(cwd, "tip")


    helpStep := fmt.Sprintf(
`    Welcome to the command line, this may look intimidating but this
    tutorial will introduce you to the concepts you need to be familiar
    with to thrive at the command line.

    When you are ready enter "%s" or "%s" to procede to the next task.`,
    styler.Style("next",  "tip"), styler.Style("n",  "tip"))



    promptStep := fmt.Sprintf(
`    Welcome to the command line, this may look intimidating but this tutorial
    will introduce you to the conceptd you need to be familiar with to thrive
    at the command line. 
    
    This funny looking line below you is called the "prompt".
    
    It displays information about the: 
    
    current user = %s <-- You!
    host machine = %s <-- Your machine
    the current directory you are in = %s <-- Where on you are on your machine.
`, user_styled, host_styled, cwd_styled)
    
    cmdStep := fmt.Sprintf(
`    Try it out for yourself!
    type %s and look at the result
`, styler.Style("whoami", "tip"))

    

    lesson1 := new(lesson.Lesson)
    lesson1.Name = "What is textual interaction?"
    lesson1.Vocabulary = []string{"pwd", "ls", "cd", "whoami", "uname", "echo", "curl", "man", "clear", "less", "vim"}
    lesson1.Description = "This lesson will teach you about the basics of command line interaction"
    lesson1.Tasks = []task.Task{
        {
            Description: helpStep,
        },
        {
            Description: promptStep,
        },
        {
            Description: cmdStep,
        },
        {
            Description: "Last spam",

        },
    }
    return *lesson1
}




