package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	var rowCmd string
	fmt.Scan(&rowCmd)
	modelAnswer := defModelAnswer()
	correctAnswer := checkAnswer10times(rowCmd, modelAnswer)
	for i, correct := range correctAnswer {
		fmt.Printf("%d 回目のテストは %t です。\n", i+1, correct)
	}
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
