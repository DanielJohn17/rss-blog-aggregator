package handlers

type Command struct {
	Name string
	Args []string
}

type Commands map[string]func(*State, Command) error

func (c *Commands) Run(s *State, cmd Command) error {
	if handeler, ok := (*c)[cmd.Name]; ok {
		return handeler(s, cmd)
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if (*c) == nil {
		(*c) = make(Commands)
	}

	(*c)[name] = f
}
