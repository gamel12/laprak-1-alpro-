package main

import "fmt"

func main() {
	var nama, kelas, nim string
	fmt.Scan(&nama, &kelas, &nim)
	fmt.Print("Perkenalkan saya adalah ", nama, ", salah satu mahasiswa Prodi S1-IF dari kelas ", kelas, " dengan NIM ", nim)
}
