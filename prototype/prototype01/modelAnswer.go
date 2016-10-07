package main

type modelCase struct {
	input  string
	output string
}

func defModelAnswer() []modelCase {
	modelAnswer := []modelCase{
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
		modelCase{input: "start", output: "finish"},
	}
	return modelAnswer
}
