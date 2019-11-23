package common

import (
	"fmt"
	"net/url"
	"os"
)

func PanicIfNonEmpty(err error) {
	if err != nil {
		panic(err)
	}
}


//Config get script env vars or panic
func Config() (string, string, string) {
	user := fmt.Sprintf(url.PathEscape(getEnvValOrPanic("JIRA_USER")))
	pass := fmt.Sprintf(url.PathEscape(getEnvValOrPanic("JIRA_PASS")))
	domain := getEnvValOrPanic("JIRA_DOMAIN")
	return user, pass, domain
}

func getEnvValOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("JIRA_USER", "JIRA_PASS", "JIRA_DOMAIN", "plz set env vars")
		panic(fmt.Sprintf("MISSING %s env var", key))
	}
	return value
}
