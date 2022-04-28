package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pipoca(done chan <- bool) {

	min := 1000
    max := 10000
	
	random := rand.Intn(max - min) + min

    fmt.Println("Preparando Pipocada")
    time.Sleep(time.Duration(random) * time.Millisecond)

    fmt.Println("Pipoca preparada")

    done <- true
}

func refri(done chan <- bool) {

	min := 1000
    max := 10000
	
	random := rand.Intn(max - min) + min

    fmt.Println("Preparando Refri")
    time.Sleep(time.Duration(random) * time.Millisecond)

    fmt.Println("Refri preparado")
    
    done <- true
}

func funcionaraia(done chan <- bool) {
    done <- true
}

func main() {

    donePipoca := make(chan bool)
    doneRefri := make(chan bool)
    doneFuncionario := make(chan bool)
    
    for {
        go funcionaraia(doneFuncionario)
		if <-doneFuncionario == true {
            go pipoca(donePipoca)
            go refri(doneRefri)
            <-donePipoca
            <-doneRefri
			fmt.Println("Pedido entregue")
		}
        time.Sleep(1000 * time.Millisecond)
    }

}