package exercise

import (
	"errors"
	"fmt"
)

var operasi = map[string]func(int, int) int{
	"+": func(i, j int) int { return i + j },
	"-": func(i, j int) int { return i - j },
	"*": func(i, j int) int { return i * j },
	"/": func(i, j int) int { return i / j },
}

func divideZeroWarn(j int, operator string) (bool, error) {
	if j == 0 && operator == "/" {
		return false, errors.New("Tidak bisa dibagi 0!")
	}

	return true, nil
}

func Calculator() {
	var i, j int
	var operator string

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[System] Terjadi kesalahan :", r)
		}
	}()

	fmt.Print("Masukkan nomor pertama : ")
	fmt.Scanln(&i)

	fmt.Print("Masukkan nomor kedua : ")
	fmt.Scanln(&j)

	fmt.Print("Pilih Operator (+, -, /, *) : ")
	fmt.Scanln(&operator)

	// checking operator
	if op, ok := operasi[operator]; ok {

		// validasi sebelum terjadi panic
		if valid, err := divideZeroWarn(j, operator); !valid {
			fmt.Println("Error :", err)
			return
		}

		hasil := op(i, j)
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Operator tidak dapat dikenali!")
	}
}
