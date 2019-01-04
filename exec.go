package main

import (
	"fmt"
	"log"
	"os/exec"
	"bufio"
	"strings"
	"regexp"
)

func main() {

    // Pipe a directory listing through sort.
	cmd1 := exec.Command("ls", "-la")
	args := []string{"-la","*.go"}
	cmd3 := exec.Command("ls", args...)
	cmd2 := exec.Command("grep", "go")

	test, err := cmd3.CombinedOutput()
	scanner2 := bufio.NewScanner(strings.NewReader(string(test)))
	for scanner2.Scan() {
		re := regexp.MustCompile(" +")
                s := re.Split(scanner2.Text(),-1)
                fmt.Println(s[8])
	}

    // Get the pipe of Stdout from cmd1 and assign it
    // to the Stdin of cmd2.
	pipe, err := cmd1.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd2.Stdin = pipe


    // Start() cmd1, so we don't block on it.
	err = cmd1.Start()
	if err != nil {
		log.Fatal(err)
	}

    // Run Output() on cmd2 to capture the output.
	output, err := cmd2.Output()
	if err != nil {
		log.Fatal(err)
	}

    // Voila!
	scanner := bufio.NewScanner(strings.NewReader(string(output))) // f is the *os.File
	for scanner.Scan() {
		re := regexp.MustCompile(" +")
		s := re.Split(scanner.Text(),-1)
		fmt.Println(s[8]) // Println will add back the final '\n'
	}

}
