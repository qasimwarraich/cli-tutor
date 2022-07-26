package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var (
	Identifier string = ""
	hostName   string
)

func InitLogging() {
	// Init Logging
	logFile, err := os.OpenFile("/tmp/tutor-log.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	hostName, _ = os.Hostname()
	log.SetPrefix(hostName + " ")
	log.SetOutput(logFile)
	log.Println("Starting new tutor session")
}

func UploadLog() {
	fileDir := "/tmp/"
	fileName := "tutor-log.txt"
	filePath := path.Join(fileDir, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		log.Println(err)
	}

	io.Copy(part, file)
	writer.Close()

	r, err := http.NewRequest("POST", "https://tutor.chistole.ch/api/upload", body)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		fmt.Println(res)
	}
}
