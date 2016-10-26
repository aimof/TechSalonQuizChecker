package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	rowCmd := readCommand()
	modelAnswer := defModelAnswer()
	correctAnswer := checkAnswer10times(rowCmd, modelAnswer)
	for i, correct := range correctAnswer {
		fmt.Printf("%d 回目のテストは %t です。\n", i+1, correct)
	}
}

func readCommand() string {
	filename := "command.txt"
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	var rowCmd string
	for scanner.Scan() {
		rowCmd = scanner.Text()
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rowCmd
}

func checkAnswer10times(rowCmd string, modelAnswer []modelCase) []bool {
	var correctAnswer []bool
	for _, InputOutput := range modelAnswer {
		correctAnswer = append(correctAnswer, checkAnswer(InputOutput.input, InputOutput.output, rowCmd))
	}
	return correctAnswer
}

func getCmd() string {
	var rowCmd string
	fmt.Scan(&rowCmd)
	return rowCmd
}

func checkAnswer(input, output, rowCmd string) bool {
	cmd := exec.Command(rowCmd)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(stdin, input)
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	if string(out) == output {
		return true
	}
	return false
}
