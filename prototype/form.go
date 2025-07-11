package prototype

type Form struct {
	FormId      string
	Category    string
	Description string
	Fields      []Field
}

func NewForm(formId, category, description string) *Form {
	return &Form{FormId: formId, Category: category, Description: description}
}

func (f *Form) AddField(tp, title string) {
	f.Fields = append(f.Fields, *Create(tp, title))
}

func (f *Form) Clone() *Form {
	var fields []Field
	for _, fs := range f.Fields {
		fields = append(fields, *fs.Clone())
	}

	return &Form{
		FormId:      f.FormId,
		Category:    f.Category,
		Description: f.Description,
		Fields:      fields,
	}
}
