package helpers

type Artifact struct {
	Folder         string
	Output string
	Worspace string
}

type PaternData struct{
	File string
	Line int
	InstructionString string
	Instruction DeclaredInstruction
	ErrorMessage string
}

type DeclaredInstruction struct{
	Resource string
	Condition string
	TargetType string
	Value string
}