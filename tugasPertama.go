package main

import (
	"fmt"
	"os"
)

func showFriend(friendIndex int) {

	type Friend struct {
		nama      string
		alamat    string
		pekerjaan string
		alasan    string
	}

	var friends = map[int]Friend{
		1: {"Mamad", "Jalan Ijen", "mahasiswa", "gabut"},
		2: {"Elsa", "Jalur Gaza", "mahasiswa", "Membuat website untuk menggalakkan perdamaian dunia"},
		3: {"Donny", "Tanah Abang", "pengacara", "ingin bikin aplikasi gugatan cerai online"},
		4: {"Budi", "Gunung Salak", "pengangguran", "ingin membuat program yang membuka jalan menuju jodoh"},
		5: {"Susi", "Lautan", "penyelam", "biar bisa bahasa programming seperti bahasa ikan"},
		6: {"Andi", "Bukit Timah", "pencari harta karun", "mengira ini kelas untuk mencari harta karun digital"},
		7: {"Fandi", "Blok M", "pencari cinta sejati", "karena dia percaya jodoh pasti bertemu, termasuk dengan bahasa pemrograman"},
		8: {"Lina", "Kampung Inggris", "guru bahasa Inggris", "mengira Golang adalah bahasa Inggris baru"},
		9: {"Rudi", "Tebing Tinggi", "pemburu hantu", "karena ingin membuat program yang bisa menangkap hantu"},
		10: {"Dina", "Gunung Kidul", "pembuat batu akik", "ingin membuat program yang bisa memprediksi harga batu akik"},
	}

	val, available := friends[friendIndex]
	if !available {
		fmt.Println("Data tidak ditemukan")
		return
	}

	println("Nama:", val.nama)
	println("Alamat:", val.alamat)
	println("Pekerjaan:", val.pekerjaan)
	println("Alasan mengikuti kelas Golang:", val.alasan)

}

func main() {
	input := os.Args

	friendIndex := 0

	_,notNumber := fmt.Sscanf(input[1], "%d", &friendIndex)

	if notNumber != nil {
		fmt.Println("Masukkan data berupa angka 1-10")
	}

	showFriend(friendIndex)
}
