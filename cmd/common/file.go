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
	f, err := os.OpenFile(fileame,
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
	file, err := os.Open(fileame)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	fmt.Println(b)
	return &b

}

func CreateIfNotExist(filename string) {
	home, _ := homedir.Dir()
	toCreate := strings.Replace(filename, "~", home, 1)
	_,e := os.Stat(toCreate)
	if os.IsNotExist(e) {
		fmt.Println("init ", toCreate)
		file, err := os.Create(toCreate)
		if err != nil {
			fmt.Println("failed initializing ", toCreate, err.Error())
		}
		file.Close()
	}
}
