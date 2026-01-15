package generator

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

func isNumeric(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

func isFloatingPoint(s string) bool {
	re := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	return re.MatchString(s)
}

func getOptionallyQuotedValue(val string) string {
	isNull := val == "null"
	isBool := val == "true" || val == "false"
	isNum := isNumeric(val)
	isFloat := isFloatingPoint(val)

	if isNull || isBool || isNum || isFloat {
		return fmt.Sprintf("%s", val)
	}
	return fmt.Sprintf("\"%s\"", val)
}

// Generate converts an AST produced by the parser into JSON bytes.
func Generate(node types.Node) ([]byte, error) {
	if node == nil {
		return []byte("{}"), nil
	}

	var buf bytes.Buffer
	if err := writeNode(&buf, node); err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		buf.WriteString("{}")
	}
	return buf.Bytes(), nil
}

func writeNode(buf *bytes.Buffer, node types.Node) error {
	switch n := node.(type) {
	case *types.ObjectNode:
		return writeObject(buf, n)
	case *types.ArrayNode:
		return writeArray(buf, n)
	case *types.ValueNode:
		buf.WriteString(getOptionallyQuotedValue(n.Value))
		return nil
	default:
		return fmt.Errorf("unknown node type %T", node)
	}
}

func writeObject(buf *bytes.Buffer, node *types.ObjectNode) error {
	buf.WriteString("{")
	for i, field := range node.Fields {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", field.Key))
		if err := writeNode(buf, field.Value); err != nil {
			return err
		}
	}
	buf.WriteString("}")
	return nil
}

func writeArray(buf *bytes.Buffer, node *types.ArrayNode) error {
	buf.WriteString("[")
	for i, elem := range node.Elements {
		if i > 0 {
			buf.WriteString(",")
		}
		if err := writeNode(buf, elem); err != nil {
			return err
		}
	}
	buf.WriteString("]")
	return nil
}
