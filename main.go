package main

import (
	"fmt"
	"sync"
)

// Sender Отправляет числа в канал.
func sender(c chan<- int, wg *sync.WaitGroup) {
	defer func() {
		close(c)
		wg.Done()
	}()
	for i := 1; i <= 100; i++ {
		c <- i
	}
}

// Recv Получает числа из канала.
func recv(c <-chan int, wg *sync.WaitGroup, method int) {
	defer func() {
		wg.Done()
	}()
	if method == 1 {
		for {
			select {
			case num, ok := <-c:
				if !ok {
					return
				}
				fmt.Println(num)
			}
		}
	} else {
		for num := range c {
			fmt.Println(num)
		}
	}
}

func main() {
	c := make(chan int, 100)
	var wg sync.WaitGroup
	var method int
	fmt.Println("Выберите метод исполнения (1 - select, 2 - range): ")
	_, err := fmt.Scanln(&method)
	if err != nil {
		fmt.Println("Неверное значение, по умолчанию метод исполнения 1")
		method = 1
	}
	if method > 2 || method < 1 {
		method = 1
	}
	go sender(c, &wg)
	go recv(c, &wg, method)
	wg.Add(2)
	wg.Wait()
}
