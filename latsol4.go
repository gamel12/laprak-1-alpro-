package main

import "fmt"

func main() {
	var C, F float64
	fmt.Print("masukan fahrenheit")
	fmt.Scan(&F)
	C = (F - 32) * 5 / 9
	fmt.Print("hasilnya adalah ", C)
}
