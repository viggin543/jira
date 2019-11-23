package common

import (
	"github.com/viggin543/jira_cli/common"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Remove("/tmp/da")
	os.Remove("/tmp/TestPrintFileContent")
	run := m.Run()
	os.Remove("/tmp/da")
	os.Remove("/tmp/TestPrintFileContent")
	os.Exit(run)
}


func TestAppendToFuke(t *testing.T) {
	common.AppendToFile("/tmp/da","da")
	common.AppendToFile("/tmp/da","da")


	content, err := ioutil.ReadFile("/tmp/da")
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "da\nda\n" {
		t.Fatal("file should contain dada and not " + string(content))
	}
}

func TestPrintFileContent(t *testing.T) {
	common.AppendToFile("/tmp/TestPrintFileContent","fun\nfun")
	actual := common.PrintFileContent("/tmp/TestPrintFileContent")
	if len(*actual) == 0 {
		t.Fatal("failed to print file content")
	}

}


