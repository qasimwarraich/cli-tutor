package lesson

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

//go:embed lessons
var EmbeddedFS embed.FS

func LoadLesson(selected string) Lesson {
	temp := template.Must(template.New(selected).Funcs(FuncMap).ParseFS(EmbeddedFS, "lessons/"+selected))
	f, _ := os.Create("expanded.md")
	defer os.Remove(f.Name())

	err := temp.Execute(f, "")
	if err != nil {
		log.Panic(err)
	}

	content, _ := os.ReadFile("expanded.md")
	currentLesson := ParseLesson(content)

	createExampleFile()
	return currentLesson
}

func createExampleFile() {
	testfileContent, err := fs.ReadFile(EmbeddedFS, "lessons/lesson-workspace/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	hiddenfileContent, err := fs.ReadFile(EmbeddedFS, "lessons/lesson-workspace/hiddenfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("file.txt", testfileContent, 0o644)
	ioutil.WriteFile(".hiddenfile.txt", hiddenfileContent, 0o644)
}
