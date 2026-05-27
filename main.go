package main

import "fmt"

func main() {
	var portofolio [NMAX]Investasi = dummyData
	var jumlahData int = 5

	var menu int
	var idx, idTarget int
	var keyword string

	for {
		tampilkanMenu()
		fmt.Scan(&menu)

		if menu == 0 {
			fmt.Println("\nTerima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			break
		}

		switch menu {
		case 1:
			if jumlahData == 0 {
				fmt.Println("\nPortofolio masih kosong. Silakan tambah data terlebih dahulu.")
			} else {
				tampilkanTabel(portofolio, jumlahData)
			}

		case 2:
			fmt.Println("\n--- TAMBAH DATA ASET ---")
			tambahAset(&portofolio, &jumlahData)

		case 3:
			fmt.Println("\n--- UBAH DATA ASET ---")
			fmt.Print("Masukkan ID Aset yang ingin diubah harganya: ")
			fmt.Scan(&idTarget)

		case 4:
			fmt.Println("\n--- HAPUS DATA ASET ---")
			fmt.Print("Masukkan ID Aset yang ingin dihapus: ")
			fmt.Scan(&idTarget)
			hapusAset(&portofolio, &jumlahData, idTarget)

		case 5:
			fmt.Println("\n--- CARI DATA ASET (SEQUENTIAL) ---")
			fmt.Print("Masukkan nama aset yang dicari: ")
			fmt.Scan(&keyword)
			
			idx = sequentialSearchNama(portofolio, jumlahData, keyword)
			if idx >= 0 {
				fmt.Printf("Sukses: Aset ditemukan pada indeks array ke-%d.\n", idx)
			} else {
				fmt.Println("Gagal: Aset dengan nama tersebut tidak ditemukan.")
			}

		case 6:
			fmt.Println("\n--- URUTKAN PORTOFOLIO (SELECTION SORT) ---")
			selectionSortNilaiInvestasi(&portofolio, jumlahData, false)
			fmt.Println("Sukses: Portofolio berhasil diurutkan berdasarkan Nilai Investasi (Ascending).")

		default:
			fmt.Println("\nPilihan salah! Silakan masukkan angka menu yang valid (0-6).")
		}

		fmt.Println()
	}
}