package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/philip727/gcc-notifier/gcc"
	"github.com/philip727/gcc-notifier/types"
)

func ConfigUI(s *types.Settings) {
	// Requests the list of products from gcc
	var products types.ProductRequest
	err := gcc.RequestProducts(&products)
	if err != nil {
		fmt.Printf("Err when making request to gcc: %s", err)
	}

	items := []string{}
    prompt := createItemSelectionPrompt(s, products)
	survey.AskOne(prompt, &items)

    // Sets the selected produts
    s.SelectedProducts = items

    if err := s.SaveSettings(); err != nil {
        fmt.Printf("Error whilst saving: %s", err)
    }
}

func createItemSelectionPrompt(s *types.Settings, p types.ProductRequest) *survey.MultiSelect {
	// Creates an array with each product name
	selection := make([]string, 0)
    checked := make([]string, 0)
	for _, product := range p.Products {
		selection = append(selection, product.Name)

        // Shows previous items as checked
        if s.ItemSelected(product.Name) {
            checked = append(checked, product.Name)
        }
	}
    
	// Prompt for choosing the items you wish to show
	prompt := &survey.MultiSelect{
		Message: "Notify when these items are in stock",
		Options: selection,
        Default: checked,
	}

    return prompt
}
