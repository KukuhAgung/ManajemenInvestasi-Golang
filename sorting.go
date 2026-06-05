package main

func selectionSortNilaiInvestasi(portfolio *[NMAX]Investasi, jumlahData int, isDescending bool) {
	var i, j, targetIdx int

	for i = 0; i < jumlahData-1; i++ {
		targetIdx = i

		for j = i + 1; j < jumlahData; j++ {
			if !isDescending {
				if (*portfolio)[j].nilaiInvestasi < (*portfolio)[targetIdx].nilaiInvestasi {
					targetIdx = j
				}
			} else {
				if (*portfolio)[j].nilaiInvestasi > (*portfolio)[targetIdx].nilaiInvestasi {
					targetIdx = j
				}
			}
		}

		if targetIdx != i {
			(*portfolio)[i], (*portfolio)[targetIdx] = (*portfolio)[targetIdx], (*portfolio)[i]
		}
	}
}

func insertionSortPersentase(portofolio *[NMAX]Investasi, jumlahData int, isDescending bool) {
	var i, j int
	var key Investasi

	for i = 1; i < jumlahData; i++ {
		key = (*portofolio)[i]
		j = i - 1

		if isDescending {
			for j >= 0 && (*portofolio)[j].persentaseKeuntungan < key.persentaseKeuntungan {
				(*portofolio)[j+1] = (*portofolio)[j]
				j--
			}
		} else {
			for j >= 0 && (*portofolio)[j].persentaseKeuntungan > key.persentaseKeuntungan {
				(*portofolio)[j+1] = (*portofolio)[j]
				j--
			}
		}
		(*portofolio)[j+1] = key
	}
}