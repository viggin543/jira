/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"github.com/viggin543/jira/cmd/common"
	"strings"
)
//     "untranslatedName": "Review",
//    "untranslatedName": "Blocked",
//    "untranslatedName": "To Do",
//     "untranslatedName": "In Progress",
//    "untranslatedName": "Done",

// transitionCmd represents the transition command
var transitionCmd = &cobra.Command{
	Use:   "transition",
	Short: "Transition a ticket to a state",
	Long: `For example:
- move a ticket from todo do code review
- or from todo to in progress
`,
	Run: func(cmd *cobra.Command, args []string) {
		ticket, _ := cmd.Flags().GetString("ticket")
		state, _ := cmd.Flags().GetString("to")
		TransitionTo(ticket,state)

	},
}

func TransitionTo(ticket string, state string) {
	//todo: insert ud if needed
	if !strings.Contains(strings.ToLower(ticket), "ud-") {
		panic("ticket number must be of the following format: UD-123")
	}
	transitionId := getTransitions(ticket, state)
	if transitionId != nil {
		body := transitionTo(ticket, *transitionId)
		fmt.Println(string(body), ticket)
	} else {
		fmt.Println(fmt.Sprintf("no transitionId '%s' available for %s", state, ticket))
	}
}


func getTransitions(ticket string, state string) *string {
	path := fmt.Sprintf("/rest/api/2/issue/%s/transitions", ticket)
	req := common.BuilGetRequest(path)
	resp := common.Execute(req)
	ticketTransition := common.ParseToSplitStr(resp, "$.transitions[*][id,name]")
	return matchingTransitionId(ticketTransition, state)
}

func matchingTransitionId(ticketTransition [][2]string, state string) *string {
	for _, v := range ticketTransition {
		if strings.ToLower(state) != "" &&
			strings.Contains(strings.ToLower(v[1]), strings.ToLower(state)) {
			return &v[0]
		}
	}
	return nil
}

func transitionTo(ticket string, id string) []byte {
	path := fmt.Sprintf("/rest/api/3/issue/%s/transitions", ticket)
	buffer := bytes.NewBuffer([]byte(fmt.Sprintf(`{"transition":{"id":%s}}`, id)))
	req := common.BuildPostRequest(path, buffer)
	return common.Execute(req)
}

func init() {
	rootCmd.AddCommand(transitionCmd)
	transitionCmd.Flags().StringP("ticket", "t", "", "ticket number")
	common.AssertFlag(transitionCmd,"ticket")
	transitionCmd.Flags().StringP("to", "s", "", "ticket number")
	common.AssertFlag(transitionCmd,"to")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transitionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
