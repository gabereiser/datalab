package log

import (
	"log"
)

func LogSetup() {
	log.SetFlags(0)
}
func Info(msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg+"\r\n", args...)
}

func Err(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg+"\r\n", args...)
}
