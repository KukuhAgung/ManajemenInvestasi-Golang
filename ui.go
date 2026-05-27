package main

import "fmt"

func tampilkanMenu() {
	fmt.Println("======================================================")
	fmt.Println("       📈 APLIKASI MANAJEMEN INVESTASI 📉")
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


func tampilkanTabel(portofolio [NMAX]Investasi, jumlahData int) {
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