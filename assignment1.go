package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Mahasiswa struct {
	Id        string
	Nama      string
	Pekerjaan string
	Alasan    string
}

var muridStudiGolang []Mahasiswa = []Mahasiswa{
	{
		Id:        "1",
		Nama:      "Alvian Ghifari",
		Pekerjaan: "Software Engineer",
		Alasan:    "mau dapet jadi backend golang",
	},
	{
		Id:        "2",
		Nama:      "Afrida Rohmatin",
		Pekerjaan: "Frontend",
		Alasan:    "Belajar backend",
	},
	{
		Id:        "3",
		Nama:      "Hafiz",
		Pekerjaan: "turu",
		Alasan:    "biar gak turu",
	},
}

func FindStudent(studentId string) (Mahasiswa, error) {
	for _, student := range muridStudiGolang {

		if student.Id == studentId {
			return student, nil
		}
	}
	return Mahasiswa{}, errors.New("no student")
}

func main() {
	input := os.Args

	if !(len(input) >= 2) {
		log.Fatalln("Rusak")
	}

	mahasiswa, err := FindStudent(input[1])

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Nama: %v\n", mahasiswa.Nama)
	fmt.Printf("Pekerjaan: %v\n", mahasiswa.Pekerjaan)
	fmt.Printf("Alasan: %v\n", mahasiswa.Alasan)
}
