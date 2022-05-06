package lesson

import (
	"os"
	"text/template"
)



var temp *template.Template

func init() {
    temp = template.Must(template.New("lesson.md").Funcs(FuncMap).ParseFiles("../../lessons/lesson1.md"))
}

func ParseLesson()  {
    f, _ := os.Create("expanded.md")
    // defer os.Remove(f.Name())

    temp.Execute(f, "" ) 
    temp.Execute(os.Stdout, "" ) 
     
}

