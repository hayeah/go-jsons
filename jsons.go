package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"
)

// cat json | jsons {1..5} {a..d}

func main() {
	args := os.Args[1:]
	input := os.Stdin
	output := os.Stdout

	decoder := json.NewDecoder(input)

	var v interface{}
	err := decoder.Decode(&v)
	if err != nil {
		fmt.Fprint(os.Stderr, "Invalid json: %v\n", err)
		os.Exit(1)
	}

	var proto interface{}
	var arrayMode bool

	switch v := v.(type) {
	case []interface{}:
		proto = v[0].(map[string]interface{})
		arrayMode = true
	default:
		proto = v
		arrayMode = false
	}

	err = expand(proto, args, output, arrayMode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func expand(obj interface{}, args []string, output io.Writer, array bool) (err error) {
	protoJSON, err := json.Marshal(obj)
	if err != nil {
		return
	}

	tmpl, err := template.New("proto").Parse(string(protoJSON))
	if err != nil {
		return
	}

	if array {
		_, err = output.Write([]byte{'['})
	}

	for i, arg := range args {
		err = tmpl.Execute(output, arg)
		if err != nil {
			return
		}

		if array && i < len(args)-1 {
			output.Write([]byte{','})
		}

		_, err = output.Write([]byte{'\n'})
		if err != nil {
			return
		}
	}

	if array {
		_, err = output.Write([]byte{']'})
	}

	return nil
}
