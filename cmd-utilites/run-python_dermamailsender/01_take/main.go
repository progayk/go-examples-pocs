package main

import (
	"bufio"
	"context"
	"fmt"
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
	chDirErr(err)

	// create a context to kill the process after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create the command
	cmd := exec.CommandContext(
		ctx,
		"/usr/bin/zsh",
		"-c", "source bin/activate && cd appfolder && flask run")

	// get the output
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("couldnt get output", err)
	}
	defer out.Close()

	// start executing the app
	err = cmd.Start()
	if err != nil {
		log.Fatal("couldn't start", err)
	}

	// scan the output
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		b := scanner.Text()
		fmt.Print(b+"\n")
	}

	log.Println("scan is aborted")

	// wait till the executing is done
	if err = cmd.Wait(); err != nil {
		log.Fatal("couldn't wait", err)
	}

}

func chDirErr(err error)  {
	if err != nil {
		log.Fatal("no such path", err)
	}
}


