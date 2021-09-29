package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {

	//go tool trace
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//------------------------

	start := time.Now()
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//lines := linesPool.Get().([]string)[:0]
	//var lines string
	scanner := bufio.NewScanner(file)

	//scanner.Scan()
	for scanner.Scan() {

		//lines = append(lines, scanner.Text())
		lines := scanner.Text()
		fmt.Println(lines)
		//willScan := scanner.Scan()
		// if !willScan {
		// 	break
		// }
	}

	//fmt.Println("Time Taken:", time.Since(start).Milliseconds())
	fmt.Println("Time Taken:", time.Since(start).Microseconds(), time.Since(start).Milliseconds(), time.Since(start).Seconds())
	//fmt.Println("Time Taken:", time.Since(start).Nanoseconds())
}
