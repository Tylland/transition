package dup

import (
	"encoding/xml"
	"strconv"
)

type Query struct {
	ObjectType    string
	NamespaceName string
	SchemaVer     string
	LimitValue    int
	FilterExpr    Expression
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

func writeQuery(enc *xml.Encoder, q *Query) error {
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

	if err := enc.EncodeToken(query); err != nil {
		return err
	}

	if q.FilterExpr != nil {
		filter := xml.StartElement{
			Name: xml.Name{Local: "FILTER"},
		}

		if err := enc.EncodeToken(filter); err != nil {
			return err
		}

		if err := writeExpression(enc, q.FilterExpr); err != nil {
			return err
		}

		if err := enc.EncodeToken(filter.End()); err != nil {
			return err
		}
	}

	return enc.EncodeToken(query.End())
}
