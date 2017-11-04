package sprompter

type Prompt struct {
	Question string
	Options map[int]string
	Result string
	OptionResult int
}

func NewPrompt(question string) (new_prompt Prompt) {
	new_prompt.Question = question
	new_prompt.Options = make(map[int]string)
	return
}

func (prompt *Prompt) AddOption(key int, option string) {
	prompt.Options[key] = option
}

func (prompt *Prompt) GetInput() {
	if len(prompt.Options) > 0 {
		prompt.OptionResult = PromptUserWithOptions(prompt.Question, prompt.Options)
		if prompt.OptionResult != 0 {
			prompt.Result = prompt.Options[prompt.OptionResult]
		}
	} else {
		prompt.Result = PromptUser(prompt.Question)
		
	}
}