package templates

import (
	dtos "Forum-back/pkg/dtos/templates"
	"html/template"
)

func GetTemplateWithLayout(headerDto *dtos.HeaderDto, pageName string, templatePath ...string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(templatePath...)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.ParseFiles("internal/templates/components/headerComponent.gohtml")
	if err != nil {
		return nil, err
	}

	headerDto.PageName = pageName

	return tmpl, nil
}
