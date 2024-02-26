package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	var (
		sleep   uint
		command string
	)
	flag.StringVar(&command, "c", "", "Command to execute")
	flag.UintVar(&sleep, "t", 60, "Sleep time in seconds")
	flag.Parse()

	if command == "" {
		flag.Usage()
		os.Exit(1)
	}

	for {
		cmd := exec.Command(command)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing %v command: %v\n", command, err)
		}
		fmt.Println(string(output))
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}
