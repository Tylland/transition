package specs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

type ExampleBlock struct {
	context *ScenarioContext
}

func Headers(exampleHeaders []interface{}) []ExampleHeader {
	var headers []ExampleHeader

	for i := 0; i < len(exampleHeaders); i += 2 {
		if len(exampleHeaders)-i > 1 {
			title := reflect.ValueOf(exampleHeaders[i]).String()

			headers = append(headers, ExampleHeader{Title: title, variable: exampleHeaders[i+1]})
		}
	}

	return headers
}

type ExampleHeader struct {
	Title    string
	variable interface{}
}

type ExampleTable struct {
	headers []ExampleHeader
	values  [][]any
}

func NewExampleBlock(context *ScenarioContext) *ExampleBlock {
	return &ExampleBlock{context: context}
}

func (eb *ExampleBlock) Example(values ...any) *ExampleBlock {
	eb.context.examples.values = append(eb.context.examples.values, values)

	return eb
}

func (eb *ExampleBlock) TitleOld(variable interface{}) string {
	varValue := reflect.ValueOf(variable)

	if varValue.Kind() == reflect.Ptr {
		varValue = varValue.Elem()
	}

	varPointer := varValue.Addr().Pointer()

	for _, hdr := range eb.context.examples.headers {

		hdrValue := reflect.ValueOf(hdr.variable)

		if hdrValue.Kind() == reflect.Ptr {
			hdrValue = hdrValue.Elem()
		}

		hdrPointer := hdrValue.Addr().Pointer()

		fmt.Printf("Title -Header addr %v \n", hdrPointer)
		fmt.Printf("Title- Variable addr %v \n", varPointer)

		if hdrPointer == varPointer {
			return hdr.Title
		}
	}

	return ""
}

type OutputRecorder struct {
	old     *os.File
	reader  *os.File
	writer  *os.File
	outChan chan string
}

func RecordOutput() func () string {
	r := &OutputRecorder{old: os.Stdout, outChan: make(chan string)}

	r.reader, r.writer, _ = os.Pipe()

	os.Stdout = r.writer

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r.reader)
		r.outChan <- buf.String()
	}()

	return r.Stop
}

func (r *OutputRecorder) Stop() string {
	r.writer.Close()
	os.Stdout = r.old
	out := <-r.outChan

	return out
}

func (eb *ExampleBlock) Execute() {
	defer RecordOutput()

	t := eb.context.t

	examples := eb.context.examples

	for index, v := range examples.values {
		for i := 0; i < len(examples.headers); i++ {
			value := reflect.ValueOf(v[i])

			eb.setValue(examples.headers[i].variable, value)

			if eb.context.debug {
				fmt.Printf("Setting %s to %v \n", examples.headers[i].Title, examples.headers[i].variable)
			}
		}

		res, err := eb.context.Execute(t, index > 0)

		S.WriteExample(v, res, err)
	}

	S.WriteExampleHeaders(examples.headers)
}

func (eb *ExampleBlock) setValue(variable interface{}, newValue reflect.Value) {
	// Get the reflect.Value of the pointer
	v := reflect.ValueOf(variable)

	// Check if it is a pointer and is addressable
	if v.Kind() == reflect.Ptr && v.Elem().CanSet() {
		// Set the new value
		v.Elem().Set(newValue)

		if eb.context.debug {
			fmt.Printf("Updating Variable at addr %v \n", v.Elem().Addr())
		}

	} else {
		if eb.context.debug {
			fmt.Print("Value is not settable")
		}
	}
}
