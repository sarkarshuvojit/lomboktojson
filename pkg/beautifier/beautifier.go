package beautifier

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sarkarshuvojit/lomboktojson/pkg/parser"
	"github.com/sarkarshuvojit/lomboktojson/pkg/scanner"
	"github.com/sarkarshuvojit/lomboktojson/types"
)

// Beautify prints the AST back into readable Lombok toString style output.
func Beautify(node types.Node, indent int) (asBytes []byte, err error) {
	if indent <= 0 {
		indent = 1
	}

	if node == nil {
		return []byte{}, nil
	}

	var buf bytes.Buffer
	if err := writeNode(&buf, node, indent, 0, false); err != nil {
		return nil, err
	}

	return []byte(strings.TrimRight(buf.String(), "\n")), nil
}

func writeNode(buf *bytes.Buffer, node types.Node, indent int, depth int, inline bool) error {
	switch n := node.(type) {
	case *types.ObjectNode:
		return writeObject(buf, n, indent, depth, inline)
	case *types.ArrayNode:
		return writeArray(buf, n, indent, depth, inline)
	case *types.ValueNode:
		buf.WriteString(n.Value)
		return nil
	default:
		return fmt.Errorf("unknown node type %T", node)
	}
}

func writeObject(buf *bytes.Buffer, node *types.ObjectNode, indent int, depth int, inline bool) error {
	if !inline {
		writeIndent(buf, indent, depth)
	}
	buf.WriteString(node.ClassName)
	buf.WriteString("(")

	if len(node.Fields) == 0 {
		buf.WriteString(")")
		return nil
	}

	buf.WriteString("\n")
	for i, field := range node.Fields {
		writeIndent(buf, indent, depth+1)
		buf.WriteString(field.Key)
		buf.WriteString("=")
		if err := writeNode(buf, field.Value, indent, depth+1, true); err != nil {
			return err
		}
		if i < len(node.Fields)-1 {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	writeIndent(buf, indent, depth)
	buf.WriteString(")")
	return nil
}

func writeArray(buf *bytes.Buffer, node *types.ArrayNode, indent int, depth int, inline bool) error {
	if !inline {
		writeIndent(buf, indent, depth)
	}
	buf.WriteString("[")
	if len(node.Elements) == 0 {
		buf.WriteString("]")
		return nil
	}

	buf.WriteString("\n")
	for i, elem := range node.Elements {
		writeIndent(buf, indent, depth+1)
		if err := writeNode(buf, elem, indent, depth+1, true); err != nil {
			return err
		}
		if i < len(node.Elements)-1 {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	writeIndent(buf, indent, depth)
	buf.WriteString("]")
	return nil
}

func writeIndent(buf *bytes.Buffer, indent int, depth int) {
	buf.WriteString(strings.Repeat(" ", indent*depth))
}

// BeautifySource scans the input string and returns the formatted output.
func BeautifySource(source string, indent int) (string, error) {
	sc := scanner.NewScanner(strings.NewReader(source))
	tokens, err := sc.Scan()
	if err != nil {
		return "", err
	}
	node, err := parser.Parse(tokens)
	if err != nil {
		return "", err
	}
	formatted, err := Beautify(node, indent)
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}
