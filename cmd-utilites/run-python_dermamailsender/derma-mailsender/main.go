package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	// go to work path
	p := path.Dir("/home/jony/Documents/dermaplast/")
	err := os.Chdir(p)
	if err != nil {
		log.Fatal("no such path", err)
	}

	// actigin seyi kapat, zorla da olsa!
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// create the command
	cmd := exec.CommandContext(
		ctx,
		"/usr/bin/zsh",
		"-c", "source bin/activate && cd appfolder && flask run")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}
