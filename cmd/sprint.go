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
	"encoding/json"
	"fmt"
	"github.com/viggin543/jira/cmd/common"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// sprintCmd represents the sprint command
var sprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "get active sprint info",
	Long: `get active sprint ID
count of days left
and count of days since start`,
	Run: func(cmd *cobra.Command, args []string) {
		GetActiveSprint().
			Print()
	},
}

func init() {
	rootCmd.AddCommand(sprintCmd)
}


type Sprint struct {
	Id int `json:Id`
	Self string `json:self`
	State string `json:state`
	Name string `json:name`
	StartDate time.Time `json:startDate`
	EndDate time.Time `json:endDate`
	OriginBoardId int `json:originBoardId`
	Goal string `json:goal`
}


func (s *Sprint) Print(){
	fmt.Println(s.Name)
	fmt.Println("State",s.State)
	fmt.Println(int(time.Until(s.EndDate).Hours()/24),"days left")
	fmt.Println(int(time.Since(s.StartDate).Hours()/24),"days since start date")
}

func GetActiveSprint() *Sprint {
	request := common.BuilGetRequest("/rest/agile/1.0/board/17/sprint?state=active")
	response := common.Execute(request)
	bytes, _ := json.Marshal(common.JPathGet(response, "$.values[0]"))
	return unmarshal(bytes)
}

func unmarshal(response []byte) *Sprint {
	var sprint = Sprint{}
	err := json.Unmarshal(response, &sprint)
	if err != nil {
		println("failed to parse response to Sprint", response)
		os.Exit(1)
	}
	return &sprint
}