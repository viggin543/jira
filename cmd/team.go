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
	"fmt"
	"github.com/spf13/viper"
	"github.com/viggin543/jira/cmd/common"

	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "list all members of UD team in jira",
	Long: `all list the jira names of all members of the UD project`,
	Run: func(cmd *cobra.Command, args []string) {
		NewListProjectTeamCommand().Execute()
	},
}

func init() {
	rootCmd.AddCommand(teamCmd)
}



func NewListProjectTeamCommand() *listProjectTeam {
	return &listProjectTeam{quiet:false}
}

//listProjectTeam ...
type listProjectTeam struct {
	quiet bool
}

func (t *listProjectTeam) NoLogs() *listProjectTeam {
	t.quiet = true
	return t
}

func (t *listProjectTeam) Execute() []string {
	project := viper.GetString("jira_project")
	req := common.BuilGetRequest(fmt.Sprintf("/rest/api/2/user/assignable/search?project=%s",project))
	body := common.Execute(req)
	jiraUsers := common.ParseToSplitStr(body, "$..name")
	t.print(jiraUsers)
	return jiraUsers
}

func (t *listProjectTeam) print(jiraUsers []string) {
	if !t.quiet {
		for _, user := range jiraUsers {
			fmt.Println(user)
		}
	}

}

