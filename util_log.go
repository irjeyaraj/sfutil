/*
 * util_log.go
 *
 * Copyright 2022 Immanuel Jeyaraj
 *
 * Author: Immanuel Jeyaraj <irj@sefier.com>
 *
 * Created date: 3 June 2019
 */

package sfutil

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var file = os.Stderr

var (
	LogTrace   *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

func LogInit(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	LogTrace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogInfo = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogWarning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogError = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func initLog(logPath string, appName string, logLevel int) {

	currentTime := time.Now()
	fileNameTime := currentTime.Format("2006-01-02_15-04-05")
	startTime := currentTime.Format("2006-01-02 15:04:05")

	logFilePath := logPath + Path_separator() + appName + "_" + fileNameTime + ".log"

	file, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0766)

	logRaw("Starting Log Server for Sedona\nStart time: " + startTime + "\n==================================\n")

	switch logLevel {
	case 0:
		LogInit(file, file, file, file)
	case 1:
		LogInit(ioutil.Discard, file, file, file)
	case 2:
		LogInit(ioutil.Discard, ioutil.Discard, file, file)
	default:
		LogInit(ioutil.Discard, ioutil.Discard, ioutil.Discard, file)
	}

}

func logRaw(msg string) {
	file.WriteString(msg)
}
