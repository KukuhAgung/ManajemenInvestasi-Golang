package main

import "fmt"

func tambahAset(portfolio *tabInvestasi, jumlahData *int) {
	var n, i, indeksBaru int

	fmt.Print("Berapa banyak aset yang ingin ditambahkan? ")
	fmt.Scan(&n)

	if *jumlahData + n > NMAX {
		fmt.Println("Gagal: Kapasitas penyimpanan portofolio tidak mencukupi!")
		return
	}
	for i = 0; i < n; i++ {
		indeksBaru = *jumlahData
		
		fmt.Printf("\n--- Input Data ke-%d ---\n", i+1)
		(*portfolio)[indeksBaru].id = *jumlahData + 1

		fmt.Print("Masukkan Nama Aset : ")
		fmt.Scan(&(*portfolio)[indeksBaru].namaAset)
		fmt.Print("Masukkan Jenis Aset: ")
		fmt.Scan(&(*portfolio)[indeksBaru].jenisAset)
		fmt.Print("Masukkan Modal Awal: ")
		fmt.Scan(&(*portfolio)[indeksBaru].modalAwal)

		(*portfolio)[indeksBaru].hargaTerkini = (*portfolio)[indeksBaru].modalAwal
		(*portfolio)[indeksBaru].nilaiInvestasi = (*portfolio)[indeksBaru].modalAwal
		(*portfolio)[indeksBaru].keuntunganRupiah = 0.0
		(*portfolio)[indeksBaru].persentaseKeuntungan = 0.0

		hitungKeuntungan(&(*portfolio)[indeksBaru])

		*jumlahData++
	}
	fmt.Println("\nSukses: Seluruh aset baru berhasil ditambahkan.")
}

func hapusAset(portfolio *tabInvestasi, jumlahData *int, id int) {
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

func ubahAset(portofolio *tabInvestasi, jumlahData int, id int) {
	var i int
	var ditemukan bool = false

	for i = 0; i < jumlahData; i++ {
		if portofolio[i].id == id {
			ditemukan = true
			fmt.Printf("Aset ditemukan: %s\n", portofolio[i].namaAset)
			
			var jenisBaru string
			fmt.Print("Masukkan Jenis Aset baru (opsional, ketik '-' jika tidak diubah): ")
			fmt.Scan(&jenisBaru)
			if jenisBaru != "-" {
				portofolio[i].jenisAset = jenisBaru
			}

			var hargaBaru float64
			fmt.Print("Masukkan Harga Terkini baru: ")
			fmt.Scan(&hargaBaru)
			portofolio[i].hargaTerkini = hargaBaru
			portofolio[i].nilaiInvestasi = hargaBaru

			hitungKeuntungan(&portofolio[i])
			
			fmt.Println("Data aset berhasil diperbarui!")
			break
		}
	}

	if !ditemukan {
		fmt.Println("Aset dengan ID tersebut tidak ditemukan.")
	}
}