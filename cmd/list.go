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
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/viggin543/jira/cmd/common"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list you'r jira tasks",
	Long: `lists jira tasks assigned to your $JIRA_USER
that are with status in: (To Do ,In Progress,Review)
`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		issues := getIssues()
		if verbose {
			printOutVerbose(issues)
		} else {
			printOut(issues)
		}
	},
}


func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "verbose output")
}



func getIssues() *[]interface{} {
	path := "/rest/api/2/search"
	user := viper.GetString("jira_user")
	body := fmt.Sprintf(`{"jql":
									"Assignee = \"%s\" AND Status in (\"To Do\" ,\"In Progress\",\"Review\")",
									"startAt":0,
									"maxResults":30}`, user)
	req := common.BuildPostRequest(
		path,
		bytes.NewBuffer([]byte(body)))
	resp := common.Execute(req)
	issues := common.JPathGet(resp, "$.issues").([]interface{})
	return &issues
}

func printOutVerbose(issues *[]interface{}) {
	fmt.Println("=========================================================================================")
	for _, v := range *issues {
		issue := v.(map[string]interface{})
		key := issue["key"]
		fmt.Println(key)
		summary := issue["fields"].(map[string]interface{})["summary"]
		fmt.Println(summary)
		fmt.Println(fmt.Sprintf("https://tg17home.atlassian.net/browse/%s", key))
		fmt.Println("=========================================================================================")
	}
}
func printOut(issues *[]interface{}) {
	for _, v := range *issues {
		ticket := v.(map[string]interface{})["key"]
		fmt.Println(ticket)
	}
}