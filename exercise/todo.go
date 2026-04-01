package exercise

import "fmt"

func TodoList() {
	var todo []string
	var todoInput string
	var option int

	// looping menu
	for {
		fmt.Println("Simple To-Do List")
		fmt.Println("Pilih opsi: \n 1 = lihat \n 2 = Tambah \n 3 = Selesai")
		fmt.Print("Input: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			// lihat To-Do List
			if len(todo) == 0 {
				fmt.Println("Anda belum menambahkan To-Do List!\n")
			} else {
				for i, list := range todo {
					fmt.Printf("List ke %d : %s\n", i+1, list)
				}
			}
		case 2:
			fmt.Println("Masukkan To-Do List (Ketik 'selesai' untuk berhenti)")

			// scan n input
			for {
				fmt.Print(">")
				fmt.Scanln(&todoInput)

				if todoInput == "selesai" {
					break
				}

				// masukkan data ke array
				todo = append(todo, todoInput)
			}
		case 3:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Perintah tidak dapat dikenali")
		}
	}
}
