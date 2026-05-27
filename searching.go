package main

func sequentialSearchNama(portofolio [NMAX]Investasi, jumlahData int, keyword string) int {
	var i int
	for i = 0; i < jumlahData; i++ {
		if portofolio[i].namaAset == keyword {
			return i
		}
	}
	return -1
}

func sequentialSearchID(portofolio [NMAX]Investasi, jumlahData int, id int) int {
	var i int
	for i = 0; i < jumlahData; i++ {
		if portofolio[i].id == id {
			return i
		}
	}
	return -1
}