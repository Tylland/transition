package dup

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type XMLRequestWriter struct {
	prefix      string
	indent      string
	indentCount int
	builder     *strings.Builder
}

func NewXMLRequestWriter(prefix string, indent string) *XMLRequestWriter {
	return &XMLRequestWriter{prefix: prefix, indent: indent, builder: &strings.Builder{}}
}

func (w *XMLRequestWriter) write(format string, a ...any) {
	w.builder.WriteString(strings.Repeat(w.indent, w.indentCount))
	w.builder.WriteString(fmt.Sprintf(format, a...))
}

func (w *XMLRequestWriter) writeln(format string, a ...any) {
	w.builder.WriteString(strings.Repeat(w.indent, w.indentCount))
	w.builder.WriteString(fmt.Sprintf(format+"\n", a...))
}

func (w *XMLRequestWriter) attr(elem xml.StartElement) string {
	var attrs []string

	for _, att := range elem.Attr {
		attrs = append(attrs, fmt.Sprintf("%s=%q", att.Name.Local, att.Value))
	}

	if len(attrs) == 0 {
		return ""
	}

	return fmt.Sprintf(" %s", strings.Join(attrs, " "))
}

func (w *XMLRequestWriter) Inline(elem xml.StartElement) {
	w.writeln("<%s%s />", elem.Name.Local, w.attr(elem))
}

func (w *XMLRequestWriter) Start(elem xml.StartElement) func() {
	w.writeln("<%s%s>", elem.Name.Local, w.attr(elem))

	w.indentCount++

	return func() {
		w.End(elem.End())
	}
}

func (w *XMLRequestWriter) End(elem xml.EndElement) {
	w.indentCount--

	if w.indentCount == 0 {
		w.write("</%s>", elem.Name.Local)
	} else {
		w.writeln("</%s>", elem.Name.Local)
	}
}
