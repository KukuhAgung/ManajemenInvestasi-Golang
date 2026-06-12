package main

const NMAX int = 100

type Investasi struct{
	id int
	namaAset string
	jenisAset string
	modalAwal float64
	hargaTerkini float64
	nilaiInvestasi float64
	keuntunganRupiah float64
	persentaseKeuntungan float64
}

var dummyData = [NMAX]Investasi{
	{
		id:                   1,
		namaAset:             "BBCA",
		jenisAset:            "Saham",
		modalAwal:            10000000.0,
		hargaTerkini:         12500000.0,
		nilaiInvestasi:       12500000.0,
		keuntunganRupiah:     2500000.0,
		persentaseKeuntungan: 25.0,
	},
	{
		id:                   2,
		namaAset:             "GOTO",
		jenisAset:            "Saham",
		modalAwal:            5000000.0,
		hargaTerkini:         2000000.0,
		nilaiInvestasi:       2000000.0,
		keuntunganRupiah:     -3000000.0, 
		persentaseKeuntungan: -60.0,     
	},
	{
		id:                   3,
		namaAset:             "Sucorinvest Equity Fund",
		jenisAset:            "ReksaDana",
		modalAwal:            2000000.0,
		hargaTerkini:         2300000.0,
		nilaiInvestasi:       2300000.0,
		keuntunganRupiah:     300000.0,
		persentaseKeuntungan: 15.0,
	},
	{
		id:                   4,
		namaAset:             "ORI023",
		jenisAset:            "Obligasi",
		modalAwal:            20000000.0,
		hargaTerkini:         20500000.0,
		nilaiInvestasi:       20500000.0,
		keuntunganRupiah:     500000.0,
		persentaseKeuntungan: 2.5,
	},
	{
		id:                   5,
		namaAset:             "BMRI",
		jenisAset:            "Saham",
		modalAwal:            15000000.0,
		hargaTerkini:         18000000.0,
		nilaiInvestasi:       18000000.0,
		keuntunganRupiah:     3000000.0,
		persentaseKeuntungan: 20.0,
	},
}