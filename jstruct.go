package jstruct

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/tidwall/gjson"
)

func iterate(st *stack, root *jsNode, data interface{}) *jsNode {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Slice {
		root.typ = Array

		for i := 0; i < v.Len(); i++ {
			n := &jsNode{
				val: v.Index(i).Interface(),
			}
			root.addChild(iterate(st, n, v.Index(i).Interface()))
		}

		return root
	}

	if v.Kind() == reflect.Map {
		root.typ = Object

		for _, k := range v.MapKeys() {
			n := &jsNode{
				name:     k.String(),
				val:      v.MapIndex(k).Interface(),
				children: make([]*jsNode, 0),
			}
			root.addChild(iterate(st, n, v.MapIndex(k).Interface()))
		}

		st.push(root)

		return root
	}

	if v.Kind() != reflect.Invalid {
		root.typ = Primitive
		root.val = v.Interface()
	}

	return root
}

func createMainStruct(packageName, structName string, st *stack, root *jsNode) string {
	pDef := fmt.Sprintf("package %s\n\n", packageName)

	N := st.len()
	for i := 0; i < N; i++ {
		if n, ok := st.pop(); ok {
			pDef += n.repr()
		}
	}

	return pDef + "\n"
}

// ReadFromFile extracts json file content as string
func ReadFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}

// WriteToFile write parsed output to a GO file
func WriteToFile(filename, content string) error {
	d := []byte(content)
	err := ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Parse json string
func Parse(packageName, rootName, jsn string) string {
	m := gjson.Parse(jsn).Value()

	s := new(stack)

	root := &jsNode{
		name:     rootName,
		children: make([]*jsNode, 0),
	}

	iterate(s, root, m)

	return createMainStruct(packageName, rootName, s, root)
}

// Convert json file to go struct definition
func Convert(inputFile, outputFile, packageName, name string) error {
	str, err := ReadFromFile(inputFile)
	if err != nil {
		return err
	}

	parsedResult := Parse(packageName, name, str)

	err = WriteToFile(outputFile, parsedResult)
	if err != nil {
		return err
	}

	return nil
}
