package logger

import (
	"log"
	"os"
	"strings"
	"time"

	helpers "github.com/nerdynz/helpers"
)

func LogPath() (string, error) {
	logPath := os.Getenv("LOG_FOLDER")
	if logPath == "" {
		logPath = os.Getenv("ATTACHMENTS_FOLDER") + "logs/"
	}
	if logPath == "" {
		logPath = "./attachments/logs/"
	}
	time := time.Now()
	fullpath := logPath + time.Format("012006") + "/"
	if _, fullErr := os.Stat(fullpath); os.IsNotExist(fullErr) {
		err := os.MkdirAll(fullpath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return fullpath, nil
}

func Log(filekey string, msg string) {
	logpath, _ := LogPath()
	f, err := os.OpenFile(logpath+strings.ToLower(helpers.KebabCase(filekey)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// dont care logging shouldn't break anything
		return
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(msg)
}
