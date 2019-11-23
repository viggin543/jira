package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
