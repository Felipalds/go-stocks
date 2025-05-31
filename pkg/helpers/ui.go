package helpers

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var confirmStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("%FFD700")).Bold(true)

func ConfirmPrompt(message string) (bool, error) {
	fmt.Print(confirmStyle.Render(message + " [y/N]:"))

	var response string

	fmt.Scanln(&response)

	if response == "n" || response == "N" || response == "" {
		return false, nil
	} else {
		if response == "y" || response == "Y" {
			return true, nil
		} else {
			return false, errors.New("Unexpected answer")
		}
	}
}
