package render

//this package is used to render the html templates
import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed *.html
	templates embed.FS
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

	tt, err := tt.ParseFS(templates, partials...)
	if err != nil {
		return err
	}

	return tt.Execute(w, dataMap)

}
