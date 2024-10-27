package ui

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
)

func SelectMenu(menuStage *string) {
	fmt.Println("--------------------------------")
	fmt.Println("Please select an option:")
	fmt.Println("1. Search Appointment")
	fmt.Println("2. Search With My Last Configuration")
	fmt.Println("3. Exit")

	SelectOption("Selection: ", []string{"1", "2", "3"}, menuStage)
}

func PrintInfoMessage(message string) {
	fmt.Println("\n--------------------------------")
	fmt.Println(message)
	fmt.Println("--------------------------------\n")
}

func SelectOption(prompt string, options []string, v *string) {
	p := &survey.Select{
		Message: prompt,
		Options: options,
	}

	if err := survey.AskOne(p, v); err != nil {
		panic(err)
	}

	if *v == "" {
		*v = constants.NO_SELECTION
	}
}

func SelectOptions(prompt string, options []string, v *[]string) {
	p := &survey.MultiSelect{
		Message: prompt,
		Options: options,
	}

	if err := survey.AskOne(p, v); err != nil {
		panic(err)
	}

	hasDefaultSelection := false
	for _, val := range *v {
		if val == constants.NO_SELECTION {
			hasDefaultSelection = true
			break
		}
	}

	if hasDefaultSelection {
		*v = []string{constants.NO_SELECTION}
	}

	if len(*v) == 0 {
		*v = []string{constants.NO_SELECTION}
	}
}

func GetInput(prompt string, v *string) {
	p := &survey.Input{
		Message: prompt,
	}

	if err := survey.AskOne(p, v); err != nil {
		panic(err)
	}
}
