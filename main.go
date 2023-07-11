package main

import (
	"fmt"
	"os"

	"github.com/philip727/gcc-scraper/notifier"
	"github.com/philip727/gcc-scraper/types"
	"github.com/philip727/gcc-scraper/ui"
)

func main() {
    settings := types.Settings{
        SelectedProducts: make([]string, 0),
    }
	cache := types.Cache{
		StockedItems: make(map[string]bool),
	}

    if err := settings.LoadSettings(); err != nil {
        fmt.Printf("Couldn't load settings file %s", err)
    }

	// Runs a thread in the background to constantly update the store
	go notifier.ConstantlyCheckStore(&settings, &cache)

	for {
		// Asks the user what he would like to do
		answer := ui.MainUI()
		// Different menus
		switch answer {
		case "Quit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		case "Config":
			ui.ConfigUI(&settings)
		}
	}
}
