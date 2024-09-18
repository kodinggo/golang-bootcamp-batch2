package subject

type SubjectMatemtika interface {
	GetSubjectName() string
}

type matematika struct {
	SubjectName string
}

func (m *matematika) GetSubjectName() string {
	return m.SubjectName
}

func NewMatematika() SubjectMatemtika {
	return &matematika{
		SubjectName: "Matematika",
	}
}
