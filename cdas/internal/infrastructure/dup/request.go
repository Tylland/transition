package dup

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Request struct {
	AuthenticationKey string
	Queries           []*Query
}

func NewRequest(key string) *Request {
	return &Request{
		AuthenticationKey: key,
	}
}

func (r *Request) Query(q *Query) *Request {
	r.Queries = append(r.Queries, q)
	return r
}

func (r *Request) XML() ([]byte, error) {
	var buf bytes.Buffer

	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	start := xml.StartElement{
		Name: xml.Name{Local: "REQUEST"},
	}

	if err := enc.EncodeToken(start); err != nil {
		return nil, err
	}

	login := xml.StartElement{
		Name: xml.Name{Local: "LOGIN"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "authenticationkey"},
				Value: r.AuthenticationKey,
			},
		},
	}

	enc.EncodeToken(login)
	enc.EncodeToken(login.End())

	for _, q := range r.Queries {
		if err := writeQuery(enc, q); err != nil {
			return nil, err
		}
	}

	enc.EncodeToken(start.End())

	if err := enc.Flush(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func writeExpression(enc *xml.Encoder, expr Expression) error {

	switch e := expr.(type) {
	case Comparison:
		start := xml.StartElement{
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

		enc.EncodeToken(start)
		enc.EncodeToken(start.End())
	case Logical:
		start := xml.StartElement{
			Name: xml.Name{Local: e.Operator},
		}

		enc.EncodeToken(start)

		for _, child := range e.Children {
			if err := writeExpression(enc, child); err != nil {
				return err
			}
		}

		enc.EncodeToken(start.End())
	}

	return nil
}
