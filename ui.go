package main

import "fmt"

func tampilkanMenu() {
	fmt.Println("======================================================")
	fmt.Println("            APLIKASI MANAJEMEN INVESTASI 			")
	fmt.Println("======================================================")
	fmt.Println("[1] Lihat Portofolio & Laporan Statistik")
	fmt.Println("[2] Tambah Data Investasi Baru")
	fmt.Println("[3] Ubah Data Investasi (Update Harga)")
	fmt.Println("[4] Hapus Data Investasi")
	fmt.Println("[5] Cari Aset (Berdasarkan Nama / Jenis)")
	fmt.Println("[6] Urutkan Portofolio (Berdasarkan Nilai / Profit)")
	fmt.Println("[0] Keluar Aplikasi")
	fmt.Println("======================================================")
	fmt.Print("Pilih menu (0-6): ")
}

func tampilkanTabel(portofolio tabInvestasi, jumlahData int) {
	var i int
	fmt.Println("\n==========================================================================================================")
	fmt.Printf("%-4s | %-30s | %-11s | %-15s | %-15s | %s\n", 
        "ID", "NAMA ASET", "JENIS", "MODAL AWAL", "HARGA TERKINI", "PROFIT/LOSS (%)")
	fmt.Println("-----------------------------------------------------------------------------------------------------")

	for i=0; i < jumlahData; i++{
		fmt.Printf("%-4d | %-30s | %-11s | Rp %-12.0f | Rp %-12.0f | %+.2f%%\n",
			portofolio[i].id,
			portofolio[i].namaAset,
			portofolio[i].jenisAset,
			portofolio[i].modalAwal,
			portofolio[i].hargaTerkini,
			portofolio[i].persentaseKeuntungan,
		)
	}
	fmt.Println("==========================================================================================================\n")
}

func tampilkanStatistik(portofolio tabInvestasi, jumlahData int) {
	var totalModal, totalNilaiTerkini, totalKeuntungan float64
	var i int

	for i = 0; i < jumlahData; i++ {
		totalModal += portofolio[i].modalAwal
		totalNilaiTerkini += portofolio[i].hargaTerkini
		totalKeuntungan += portofolio[i].keuntunganRupiah
	}

	fmt.Println("================ RINGKASAN STATISTIK ===============")
	fmt.Printf("Total Modal Keseluruhan : Rp %.2f\n", totalModal)
	fmt.Printf("Total Nilai Terkini     : Rp %.2f\n", totalNilaiTerkini)
	fmt.Printf("Total Keuntungan/Rugi   : Rp %.2f\n", totalKeuntungan)
	fmt.Println("====================================================\n")
}

func ulangProgram() bool {
	var pilihan int
	fmt.Println("Apakah Anda ingin melakukan aktivitas lain?")
	fmt.Println("[1] Ya, kembali ke menu utama")
	fmt.Println("[0] Tidak, Keluar Aplikasi")
	fmt.Print("Pilih (0/1): ")
	fmt.Scan(&pilihan)

	if pilihan == 0 {
		fmt.Println("\nTerima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
		return false
	}
	
	fmt.Println() 
	return true 
}