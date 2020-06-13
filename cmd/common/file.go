package common

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func AppendToFile(fileame string, text string){
	f, err := os.OpenFile(ExpandHomeDir(fileame),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text+"\n"); err != nil {
		log.Println(err)
	}
}

func PrintFileContent(fileame string) *[]byte {
	file, err := os.Open(ExpandHomeDir(fileame))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	fmt.Println(string(b))
	return &b

}

func CreateIfNotExist(filename string) {
	toCreate := ExpandHomeDir(filename)
	isNotExists := IsNotExist(filename)
	if isNotExists {
		fmt.Println("init ", toCreate)
		file, err := os.Create(toCreate)
		if err != nil {
			fmt.Println("failed initializing ", toCreate, err.Error())
		}
		file.Close()
	}
}

func IsNotExist(filename string)  bool {
	toCreate := ExpandHomeDir(filename)
	_, e := os.Stat(toCreate)
	isNotExists := os.IsNotExist(e)
	return  isNotExists
}

func ExpandHomeDir(filename string)  string {
	home, _ := homedir.Dir()
	toCreate := strings.Replace(filename, "~", home, 1)
	return toCreate
}
