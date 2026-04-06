package exercise

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}

type student struct {
	person
	grade int
	major string
}

func Absensi() {
	var students []student
	var input int

	var s student

	// test data
	students = append(students, student{
		person: person{
			name: "Jamal",
			age:  20,
		},
		grade: 11,
		major: "Computer Science",
	})

	for {
		fmt.Println("Selamat datang di sistem pengabsenan siswa, silahkan pilih opsi yang ada")
		fmt.Println("1. Lihat list liswa\n2. Tambah Data siswa\n3. Selesai")

		fmt.Print("input: ")
		fmt.Scanln(&input)

		// logic section
		switch input {
		case 1:
			if len(students) == 0 {
				fmt.Println("Anda belum menginput data siswa!")
			} else {
				fmt.Println("Daftar siswa yang terverifikasi")
				for i, s := range students {
					fmt.Printf("%d. Nama: %s. Umur: %d. Kelas: %d. Jurusan: %s\n", i+1, s.person.name, s.person.age, s.grade, s.major)
				}
			}
		case 2:
			inputLabel := []string{"Nama", "Umur", "Kelas", "Jurusan"}

			fmt.Println("Silahkan masukkan data diri dengan lengkap.")
			for i, inputLabel := range inputLabel {
				fmt.Print(inputLabel, ": ")

				switch i { // case sesuai index array yang ingin ditampilkan
				case 0:
					fmt.Scan(&s.person.name) // embedded struct
				case 1:
					for {
						// kuduna validasi program ieu bisa dipangkas jadi function, ngarah teu duplikat., tapi nya.., ngke deui weh. wkwkwk
						var ageInput string
						fmt.Scanln(&ageInput)

						age, err := strconv.Atoi(ageInput)
						if err != nil {
							fmt.Println("Masukkan angka yang valid!")
							fmt.Print(inputLabel, ": ")
							continue
						}

						s.person.age = age
						break
					}
				case 2:
					for {
						var gradeInput string
						fmt.Scanln(&gradeInput)

						grade, err := strconv.Atoi(gradeInput)
						if err != nil {
							fmt.Println("Gunakan angka untuk mengisi data")
							fmt.Print(inputLabel, ": ")
							continue
						}

						s.grade = grade
						break
					}
				case 3:
					fmt.Scan(&s.major)
				default:
					fmt.Println("Tidak dapat dikenali")
					return
				}
			}

			// appending to students data
			students = append(students, s)

			// notification
			fmt.Println("Data murid atas nama ", s.person.name, " berhasil dimasukkan!")

		case 3:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Perintah tidak dapat dikenali")
		}
	}

}
