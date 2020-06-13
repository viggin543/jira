package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/viggin543/jira/cmd/common"
	"io/ioutil"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "jira",
	Short: "Jira command line client",
	Long: `Jira is super slow, especially 4 creating tasks
this client uses jira api in order to automate things
its main function is to create tickets for current sprint
with an assignee from you't team
optionally related to some epic.

the cli requires the following env vars.
export JIRA_USER="banana@opa.com"
export JIRA_PASS="API_TOKEN_HERE"
export JIRA_DOMAIN="opa.atlassian.net"
export JIRA_PROJECT="UD"

Or create the a config file in ~/.jira.yaml with the above keys if you dont want to export env vars

To create a token browse here:
https://Id.atlassian.com/manage/api-tokens

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {


	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	common.InnitConfig()
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run

	common.CreateIfNotExist(history_file)
	common.CreateIfNotExist(epics_file)
	common.CreateIfNotExist(epics_file)
	createConfigTemplate()

}

func createConfigTemplate() {
	config_path := "~/.jira.yaml"
	if common.IsNotExist(config_path) {
		ioutil.WriteFile(common.ExpandHomeDir(config_path), []byte(`
JIRA_USER: "some@ourbond.com"
JIRA_PASS: "api-token-from -> https://Id.atlassian.com/manage/api-tokens"
JIRA_DOMAIN: "some.atlassian.net"
JIRA_PROJECT: "some-project-name"
DEFAULT_ASSIGNEE: "some-project-name"
DEFAULT_STATUE: "to do"
`), 0644)
	}
}

// initConfig reads in config file and ENV variables if set.
