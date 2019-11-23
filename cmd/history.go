/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/viggin543/jira/cmd/common"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "print recently created tickets",
	Long: `prints recent tickets, every ticket that was created is saved in ~/.jira_tickets`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintFileContent(epics_file)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
