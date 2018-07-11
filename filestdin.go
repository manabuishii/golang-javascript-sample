package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {

	path := "sample.txt"
	fmt.Println(path)
	var cmd *exec.Cmd
	sF, _ := os.Open(path)
	inputString, _ := ioutil.ReadAll(sF)
	//cmd.Stdin = sF
	cmd = exec.Command("wc", "-l")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(stdin, string(inputString))
	stdin.Close()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// if err = cmd.Start(); err != nil { //Use start, not run
	// 	fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	// }
	// cmd.Wait()

	if err := cmd.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}
	cmd.Wait()
	//fmt.Println(out)
	fmt.Println("fin")
}
