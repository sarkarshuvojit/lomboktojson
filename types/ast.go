package types

// NodeType identifies the kind of AST node.
type NodeType string

const (
	NodeObject NodeType = "OBJECT"
	NodeArray  NodeType = "ARRAY"
	NodeValue  NodeType = "VALUE"
)

// Node is implemented by all AST nodes.
type Node interface {
	NodeType() NodeType
}

// ObjectField represents a single key/value pair inside an object.
type ObjectField struct {
	Key   string
	Value Node
}

// ObjectNode models a Lombok object instance.
type ObjectNode struct {
	ClassName string
	Fields    []ObjectField
}

func (ObjectNode) NodeType() NodeType { return NodeObject }

// ArrayNode models an array of values or nested objects.
type ArrayNode struct {
	Elements []Node
}

func (ArrayNode) NodeType() NodeType { return NodeArray }

// ValueNode stores a primitive or string literal.
type ValueNode struct {
	Value string
}

func (ValueNode) NodeType() NodeType { return NodeValue }
