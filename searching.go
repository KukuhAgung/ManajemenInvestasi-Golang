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

func sortUntukBinarySearch(portofolio *[NMAX]Investasi, jumlahData int) {
	var i, j int
	for i = 0; i < jumlahData-1; i++ {
		for j = 0; j < jumlahData-i-1; j++ {
			if (*portofolio)[j].jenisAset > (*portofolio)[j+1].jenisAset {
				(*portofolio)[j], (*portofolio)[j+1] = (*portofolio)[j+1], (*portofolio)[j]
			}
		}
	}
}

func binarySearchJenis(portofolio [NMAX]Investasi, jumlahData int, keyword string) int {
	var left, right, mid int
	left = 0
	right = jumlahData - 1

	for left <= right {
		mid = left + (right - left) / 2

		if portofolio[mid].jenisAset == keyword {
			return mid
		}
		
		if portofolio[mid].jenisAset < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}