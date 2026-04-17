package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Orders struct {
	Nama     string
	Harga    int
	Jumlah   int
	Subtotal int
}

var menu = map[string]int{
	"Mie Ayam":    10000,
	"Nasi Goreng": 12000,
	"Es Teh":      5000,
	"Ayam Goreng": 18000,
}

var order []Orders
var optionInput int
var selectMenuOption string

func selectMenu(message string) {
	var qty int

	for {
		fmt.Print(message)

		// scan semua input
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			selectMenuOption = scanner.Text()
		}

		if strings.ToLower(selectMenuOption) == "selesai" {
			fmt.Println("Selesai memilih...")
			break
		} else {
			fmt.Print("Jumlah pesanan : ")
			fmt.Scanln(&qty)
		}

		var namaPesanan string = selectMenuOption
		var harga int = menu[selectMenuOption]
		var subtotal int = qty * harga

		if _, ok := menu[selectMenuOption]; ok {
			order = append(order, Orders{
				Nama:     namaPesanan,
				Harga:    harga,
				Jumlah:   qty,
				Subtotal: subtotal,
			})
		} else {
			if strings.ToLower(selectMenuOption) != "selesai" {
				fmt.Printf("Maaf, untuk item %s, tidak tersedia dalam menu kami\n", selectMenuOption)
			}
		}
	}
}

func deleteOrder() {
	var index int
	var deleteQuestion string

	if len(order) == 0 {
		fmt.Println("Tidak ada pesanan yang harus dihapus")
		return
	}

	fmt.Println("Pesanan mana yang akan dibatalkan?")
	fmt.Println("==================================")
	for i, v := range order {
		fmt.Printf("%d. %s - %d x %d = %d\n", i+1, v.Nama, v.Harga, v.Jumlah, v.Subtotal)
	}
	fmt.Println("==================================")

	fmt.Scanln(&index)

	if index < 1 || index > len(order) {
		fmt.Println("Nomor tidak valid")
		return
	}

	fmt.Printf("Anda yakin ingin menghapusnya? (y / t) : ")
	fmt.Scanln(&deleteQuestion)

	if strings.ToLower(deleteQuestion) != "y" {
		return
	} else {
		idx := index - 1
		order = append(order[:idx], order[idx+1:]...)

		fmt.Println("Pesanan berhasil dihapus!")
	}
}

func selectMenuAgain() {
	var questionInput int

	fmt.Println("Apakah anda ingin memesan lagi? : ")
	fmt.Println("1. Ya, saya mau pesan lagi")
	fmt.Println("2. Tidak, sudah cukup")
	fmt.Println("3. Saya mau membatalkan pesanan")
	fmt.Scanln(&questionInput)

	switch questionInput {
	case 1:
		for key, v := range menu {
			fmt.Printf("%s : %d\n", key, v)
		}
		selectMenu("Silahkan pilih menu lagi : ")
	case 2:
		fmt.Println("Oke, silahkan konfirmasi pesanan ya..,")
	case 3:
		deleteOrder()
	default:
		fmt.Println("Hell nawh.., lu niat pesan gak sih?")
	}
}

func RunOrderManagement() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Occured : ", r)
		}
	}()

	for {
		fmt.Println("Selamat datang di restoran kami, Silahkan pilih opsi yang tersedia")
		fmt.Println("1. Pilih pesanan")
		fmt.Println("2. Tampilkan pesanan anda")
		fmt.Println("3. Selesai & Konfirmasi pesanan")

		fmt.Print("> ")
		fmt.Scanln(&optionInput)

		// logic section
		switch optionInput {
		case 1:
			// tampilkan pesanan
			fmt.Println("Berikut menu andalan yang tersedia di toko kami :")
			for key, v := range menu {
				fmt.Printf("%s : %d\n", key, v)
			}

			selectMenu("Pilih pesanan anda : ")
		case 2:
			if len(order) == 0 {
				fmt.Println("Anda belum memesan menu apapun!")
			} else {
				fmt.Println("==========")
				for i, v := range order {
					fmt.Printf("%d. %s - %d x %d = %d\n", i+1, v.Nama, v.Harga, v.Jumlah, v.Subtotal)
				}
				// fmt.Println(order)
				fmt.Println("==========")

				// tambah logic display total

				selectMenuAgain()
			}
		case 3:
			fmt.Println("Terimakasih, Pesanan anda akan segera kami proses :)")
			return
		default:
			fmt.Println("Perintah tidak dikenali")
		}
	}
}
