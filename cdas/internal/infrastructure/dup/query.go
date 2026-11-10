package dup

import (
	"encoding/xml"
	"strconv"
)

type Query struct {
	ObjectType    string     `xml:"objecttype,attr"`
	NamespaceName string     `xml:"namespace,attr,omitempty"`
	SchemaVer     string     `xml:"schemaversion,attr,omitempty"`
	LimitValue    int        `xml:"limit,attr,omitempty"`
	FilterExpr    Expression `xml:"FILTER"`
}

func NewQuery(objectType string) *Query {
	return &Query{
		ObjectType: objectType,
	}
}

func (q *Query) Namespace(v string) *Query {
	q.NamespaceName = v
	return q
}

func (q *Query) SchemaVersion(v string) *Query {
	q.SchemaVer = v
	return q
}

func (q *Query) Limit(v int) *Query {
	q.LimitValue = v
	return q
}

func (q *Query) Filter(expr Expression) *Query {
	q.FilterExpr = expr
	return q
}

func writeQuery(w RequestWriter, q *Query) {
	attrs := []xml.Attr{
		{
			Name:  xml.Name{Local: "objecttype"},
			Value: q.ObjectType,
		},
	}

	if q.NamespaceName != "" {
		attrs = append(attrs, xml.Attr{
			Name:  xml.Name{Local: "namespace"},
			Value: q.NamespaceName,
		})
	}

	if q.SchemaVer != "" {
		attrs = append(attrs, xml.Attr{
			Name:  xml.Name{Local: "schemaversion"},
			Value: q.SchemaVer,
		})
	}

	if q.LimitValue > 0 {
		attrs = append(attrs, xml.Attr{
			Name:  xml.Name{Local: "limit"},
			Value: strconv.Itoa(q.LimitValue),
		})
	}

	query := xml.StartElement{
		Name: xml.Name{Local: "QUERY"},
		Attr: attrs,
	}

	w.Start(query)

	if q.FilterExpr != nil {
		filter := xml.StartElement{
			Name: xml.Name{Local: "FILTER"},
		}

		w.Start(filter)

		writeExpression(w, q.FilterExpr)

		w.End(filter.End())
	}

	w.End(query.End())
}
