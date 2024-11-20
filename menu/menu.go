package menu

import (
	"fmt"
	"log"
	"strings"
	"todo-cli/helpers"

	"github.com/eiannone/keyboard"
)

// MenuConfig allows customization of menu behavior
type MenuConfig struct {
	Title            string
	CancelEnabled    bool
	ExitOnCancel     bool
	CancelLabel      string
	SelectedPrefix   string
	UnselectedPrefix string
}

// DefaultConfig provides a standard menu configuration
func DefaultConfig() MenuConfig {
	return MenuConfig{
		Title:            "",
		CancelEnabled:    true,
		ExitOnCancel:     false,
		CancelLabel:      "Cancel",
		SelectedPrefix:   "* ",
		UnselectedPrefix: "  ",
	}
}

// SelectOption displays an interactive menu and returns the selected index
func SelectOption(options []string, config ...MenuConfig) (int, bool) {
	cfg := DefaultConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	selIndex := 0
	running := true

	for running {
		helpers.ClearConsole()

		if cfg.Title != "" {
			fmt.Println(cfg.Title)
			fmt.Println(strings.Repeat("-", len(cfg.Title)))
		}

		for i, action := range options {
			if i == selIndex {
				fmt.Printf("%s%s\n", cfg.SelectedPrefix, action)
			} else {
				fmt.Printf("%s%s\n", cfg.UnselectedPrefix, action)
			}
		}

		if cfg.CancelEnabled {
			fmt.Printf("\n%sPress ESC for %s\n", cfg.UnselectedPrefix, cfg.CancelLabel)
		}

		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selIndex > 0 {
				selIndex--
			}
		case keyboard.KeyArrowDown:
			if selIndex < len(options)-1 {
				selIndex++
			}
		case keyboard.KeyEnter:
			running = false
		case keyboard.KeyEsc:
			if cfg.CancelEnabled {
				if cfg.ExitOnCancel {
					return -1, false
				}
				return selIndex, false
			}
		case keyboard.KeyCtrlC:
			return -1, false
		}
	}

	helpers.ClearConsole()
	return selIndex, true
}
