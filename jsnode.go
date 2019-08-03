package jstruct

import (
	"fmt"
	"sort"
)

const (
	// RootObject root
	RootObject = iota
	// Object type
	Object
	// Array type
	Array
	// Primitive type
	Primitive
)

type jsNode struct {
	children []*jsNode
	typ      int
	name     string
	val      interface{}
}

func (n *jsNode) addChild(c *jsNode) *jsNode {
	n.children = append(n.children, c)

	return n
}

func (n *jsNode) isEmpty() bool {
	return len(n.children) == 0
}

func (n *jsNode) repr() string {
	switch n.typ {
	case Object, RootObject:
		return reprObject(n)
	case Array:
		return reprArray(n)
	case Primitive:
		return reprPrimitive(n)
	default:
		panic(fmt.Sprintf("Wrong node type %#v", n.typ))
	}
}

func reprPrimitive(n *jsNode) string {
	switch n.val.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case map[string]interface{}:
		return "map[string]interface{}"
	default:
		return fmt.Sprintf("%#v %s", n.val, n.name)
	}
}

func reprObject(n *jsNode) string {
	if n.name == "" {
		n.name = "object"
	}

	expr := fmt.Sprintf("\ntype %s struct {\n", n.name)

	var keys []string
	for _, c := range n.children {
		if c.name != "" {
			keys = append(keys, c.name)
		}
	}

	sort.Strings(keys)

	for _, k := range keys {
		for i, c := range n.children {
			if c.name == k {
				expr += fmt.Sprintf("\t%s\t%s\n", k, typeAsString(n.children[i]))
			}
		}

	}

	return expr + "}\n"
}

func reprArray(n *jsNode) string {
	var typ string
	for _, v := range n.children {
		if typ == "" {
			typ = typeAsString(v)
		}

		if typ != typeAsString(v) {
			return "[]interface{}"
		}
	}

	if typ == "" {
		return "[]interface{}"
	}

	return fmt.Sprintf("[]%s", typ)
}

func reprMap(n *jsNode) string {
	var repr string

	for _, c := range n.children {
		repr += fmt.Sprintf("%s", typeAsString(c))
	}

	return repr
}

func typeAsString(n *jsNode) string {
	switch n.typ {
	case Primitive:
		return reprPrimitive(n)
	case Array:
		return reprArray(n)
	case Object, RootObject:
		return fmt.Sprintf("%s", n.name)
	default:
		return "interface{}"
	}
}
