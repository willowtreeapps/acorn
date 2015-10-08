package acorn

import (
	"io"
	"text/template"
)

// Func defines the interface necessary to generate a func
type Func interface {
	// The Type the func should be applied to.
	// e.g. use "" for a package func, "(t *MyType)" for your type
	Type() string

	// The Signature of the function.
	// This is what would be defined in an interface.
	Signature() string

	// The Body of the function
	Body() string
}

// Interface defines the components necessary to generate an interface
type Interface interface {
	// The Name of the interface
	Name() string

	// The Signatures of the methods the interface defines
	Signatures() []string
}

// WriteFunction outputs a generated func to the writer
func WriteFunction(w io.Writer, f Func) {
	fmap := map[string]interface{}{
		"type":      f.Type,
		"signature": f.Signature,
		"body":      f.Body,
	}
	const pattern = `
    func {{ type }} {{ signature }} {
      {{ body }}
    }
  `
	templ := template.Must(template.New("func").Funcs(fmap).Parse(pattern))
	templ.Execute(w, nil)
}

// WriteInterface outputs a generated interface to the writer
func WriteInterface(w io.Writer, i Interface) {
	fmap := map[string]interface{}{
		"name":       i.Name,
		"signatures": i.Signatures,
	}
	const pattern = `
    type {{ name }} interface {
      {{ range signatures }}
        {{ . }}
      {{ end }}
    }
  `
	templ := template.Must(template.New("interface").Funcs(fmap).Parse(pattern))
	templ.Execute(w, nil)
}
