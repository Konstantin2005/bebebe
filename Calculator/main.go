package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Калькулятор: введите выражение, например (2+3)*4-10/2")
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !sc.Scan() {
			break
		}
		line := sc.Text()
		// пустую строку можно пропустить
		if line == "" {
			continue
		}
		// вычисление
		result := OneCalculate(line)
		fmt.Printf("%s = %v\n", line, result)
	}
}
