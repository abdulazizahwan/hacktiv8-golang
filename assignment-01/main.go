// Abd Aziz
// Scalable Web Service with Golang - DTS Kominfo

package main

import (
	"fmt"
	"os"
	"strconv"
	"assignment-01/student"
)

var studentBiodata []student.StudentBiodata

func init() {
	studentBiodata = []student.StudentBiodata{
		{StudentID: 1, Nama: "Frigyes Antigonos", Alamat: "CGK", Pekerjaan: "Junior BackEnd Engineer", Alasan: "Lorem Ipsum is simply dummy text of the printing and typesetting industry"},
		{StudentID: 2, Nama: "Joele Indrek", Alamat: "SRG", Pekerjaan: "Fullstack Engineer", Alasan: "Contrary to popular belief, Lorem Ipsum is not simply random text"},
		{StudentID: 3, Nama: "Regina Kaelee", Alamat: "SBY", Pekerjaan: "FrontEnd Engineer", Alasan: "The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested"},
		{StudentID: 4, Nama: "Pau Euphranor", Alamat: "PKL", Pekerjaan: "Product Owner", Alasan: "Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text"},
		{StudentID: 5, Nama: "Jonathan Nikolaj", Alamat: "YGY", Pekerjaan: "Mobile Engineer", Alasan: "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium"},
	}
}

func main() {
	allArgs := os.Args[1:]
	if len(os.Args) > 1 {
		for i := 0; i < len(allArgs); i++ {
			id, err := strconv.Atoi(allArgs[i])
			if err != nil {
				fmt.Println("Wrong! Please check your input again")
			} else {
				find := 0
				for i := 0; i < len(studentBiodata); i++ {
					if studentBiodata[i].StudentID == id {
						student.GetStudentById(studentBiodata[i])
						find++
						break
					}
				}
				if find == 0 {
					fmt.Println("404 Not Found | No result for you")
				}
			}
		}
	} else {
		fmt.Println("Please input the StudentID number in Arguments")
	}
}
