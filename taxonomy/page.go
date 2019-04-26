package taxonomy

import (
	"html/template"
	"strings"
	"time"
)

type Page struct {
	Filename        string
	CreatedAt       time.Time
	Name            string
	Title           string `yaml:"title"`
	Layout          string `yaml:"layout"`
	MarkDownContent string `fm:"content"`
	HtmlContent     template.HTML
	FullHtml        string
	BlogName        string
}

func (p *Page) IsValid() (isValid bool) {
	isValid = false

	if len(p.GetFilename()) == 0 {
		return isValid
	}

	explodedfn := strings.Split(p.GetFilename(), "-")
	layout := "2006-01-02"

	if len(explodedfn) > 3 {
		dateSlice := explodedfn[:3]
		date := strings.Join(dateSlice, "-")
		_, dateConversionError := time.Parse(layout, date)

		if dateConversionError == nil {
			isValid = true
		}
	}

	return isValid
}

func (p *Page) GetFilename() string {
	return p.Filename
}

func (p *Page) GetFullHtml() string {
	return p.FullHtml
}
