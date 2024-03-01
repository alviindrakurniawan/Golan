package helper

import "fmt"

func ShowData(listBiodata []Biodata,no int){
	no = int(no)
	if no>0 && no<=len(listBiodata) {
		fmt.Println("Nama:", listBiodata[no-1].Nama)
		fmt.Println("Alamat:", listBiodata[no-1].Alamat)
		fmt.Println("Pekerjaan:", listBiodata[no-1].Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang:", listBiodata[no-1].Alasan)
	}else{
		fmt.Println("Data tidak ada")
	}
	
}