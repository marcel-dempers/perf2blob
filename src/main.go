package main

import (
	"fmt"
)

func main() {
	fmt.Println("running perf")
	err := PerfExec()

	if err != nil {
		panic(err)
	}

}
