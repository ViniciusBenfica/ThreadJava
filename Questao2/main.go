package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func getFile(done chan bool, files string) {

	file, err := os.Open("files/" + files + ".txt")
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

	done <- true
}

func main() {
	doneGetFile := make(chan bool)

	files := []string{"ADMINISTRACAO",
					  "COMPUTACAO", 
					  "DESIGN", 
					  "DIREITO", 
					  "EDUCACAO FISICA", 
					  "ENFERMAGEM", 
					  "ENGENHARIA", 
					  "FOTOGRAFIA", 
					  "JORNALISMO", 
					  "MEDICINA", 
					  "PSICOLOGIA", 
					  "PUBLICIDADE", 
					  "TURISMO", 
					  "SI", 
					  "MATEMATICA",
					}

	for i := 0; i <= len(files)-1; i++ {
		go getFile(doneGetFile, files[i])
	}

	time.Sleep(3 * time.Second)
}
