package main

import (
	"os"
	"text/template"

	"github.com/maja42/goval"
)

func main() {
	// Init evaluator
	eval := goval.NewEvaluator()

	// Data
	data := map[string]interface{}{
		"isTrue": false,
		"state": map[string]interface{}{
			"id": "asdf",
		},
	}

	// Create template
	tmpl := template.New("test")

	// Add function to template
	tmpl.Funcs(template.FuncMap{
		"eval": func(expr string, v map[string]interface{}) (any, error) {
			return eval.Evaluate(expr, v, nil)
		},
	})

	// Parse template
	tmpl.Parse(`
		Evaluator: {{eval "isTrue == true" .}}
		JSON {{eval "state.id != \"\"" . }} 
	`)

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
