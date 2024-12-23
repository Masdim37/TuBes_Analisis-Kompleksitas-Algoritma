package main

import (
	"fmt"
	"time"
)

// Karyawan menyimpan informasi tentang karyawan
type Karyawan struct {
	nama    string
	jabatan string
}

// Fungsi untuk menampilkan daftar nama karyawan berdasarkan jabatan
func TampilkanDaftarNama(karyawan []Karyawan) {
	var jabatan string
	fmt.Print("Masukkan jabatan untuk melihat daftar nama: ")
	fmt.Scanln(&jabatan)

	fmt.Println("Daftar Nama Karyawan dengan Jabatan:", jabatan)
	for _, k := range karyawan {
		if k.jabatan == jabatan {
			fmt.Println(k.nama)
		}
	}
}

// Fungsi untuk menghitung gaji berdasarkan hari masuk
func HitungGaji() {
	var jabatan string
	var hariMasuk int
	fmt.Print("Masukkan jabatan: ")
	fmt.Scanln(&jabatan)
	fmt.Print("Masukkan jumlah hari masuk: ")
	fmt.Scanln(&hariMasuk)

	var gaji int
	switch jabatan {
	case "CEO":
		gaji = hariMasuk * 1000000
	case "Direktur":
		gaji = hariMasuk * 750000
	case "Manajer":
		gaji = hariMasuk * 500000
	case "Pekerja":
		gaji = hariMasuk * 250000
	}
	fmt.Printf("Gaji karyawan: %d\n", gaji)
}

// Fungsi untuk menghitung gaji berdasarkan performa
func HitungGajiPerforma() {
	var jabatan string
	var hariMasuk, performa int
	fmt.Print("Masukkan jabatan: ")
	fmt.Scanln(&jabatan)
	fmt.Print("Masukkan jumlah hari masuk: ")
	fmt.Scanln(&hariMasuk)
	fmt.Print("Masukkan nilai performa (0-100): ")
	fmt.Scanln(&performa)

	gaji := HitungGajiManual(jabatan, hariMasuk)

	if performa >= 90 {
		gaji += gaji * 20 / 100
	} else if performa >= 70 {
		gaji += gaji * 10 / 100
	}
	fmt.Printf("Gaji berdasarkan performa: %d\n", gaji)
}

// Fungsi manual untuk menghitung gaji
func HitungGajiManual(jabatan string, hariMasuk int) int {
	var gaji int
	switch jabatan {
	case "CEO":
		gaji = hariMasuk * 1000000
	case "Direktur":
		gaji = hariMasuk * 750000
	case "Manajer":
		gaji = hariMasuk * 500000
	case "Pekerja":
		gaji = hariMasuk * 250000
	}
	return gaji
}

// Fungsi untuk menghitung gaji overtime
func HitungGajiOvertime() {
	var jabatan string
	var hariMasuk, jamOvertime int
	fmt.Print("Masukkan jabatan: ")
	fmt.Scanln(&jabatan)
	fmt.Print("Masukkan jumlah hari masuk: ")
	fmt.Scanln(&hariMasuk)
	fmt.Print("Masukkan jumlah jam overtime: ")
	fmt.Scanln(&jamOvertime)

	gaji := HitungGajiManual(jabatan, hariMasuk)
	var gajiPerJamOvertime int

	switch jabatan {
	case "CEO":
		gajiPerJamOvertime = 100000
	case "Direktur":
		gajiPerJamOvertime = 75000
	case "Manajer":
		gajiPerJamOvertime = 50000
	case "Pekerja":
		gajiPerJamOvertime = 25000
	}

	gaji += jamOvertime * gajiPerJamOvertime
	fmt.Printf("Gaji dengan overtime: %d\n", gaji)
}

// Fungsi iteratif untuk mencari karyawan berdasarkan nama
func CariKaryawanBerdasarkanNamaIteratif(karyawan []Karyawan) {
	var nama string
	fmt.Print("Masukkan nama karyawan yang dicari: ")
	fmt.Scanln(&nama)

	// Start running time measurement here
	start := time.Now()

	hasil := cariNamaIteratif(nama, karyawan)
	if hasil == -1 {
		fmt.Println("Karyawan tidak ditemukan")
	} else {
		fmt.Printf("Gaji karyawan %s ditemukan: %d\n", nama, hasil)
	}

	// Print running time for this function in standard format (seconds)
	duration := time.Since(start)
	fmt.Printf("Running time CariKaryawanBerdasarkanNamaIteratif: mikro detik\n", duration.Microseconds())
}

