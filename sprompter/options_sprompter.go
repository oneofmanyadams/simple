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
	var real_key int
	for i := 1; real_key < len(options); i++ {
		if _, exists := options[i]; exists {
			prompt_string = prompt_string + "	"+strconv.Itoa(i)+": "+options[i]+"\n"
			real_key = real_key + 1
		}		
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