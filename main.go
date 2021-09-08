package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
)

func readTxt(name string) []string{
	//Function read a txt file and return a string list
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	return txtlines
}





//We going to cipher the message with Feistel cipher permutation

func main(){

	//Exampling reading file
	file:=readTxt("mensajedeentrada.txt") //read the text
	hash := sha1.New()//generate the sha1
	hash.Write([]byte(file[0]))//using the sha1 function
	bs := hash.Sum(nil)
	println(file[0])
	println(bs)
	fmt.Printf("%x\n", bs)




}