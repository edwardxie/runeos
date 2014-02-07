package plugins

type Plugin struct {
	Name  string
	Type  string
	Mode  int
	Entry func()
}
