package log

import (
	"log"
)

func Info(fmt string, args ...interface{}) {
	log.Printf("[INFO] "+fmt+"\r\n", args...)
}

func Err(fmt string, args ...interface{}) {
	log.Printf("[ERROR] "+fmt+"\r\n", args...)
}
