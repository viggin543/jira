
package cmd

import (
  "fmt"
  "github.com/viggin543/jira/cmd/common"
  "os"
  "github.com/spf13/cobra"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)


var cfgFile string
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
https://id.atlassian.com/manage/api-tokens

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
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jira.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


  common.CreateIfNotExist(history_file)
  common.CreateIfNotExist(epics_file)
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".jira" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".jira")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

