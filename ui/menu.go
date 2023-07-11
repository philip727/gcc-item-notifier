package ui

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func MainUI() string {
	// Asks the user what he would like to do
	answer := ""
	prompt := &survey.Select{
		Options: []string{"Config", "Quit"},
	}

	err := survey.AskOne(prompt, &answer)
	if err != nil {
		// Checks if the user interrupted the console
		if err == terminal.InterruptErr {
			fmt.Println("Interuppted")
            os.Exit(0)
		}
	}

    return answer
}
