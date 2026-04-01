package exercise

import "fmt"

var operasi = map[string]func(int, int) int{
	"+": func(i, j int) int { return i + j },
	"-": func(i, j int) int { return i - j },
	"*": func(i, j int) int { return i * j },
	"/": func(i, j int) int { return i / j },
}

func Calculator() {
	var i, j int
	var operator string

	fmt.Println("Masukkan nomor pertama")
	fmt.Scanln(&i)

	fmt.Println("Masukkan nomor kedua")
	fmt.Scanln(&j)

	fmt.Println("Pilih Operator (+, -, /, *):")
	fmt.Scanln(&operator)

	// checking operator
	if op, ok := operasi[operator]; ok {
		hasil := op(i, j)
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Operator tidak dapat dikenali!")
	}
}
