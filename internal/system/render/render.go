package render

//this package is used to render the html templates
import (
	"html/template"
	"io"

	"github.com/MateoCaicedoW/sqliteManager/internal/system/templates"
)

var dataMap = make(map[string]any)

func SetData(key string, value any) {
	dataMap[key] = value
}

func GetData(key string) any {
	return dataMap[key]
}

func Render(w io.Writer, partials ...string) error {
	if len(partials) == 0 {
		return nil
	}

	tt := template.New(partials[0])

	tt, err := tt.ParseFS(templates.Templates, partials...)
	if err != nil {
		return err
	}

	return tt.Execute(w, dataMap)

}
