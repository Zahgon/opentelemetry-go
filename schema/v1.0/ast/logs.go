package ast

type Logs struct {
	Changes []LogsChange
}

type LogsChange struct {
	RenameAttributes *RenameAttributes `yaml:"rename_attributes"`
}
