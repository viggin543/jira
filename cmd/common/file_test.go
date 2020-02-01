package common_test

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/viggin543/jira/cmd/common"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	createFile("/tmp/da")
	createFile("/tmp/TestPrintFileContent")
	fmt.Println("b4 test")
	run := m.Run()
	fmt.Println("after test")
	os.Remove("/tmp/da")
	os.Remove("/tmp/TestPrintFileContent")
	os.Exit(run)
}

func createFile(s string)  {
	f,_  := os.Create(s)
	f.Close()
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

func TestCreateIfNotExist(t *testing.T) {
	fmt.Println("TestCreateIfNotExist")
	home,_ :=homedir.Dir()

	type args struct {
		filename string
	}

	isExist := func (this * args) bool {
		home,_ :=homedir.Dir()
		path := fmt.Sprintf("%s/.%s", home, strings.Split(this.filename, ".")[1])
		file,_ := os.Stat(path)
		return file != nil
	}
	clean := func (this * args)  {
		toClean := fmt.Sprintf("%s/.%s", home, strings.Split(this.filename, ".")[1])
		os.Remove(toClean)
		fmt.Println("cleaning", toClean)
	}


	tests := []struct {
		name string
		args args
	}{
		{
			name: "create existing file in homedir, do nothing",
			args: args{filename: "~/.existing"},
		},
		{
			name: "create non existing file. create it",
			args: args{filename: "~/.non-existing"},
		},
	}


	createFile(fmt.Sprintf("%s/.existing",home))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := tt.args.filename
			common.CreateIfNotExist(file)
			if !isExist(&tt.args){
				t.Fail()
			}
		})
	}

	for _,tt := range tests {
		clean(&tt.args)
	}

}

func TestFail(t *testing.T) {
	t.Fail()
}


