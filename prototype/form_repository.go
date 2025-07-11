package prototype

type FormRepository interface {
	Save(f Form)
	GetById(formId string) *Form
}

type FormRepositoryMemory struct {
	Forms []Form
}

func NewFormRepositoryMemory() *FormRepositoryMemory {
	return &FormRepositoryMemory{}
}

func (fm *FormRepositoryMemory) Save(form Form) {
	fm.Forms = append(fm.Forms, form)
}

func (fm *FormRepositoryMemory) GetById(formId string) *Form {
	for _, f := range fm.Forms {
		if f.FormId == formId {
			return &f
		}
	}
	return nil
}
