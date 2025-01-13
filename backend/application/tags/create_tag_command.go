package tags

type CreateTagCommand struct {
	Name string
}

func NewCreateTagCommand(name string) *CreateTagCommand {
	return &CreateTagCommand{name}
}
