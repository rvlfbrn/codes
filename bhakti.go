package main

import "fmt"

type siswa struct {
	nama  string
	nilai float64
}

func main() {
	var hasil, gagal []siswa
	var jumlahSiswa int
	var total float64

	total = 0

	fmt.Print("Masukkan Jumlah Siswa: ")
	fmt.Scanln(&jumlahSiswa)

	for i := 0; i < jumlahSiswa; i++ {
		fmt.Println("\nData ke:", i+1)

		fmt.Print("Masukkan Nama: ")
		var nama string
		fmt.Scanln(&nama)

		fmt.Print("Masukkan Nilai: ")
		var nilai float64
		fmt.Scanln(&nilai)

		dataSiswa := siswa{
			nama:  nama,
			nilai: nilai,
		}
		hasil = append(hasil, dataSiswa)

		total += nilai
	}

	for _, v := range hasil {
		fmt.Println("\nNama:", v.nama)
		fmt.Println("Nilai:", v.nilai)

		if v.nilai < 60 {
			gagal = append(gagal, v)
		}
	}

	fmt.Println("\nrata - ratanya adalah:", total/float64(jumlahSiswa))

	fmt.Println("\n----- Siswa yang gagal -----")
	for _, v := range gagal {
		fmt.Println("\nNama:", v.nama)
		fmt.Println("Nilai:", v.nilai)
	}
}
