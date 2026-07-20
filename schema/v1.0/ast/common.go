package ast

type RenameAttributes struct {
	AttributeMap AttributeMap `yaml:"attribute_map"`
}

type AttributeMap map[string]string
