package student

import (
	"fmt"
)

type StudentBiodata struct {
	StudentID int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func GetStudentById(studentBiodata StudentBiodata) {
	fmt.Println("Nama :", studentBiodata.Nama)
	fmt.Println("Alamat :", studentBiodata.Alamat)
	fmt.Println("Pekerjaan :", studentBiodata.Pekerjaan)
	fmt.Println("Alasan :", studentBiodata.Alasan)
}
