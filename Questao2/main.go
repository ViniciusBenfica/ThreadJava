package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
	"time"
)

func getFile(done chan bool, files string){

	file, err := os.Open(files+".txt")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	}
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		nomeString := strings.Split(fileScanner.Text(), " ")
		if nomeString[len(nomeString)-1] == "CONCLUIDO" {
			fmt.Println(fileScanner.Text())
		}
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error ao ler o arquivo: %s", err)
	}

    // file.Close()

	done <- true
}

func main() {
	doneGetFile := make(chan bool)

	files := []string{"COMPUTACAO", "DIREITO"}

	for i:= 0; i <= len(files)-1; i++ {
    	go getFile(doneGetFile, files[i])
	}

	// fmt.Println(<-doneGetFile)
	time.Sleep(3 * time.Second)
}