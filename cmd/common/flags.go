package common

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func AssertFlag(command *cobra.Command ,name string) {
	if err := command.MarkFlagRequired(name); err != nil {
		fmt.Println("missing " + name + " flag")
		os.Exit(1)
	}
}
