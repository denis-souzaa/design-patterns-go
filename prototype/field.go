package prototype

import "github.com/google/uuid"

type Field struct {
	FieldId string
	Type    string
	Title   string
}

func Create(tp, title string) *Field {
	fieldId := uuid.New().String()
	return &Field{FieldId: fieldId, Title: title, Type: tp}
}

func (f *Field) Clone() *Field {
	return &Field{FieldId: f.FieldId, Type: f.Type, Title: f.Title}
}
