package main

import (
	"fmt"
)

type Person struct {
	name, kegiatan, daerah string
	age                    int
}

var mante = []Person{
	{name: "Rika", age: 22, kegiatan: "Seorang mahasiswa akhir", daerah: "Kalimantan Timur"},
	{name: "Dwi", age: 25, kegiatan: "Jr Frontend", daerah: "Sulawesi Tengah"},
	{name: "Fitri", age: 27, kegiatan: "Fullstack", daerah: "Jakarta"},
	{name: "Aya", age: 29, kegiatan: "Fullstack", daerah: "Makassar"},
	{name: "Kartika", age: 22, kegiatan: "Mobile Software Engineer", daerah: "Jakarta"},
}

func main() {
	var cari string
	fmt.Printf("Masukkan Nama yang dicari : ")
	fmt.Scan(&cari)
	for k, mantes := range mante {
		if cari == mantes.name {
			fmt.Println("-----------------------------------")
			fmt.Println("Nama : ", mantes.name, "\nUmur : ", mantes.age, "\nKegiatan : ", mantes.kegiatan, "\nDaerah : ", mantes.daerah)
			fmt.Println("-----------------------------------")
			break
		} else if cari != mantes.name && k+1 >= cap(mante) {
			fmt.Println("Nama tidak tersedia, perhatikan huruf kapital")
		}
	}
}
