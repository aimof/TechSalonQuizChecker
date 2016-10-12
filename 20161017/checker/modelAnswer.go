package main

type modelCase struct {
	input  string
	output string
}

func defModelAnswer() []modelCase {
	modelAnswer := []modelCase{
		modelCase{input: "8 5", output: "5\n"},
		modelCase{input: "17 2", output: "4\n"},
		modelCase{input: "47 5", output: "30\n"},
		modelCase{input: "6 19", output: "0\n"},
		modelCase{input: "1000 1000", output: "125000\n"},
		modelCase{input: "1 21", output: "0\n"},
		modelCase{input: "500 1000", output: "62000\n"},
		modelCase{input: "7 10", output: "10\n"},
		modelCase{input: "8 380", output: "380\n"},
		modelCase{input: "6 490", output: "0\n"},
	}
	return modelAnswer
}
