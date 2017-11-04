package sprompter

import (
	"fmt"
	"strconv"
)

// resturns 0 if no valid option is selected
//		-Maybe change that to returning an error?
//		-Needs to make sure options start with 1
//			-And don't skip any numbers

func PromptUserWithOptions(question string, options map[int]string) (selected_option int) {
	// build prompt string
	var prompt_string string	
	prompt_string = question+"\n"
	for i := 1; i <= len(options); i++ {
		prompt_string = prompt_string + "	"+strconv.Itoa(i)+": "+options[i]+"\n"
	}
	
	// prompt user for input
	user_input := PromptUser(prompt_string)
	
	// verify user input is a valid option
	user_selected_option, _ := strconv.Atoi(user_input)
	if _, exists := options[user_selected_option]; exists {
		selected_option = user_selected_option
	} else {
		fmt.Println("Option Provided is not valid!")
		selected_option = 0
	}
	
	return 
}