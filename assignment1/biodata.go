package main

import (
	"assignment1/helper"
	"fmt"
	"os"
	"strconv"
)

func main (){
	var peserta= []helper.Biodata{
		{Nama:"Acak", Alamat: "Jln monyet", Pekerjaan:"Pekerjaan 2", Alasan:"Alasan memilih Golang 1"},
		{Nama:"Bambang", Alamat: "Jln Parus", Pekerjaan:"Pekerjaan 2", Alasan:"Alasan memilih Golang 2"},	 
		{"Supri","Jln Padar","Pekerjaan 3","Alasan memilih Golang 3"},
	}
	
	index, err := strconv.Atoi(os.Args[1])
	if err == nil {
		helper.ShowData(peserta,index)
	} else {
		fmt.Printf("%s is not a number.\n",os.Args[1])
	}

	



	

	
	
	// person:= Biodata{name:"bambang",
	// alamat:"Jalan musung",
	// pekerjaan:"Banker",
	// alasan:"Memperdalam ilmu"}

	
	
}
	