package logger

import (
	"errors"
	"log"
	"os"
	"strings"

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

	if logsFolder := os.Getenv("LOGS_FOLDER"); logsFolder != "" {
		logPath = logsFolder
	}

	// time := time.Now()
	// fullpath := logPath + time.Format("012006") + "/"
	if _, fullErr := os.Stat(logPath); os.IsNotExist(fullErr) {
		return "", errors.New("Log folder is unavailable: " + fullErr.Error())
	}
	return logPath, nil
}

func Log(filekey string, msg string) {
	logpath, err := LogPath()
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(logpath+strings.ToLower(helpers.KebabCase(filekey)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// dont care logging shouldn't break anything
		return
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(msg)
}
