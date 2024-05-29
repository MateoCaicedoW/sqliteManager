package render

//this package is used to render the html templates
import (
	"html/template"
	"net/http"
)

var dataMap = make(map[string]any)

func SetData(key string, value any) {
	dataMap[key] = value
}

func GetData(key string) any {
	return dataMap[key]
}

// this function is used to render the html templates
func HTML(w http.ResponseWriter, tmpl string) {
	// ParseFiles parses the named files and associates the resulting templates with t.
	// If an error occurs, parsing stops and the returned template is nil;
	// otherwise it is t. There must be at least one file.
	t, err := template.ParseFiles("./" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute applies a parsed template to the specified data object,
	// writing the output to wr.
	err = t.Execute(w, dataMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderWithLayout(w http.ResponseWriter, tmpl string, layout string) {
	// ParseFiles parses the named files and associates the resulting templates with t.
	// If an error occurs, parsing stops and the returned template is nil;
	// otherwise it is t. There must be at least one file.
	t, err := template.ParseFiles("./"+tmpl, "./layouts/"+layout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute applies a parsed template to the specified data object,
	// writing the output to wr.
	err = t.ExecuteTemplate(w, layout, dataMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
