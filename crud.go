package main

import "fmt"

func tambahAset(portfolio *[NMAX]Investasi, jumlahData *int) {
	var i int
	if *jumlahData >= NMAX {
		fmt.Println("Gagal: Kapasitas penyimpanan portofolio sudah penuh!")
		return
	}

	i = *jumlahData

	(*portfolio)[i].id = *jumlahData + 1

	fmt.Print("Masukkan Nama Aset : ")
	fmt.Scan(&(*portfolio)[i].namaAset)
	fmt.Print("Masukkan Jenis Aset: ")
	fmt.Scan(&(*portfolio)[i].jenisAset)
	fmt.Print("Masukkan Modal Awal: ")
	fmt.Scan(&(*portfolio)[i].modalAwal)

	(*portfolio)[i].hargaTerkini = (*portfolio)[i].modalAwal
	(*portfolio)[i].nilaiInvestasi = (*portfolio)[i].modalAwal
	(*portfolio)[i].keuntunganRupiah = 0.0
	(*portfolio)[i].persentaseKeuntungan = 0.0

	*jumlahData++
	fmt.Println("Sukses: Aset baru berhasil ditambahkan.")
}

func hapusAset(portfolio *[NMAX]Investasi, jumlahData *int, id int) {
	var idx int
	idx = sequentialSearchID(*portfolio, *jumlahData, id)

	if idx < 0 {
		fmt.Println("Gagal: ID data tidak ditemukan.")
		return 
	}

	for i := idx; i < *jumlahData-1; i++ {
		(*portfolio)[i] = (*portfolio)[i+1]
	}

	*jumlahData--

	fmt.Println("Sukses: Aset berhasil dihapus secara manual.")
}