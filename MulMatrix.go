package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
	"regexp"
)

func readFile(path string) (string,error){
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return "",err
	}
	return string(data),nil
}
//
func getMatrixMetadata(metadataLine string){
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(metadataLine,-1))
}
func main() {
	content,ok := readFile("./matrices-2.txt")
	if ok !=nil {
		panic("Cannot read file")
	}
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
