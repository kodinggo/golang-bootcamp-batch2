package user

import (
	"fmt"
	"go-project/subject"
)

// Interface mahasiswa
type UserMahasiswa interface {
	SetNama(nama string)
	GetProfile() string
}

// Object mahasiswa
type mahasiswa struct {
	ID   int64
	Nama string

	SubjectMtk subject.SubjectMatemtika
}

func (m *mahasiswa) SetNama(nama string) {
	m.Nama = nama
}

func (m *mahasiswa) GetProfile() string {
	return fmt.Sprintf("Halo, nama saya %s, saya suka pelajaran %s", m.Nama, m.SubjectMtk.GetSubjectName())
}

func NewMahasiswa(SubjectMtk subject.SubjectMatemtika) UserMahasiswa {
	return &mahasiswa{SubjectMtk: SubjectMtk}
}
