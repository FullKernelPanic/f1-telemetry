package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriterDecorator struct {
	writer         io.Writer
	printToConsole bool
}

func Init(logFileName string, logToConsole bool) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(logWriterDecorator{f, logToConsole})
}

func (d logWriterDecorator) Write(bytes []byte) (int, error) {
	if d.printToConsole {
		fmt.Print(string(bytes))
	}

	return d.writer.Write(bytes)
}
