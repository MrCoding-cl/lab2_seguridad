package main

import (
	"bufio"
	"github.com/cyrildever/feistel"
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
	file:=readTxt("mensajedeentrada.txt")
	println(file[0])

	cipher := feistel.NewCipher("some-32-byte-long-key-to-be-safe", 10)



}