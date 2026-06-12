package main

import "fmt"

func main() {
	var portofolio tabInvestasi = dummyData
	var jumlahData int = 5

	var menu int
	var idx, idTarget int
	var keyword string
	var lanjut bool

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
				tampilkanStatistik(portofolio, jumlahData)
			}

		case 2:
			fmt.Println("\n--- TAMBAH DATA ASET ---")
			tambahAset(&portofolio, &jumlahData)
			tampilkanTabel(portofolio, jumlahData)

		case 3:
			fmt.Println("\n--- UBAH DATA ASET ---")
			fmt.Print("Masukkan ID Aset yang ingin diubah harganya: ")
			fmt.Scan(&idTarget)
			ubahAset(&portofolio, jumlahData, idTarget)
			tampilkanTabel(portofolio, jumlahData)

		case 4:
			fmt.Println("\n--- HAPUS DATA ASET ---")
			fmt.Print("Masukkan ID Aset yang ingin dihapus: ")
			fmt.Scan(&idTarget)
			hapusAset(&portofolio, &jumlahData, idTarget)
			tampilkanTabel(portofolio, jumlahData)

		case 5:
			fmt.Println("\n--- CARI DATA ASET ---")
			fmt.Println("1. Cari Berdasarkan Nama (Sequential Search)")
			fmt.Println("2. Cari Berdasarkan Jenis (Binary Search)")
			fmt.Print("Pilih metode pencarian (1/2): ")
			var subMenu int
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				fmt.Print("Masukkan Nama Aset yang dicari: ")
				fmt.Scan(&keyword)
				idx = sequentialSearchNama(portofolio, jumlahData, keyword)
				if idx >= 0 {
					fmt.Printf("Sukses: Aset ditemukan pada indeks array ke-%d.\n", idx)
				} else {
					fmt.Println("Gagal: Aset dengan nama tersebut tidak ditemukan.")
				}
			} else if subMenu == 2 {
				fmt.Print("Masukkan Jenis Aset yang dicari: ")
				fmt.Scan(&keyword)
				
				sortUntukBinarySearch(&portofolio, jumlahData)
				
				idx = binarySearchJenis(portofolio, jumlahData, keyword)
				if idx >= 0 {
					fmt.Printf("Sukses: Aset ditemukan. Perhatikan bahwa tabel mungkin telah diurutkan ulang berdasarkan Jenis.\n")
				} else {
					fmt.Println("Gagal: Aset dengan jenis tersebut tidak ditemukan.")
				}
			} else {
				fmt.Println("Pilihan pencarian tidak valid.")
			}

		case 6:
			fmt.Println("\n--- URUTKAN PORTOFOLIO ---")
			fmt.Println("1. Berdasarkan Nilai Investasi (Selection Sort)")
			fmt.Println("2. Berdasarkan Persentase Profit (Insertion Sort)")
			fmt.Print("Pilih metode pengurutan (1/2): ")
			var subMenu int
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				selectionSortNilaiInvestasi(&portofolio, jumlahData, false) // false = Ascending
				fmt.Println("Sukses: Portofolio berhasil diurutkan berdasarkan Nilai Investasi (Terkecil ke Terbesar).")
				tampilkanTabel(portofolio, jumlahData)
			} else if subMenu == 2 {
				insertionSortPersentase(&portofolio, jumlahData, true) // true = Descending (Profit tertinggi di atas)
				fmt.Println("Sukses: Portofolio berhasil diurutkan berdasarkan Persentase Profit (Terbesar ke Terkecil).")
				tampilkanTabel(portofolio, jumlahData)
			} else {
				fmt.Println("Pilihan pengurutan tidak valid.")
			}
		default:
			fmt.Println("\nPilihan salah! Silakan masukkan angka menu yang valid (0-6).")
		}

		if menu != 0 { 
			lanjut = ulangProgram()
			if !lanjut {
				break
			}
		}
	}
}