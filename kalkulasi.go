package main

func hitungKeuntungan(aset *Investasi) {
	aset.keuntunganRupiah = aset.hargaTerkini - aset.modalAwal

	if aset.modalAwal > 0 {
		aset.persentaseKeuntungan = (aset.keuntunganRupiah / aset.modalAwal) * 100
	} else {
		aset.persentaseKeuntungan = 0
	}
}