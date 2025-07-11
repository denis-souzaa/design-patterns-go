package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCopyForm(t *testing.T) {
	formRepo := NewFormRepositoryMemory()
	form := NewForm("1", "Marketing", "Leads v1")
	form.AddField("text", "name")
	form.AddField("text", "email")
	formRepo.Save(*form)
	copyForm := NewCopyForm(formRepo)
	input := Input{FromFormId: "1", NewFormId: "2", NewCategory: "Marketing", NewDescription: "Leads V2"}
	copyForm.Execute(input)
	newForm := formRepo.GetById("2")
	assert.Len(t, newForm.Fields, 2)
	assert.Equal(t, "Marketing", newForm.Category)
	assert.Equal(t, "Leads V2", newForm.Description)
	assert.Equal(t, "name", newForm.Fields[0].Title)
	assert.Equal(t, "text", newForm.Fields[0].Type)
}