func cariNamaIteratif(nama string, karyawan []Karyawan) int {
	for _, k := range karyawan {
		if k.nama == nama {
			return HitungGajiManual(k.jabatan, 1)
		}
	}
	return -1
}

// Fungsi rekursif untuk menghapus karyawan berdasarkan nama
func HapusKaryawanRekursif(karyawan *[]Karyawan) {
	var nama string
	fmt.Print("Masukkan nama karyawan yang ingin dihapus: ")
	fmt.Scanln(&nama)

	hasil := hapusNamaRekursif(nama, karyawan, 0)
	if !hasil {
		fmt.Println("Karyawan tidak ditemukan")
	} else {
		fmt.Printf("Karyawan %s berhasil dihapus\n", nama)
	}
}

func hapusNamaRekursif(nama string, karyawan *[]Karyawan, indeks int) bool {
	if indeks >= len(*karyawan) {
		return false
	}
	if (*karyawan)[indeks].nama == nama {
		*karyawan = append((*karyawan)[:indeks], (*karyawan)[indeks+1:]...)
		return true
	}
	return hapusNamaRekursif(nama, karyawan, indeks+1)
}

func main() {
	var karyawan []Karyawan

	// Input jumlah CEO
	var jumlahCEO int
	fmt.Print("Masukkan jumlah CEO: ")
	fmt.Scanln(&jumlahCEO)
	// Input data CEO
	for i := 0; i < jumlahCEO; i++ {
		var nama string
		fmt.Printf("Masukkan Nama CEO ke-%d: ", i+1)
		fmt.Scanln(&nama)
		karyawan = append(karyawan, Karyawan{nama: nama, jabatan: "CEO"})
	}

	// Input jumlah Direktur
	var jumlahDirektur int
	fmt.Print("Masukkan jumlah Direktur: ")
	fmt.Scanln(&jumlahDirektur)
	// Input data Direktur
	for i := 0; i < jumlahDirektur; i++ {
		var nama string
		fmt.Printf("Masukkan Nama Direktur ke-%d: ", i+1)
		fmt.Scanln(&nama)
		karyawan = append(karyawan, Karyawan{nama: nama, jabatan: "Direktur"})
	}

	// Input jumlah Manajer
	var jumlahManajer int
	fmt.Print("Masukkan jumlah Manajer: ")
	fmt.Scanln(&jumlahManajer)
	// Input data Manajer
	for i := 0; i < jumlahManajer; i++ {
		var nama string
		fmt.Printf("Masukkan Nama Manajer ke-%d: ", i+1)
		fmt.Scanln(&nama)
		karyawan = append(karyawan, Karyawan{nama: nama, jabatan: "Manajer"})
	}

	// Input jumlah Pekerja
	var jumlahPekerja int
	fmt.Print("Masukkan jumlah Pekerja: ")
	fmt.Scanln(&jumlahPekerja)
	// Input data Pekerja
	for i := 0; i < jumlahPekerja; i++ {
		var nama string
		fmt.Printf("Masukkan Nama Pekerja ke-%d: ", i+1)
		fmt.Scanln(&nama)
		karyawan = append(karyawan, Karyawan{nama: nama, jabatan: "Pekerja"})
	}

	// Menu interaktif
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan daftar nama karyawan")
		fmt.Println("2. Hitung gaji")
		fmt.Println("3. Hitung gaji berdasarkan performa")
		fmt.Println("4. Hitung gaji overtime")
		fmt.Println("5. Cari karyawan berdasarkan nama")
		fmt.Println("6. Hapus karyawan berdasarkan nama")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			TampilkanDaftarNama(karyawan)
		case 2:
			HitungGaji()
		case 3:
			HitungGajiPerforma()
		case 4:
			HitungGajiOvertime()
		case 5:
			CariKaryawanBerdasarkanNamaIteratif(karyawan)
		case 6:
			HapusKaryawanRekursif(&karyawan)
		case 7:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}