package main

import (
	"attendance"
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := attendance.NewAttendanceCmd().Execute()

	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}