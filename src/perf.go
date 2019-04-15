package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"errors"
	"bytes"
)

type Perf struct {}
type Cmd struct {}

func PerfExec() (err error) {

	fmt.Println("Perf...")

		p := Perf{}
		arguments := os.Args[1:]

		//err = p.exec([]string{"--help"})
		err = p.exec(arguments)
		if err != nil {
			panic(err)
		}

		return err
}

func (d Perf) exec(args []string) (errr error) {
	e:= Cmd{}
	return e.exec("/tmp/perf", args, 120)
}

func (e Cmd) exec(program string, args []string, timeoutInSec time.Duration) (err error){
	
	cmd := exec.Command(program, args...)
    
	var outputbuf, errbuf bytes.Buffer
    cmd.Stdout = &outputbuf
	cmd.Stderr = &errbuf
	
	if err := cmd.Start(); err != nil {
		fmt.Println("Cmd returning error from Start...")
		fmt.Print(err)
		return err
	}
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(timeoutInSec * time.Second)
	select {
	case <-timeout:
		cmd.Process.Kill()
		
		return errors.New("There is a problem with the request")
	case err := <-done:
	
		fmt.Println("Cmd done")
		
		if err != nil {
			fmt.Println("Cmd returned error after completion", err)
			println(outputbuf.String())
			println(errbuf.String())
			return errors.New(errbuf.String())
		}

		println(outputbuf.String())
	}

	fmt.Println("Cmd processing complete")
	return nil
}
