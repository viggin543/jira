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
	"github.com/viggin543/jira/cmd/common"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a jira task",
	Long: `creates a jira task in current sprint
assigned to assignee, with title and description
and optionally a link to an epic task`,
	Run: func(cmd *cobra.Command, args []string) {
		epic, _ := cmd.Flags().GetInt("epic")

		common.AppendToFile(epics_file, string(epic))

		assignee, _ := cmd.Flags().GetString("assignee")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		issue := createIssue{
			Epic:        epic,
			Title:       title,
			Description: description}
		issue.withAssignee(assignee).Execute()
	},
}

var _, _, domain = common.Config()
var history_file = "~/.jira_tickets"
var epics_file = "~/.jira_epics"

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "title of the ticket")
	assertFlag("title")
	createCmd.Flags().StringP("assignee", "a", "", "the assignee")
	assertFlag("assignee")
	createCmd.Flags().StringP("description", "d", "", "short description of the task")
	assertFlag("description")
	createCmd.Flags().IntP("epic", "e", -1, "link to the epic task if exists")
}

func assertFlag(name string) {
	if err := createCmd.MarkFlagRequired(name); err != nil {
		fmt.Println("missing " + name + " flag")
		os.Exit(1)
	}
}


type createIssue struct {
	Title       string
	Assignee    string
	Description string
	Epic int
}


func (t *createIssue) withAssignee(assignee string) *createIssue {
	team := NewListProjectTeamCommand().NoLogs().Execute()
	for _, member := range team {
		if strings.Contains(member, assignee) {
			t.Assignee = member
			break
		}
	}
	if t.Assignee == "" && assignee != "" {
		fmt.Println("cant find assignee in team:",assignee)
		os.Exit(1)
	}
	return t
}

func (t *createIssue) Execute()  {
	req := common.BuildPostRequest("/rest/api/2/issue/", t.postBody())
	body := common.Execute(req)
	taskNumber := common.ParseToSting(body, "$.key")
	createdTask := fmt.Sprintf("https://%s/browse/%s", domain, taskNumber)

	common.AppendToFile(history_file,createdTask)
	fmt.Println(createdTask)
}


func (t *createIssue) postBody() *bytes.Buffer {

	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{
	"fields": {
	   "project": {"key": "UD"},
	   "summary": "%s",
	   "Description": "%s",
	   "customfield_10064": {"value": "Backend"},
		"customfield_10010":%d,
	   "issuetype": {"name": "Task"},
	   "Assignee": {"name":"%s"}
		%s
		}
	}`,
		t.Title,
		t.Description,
		GetActiveSprint().Id,
		t.Assignee,
		t.getEpicLink())))
	return body
}

func (t *createIssue) getEpicLink() string {
	if t.Epic != 0 {
		return fmt.Sprintf(`,"customfield_10008":"UD-%d"`, t.Epic)
	} else {
		return ""
	}
}

