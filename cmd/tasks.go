package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/viggin543/jira/cmd/common"
)

type Task struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "print current sprint tasks",
	Long: `
print current tasks issue number / Title`,
	Run: func(cmd *cobra.Command, args []string) {
		user := viper.GetString("jira_user")
		bodyString := fmt.Sprintf(
			`{"jql":"assignee = \"%s\"AND Status in (%s) ","startAt":0,"maxResults":30}`,
			user,
			getJiraFilter(cmd))
		req := common.BuildPostRequest("/rest/api/2/search",
			bytes.NewBuffer([]byte(bodyString)))
		respBody := common.Execute(req)
		tasks := parseResp(respBody)
		prettyPrint(tasks)
	},
}


func getJiraFilter(cmd *cobra.Command) string {
	progress, _ := cmd.Flags().GetBool("in-progress")
	if progress {
		return `\"In Progress\"`
	} else {
		return `\"To Do\" ,\"In Progress\",\"Review\"`
	}
}
func prettyPrint(tasks []Task) {
	jsonTasks, _ := json.Marshal(tasks)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, jsonTasks, "", "\t")
	fmt.Println(prettyJSON.String())
}

func parseResp(respBody []byte) []Task {
	issues := common.JPathGet(respBody, "$.issues").([]interface{})
	tasks := make([]Task, len(issues))
	for i, e := range issues {
		task := Task{
			Id:    e.(map[string]interface{})["key"].(string),
			Title: e.(map[string]interface{})["fields"].(map[string]interface{})["summary"].(string)}
		tasks[i] = task
	}
	return tasks
}

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.Flags().BoolP("in-progress", "i", false, "in progress only")
}
