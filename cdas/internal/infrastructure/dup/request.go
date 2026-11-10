package dup

import (
	"encoding/xml"
	"fmt"
)

type RequestWriter interface {
	Inline(elem xml.StartElement)
	Start(elem xml.StartElement) func()
	End(elem xml.EndElement)
}

type Login struct {
	AuthenticationKey string `xml:"nickname,omitempty"`
}

type Request struct {
	Login   *Login `xml:"LOGIN,omitempty"`
	Queries []*Query
}

func NewRequest(key string) *Request {
	return &Request{Login: &Login{AuthenticationKey: key}}
}

func (r *Request) Query(q *Query) *Request {
	r.Queries = append(r.Queries, q)
	return r
}

func (r *Request) XML() ([]byte, error) {

	w := NewXMLRequestWriter("", "  ")

	request := xml.StartElement{
		Name: xml.Name{Local: "REQUEST"},
	}

	w.Start(request)

	login := xml.StartElement{
		Name: xml.Name{Local: "LOGIN"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "authenticationkey"},
				Value: r.Login.AuthenticationKey,
			},
		},
	}

	w.Inline(login)

	for _, q := range r.Queries {
		writeQuery(w, q)
	}

	w.End(request.End())

	return []byte(w.builder.String()), nil
}

func writeExpression(w RequestWriter, expr Expression) {

	switch e := expr.(type) {
	case Comparison:
		comp := xml.StartElement{
			Name: xml.Name{Local: e.Operator},
			Attr: []xml.Attr{
				{
					Name:  xml.Name{Local: "name"},
					Value: e.Name,
				},
				{
					Name:  xml.Name{Local: "value"},
					Value: fmt.Sprintf("%v", e.Value),
				},
			},
		}

		w.Inline(comp)
	case Logical:
		logic := xml.StartElement{
			Name: xml.Name{Local: e.Operator},
		}

		w.Start(logic)

		for _, child := range e.Children {
			writeExpression(w, child)
		}

		w.End(logic.End())
	}
}
