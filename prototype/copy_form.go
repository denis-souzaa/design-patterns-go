package prototype

type CopyForm struct {
	FormRepo FormRepository
}

type Input struct {
	FromFormId     string
	NewFormId      string
	NewCategory    string
	NewDescription string
}

func NewCopyForm(formRepo FormRepository) *CopyForm {
	return &CopyForm{FormRepo: formRepo}
}

func (cf *CopyForm) Execute(i Input) {
	form := cf.FormRepo.GetById(i.FromFormId)
	newForm := form.Clone()
	newForm.FormId = i.NewFormId
	newForm.Category = i.NewCategory
	newForm.Description = i.NewDescription
	cf.FormRepo.Save(*newForm)
}
