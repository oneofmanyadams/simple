package sprompter

import (
	"bufio"
	"os"
	"fmt"
	"strings"
/*
	"strconv"
*/
)

func PromptUser(question string) (answer string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(question)
	fmt.Print("#:")

	user_input, _ := reader.ReadString('\n')
	
	answer = clean_up_input(user_input)

	return
}

func clean_up_input(input string) string {
	remover := strings.NewReplacer("\n","")	

	return strings.TrimSpace(remover.Replace(input))
}