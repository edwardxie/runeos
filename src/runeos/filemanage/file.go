package filemanage

const (
	FileType Type = 1 << iota
	FolerType
)

type (
	File        interface{}
	FileInfo    interface{}
	FileManager interface{}
	Attr        interface{}
	ReadFiler   interface{}
	WriteFiler  interface{}
	Folder      interface{}

	Type int

	Osio struct {
		root []*Mount
	}

	Mount struct {
		Name    string
		Type    int
		Path    string
		Mappath string
	}
)

func New() FileManager {
	return nil
}

func AddMount() {

}

func DelMount() {

}
