package subject

type SubjectMatemtika interface {
	GetSubjectName() string
}

type Matematika struct {
	SubjectName string
}

func (m *Matematika) GetSubjectName() string {
	return m.SubjectName
}

func NewMatematika() SubjectMatemtika {
	return &Matematika{
		SubjectName: "Matematika",
	}
}
