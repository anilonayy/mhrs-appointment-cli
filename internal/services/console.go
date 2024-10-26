package services

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func SelectMenu(menuStage *string) {
	fmt.Println("--------------------------------")
	fmt.Println("Please select an option:")
	fmt.Println("1. Search Appointment")
	fmt.Println("2. Search With My Last Configuration")
	fmt.Println("3. Exit")

	SelectOption("Selection: ", []string{"1", "2", "3"}, menuStage)
}

func SelectOption(prompt string, options []string, v *string) {
	p := &survey.Select{
		Message: prompt,
		Options: options,
	}

	if err := survey.AskOne(p, v); err != nil {
		panic(err)
	}
}
