package main

import (
	"fmt"
	"time"
)

/*
usar a palavra reservada "go" permite criar uma nova go routine
ou seja uma thread para que as funçoes sejam executadas paralelamente
ao invés de sequencial

As go routines, diferente de threads trabalham para consumir menos memoria
diferente de threads comuns onde voce realiza diversas chamadas ao SO (call systems)
as go routines sao criadas na userland, permitindo assim a criaçao de milhares de threads
para lidar com multiplas tarefas em multi cores
*/
func goRoutines() {
	// t1
	println("teste 1")
	// t1
	println("teste 2")
	// t1
	println("teste 3")
	time.Sleep(time.Second)
}

func task(name string) {
	for i := 1; i <= 10; i++ {
		println(fmt.Sprintf("%s:%d", name, i))
		time.Sleep(time.Second)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int) {
	for c := range ch {
		fmt.Println(c)
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	workerQuantity := 3

	for i := 0; i < workerQuantity; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 10000; i++ {
		canal <- i
	}
}

