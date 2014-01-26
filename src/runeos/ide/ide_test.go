package ide

import (
	"testing"
)

func TestOverridden(t *testing.T) {
	// ws := &Workspace{}
	ws.FontFace = "Droid Sans Mono"
	ws.FontSize = 11
	ws.Theme = "Soda Drak 3 Theme"
	// t.Logf("Workspace: %#v\nRun path: %v\n", ws, runpath)

}

func TestTojson(t *testing.T) {
	InitIDE()
	if err := toJson(); err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestToxml(t *testing.T) {
	InitIDE()
	if err := toXml(); err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestFromjson(t *testing.T) {
	ws = &Workspace{}
	if err := fromJson(); err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	t.Logf("Read json file ok. -> \n%#v", ws)
}

func TestFromxml(t *testing.T) {
	ws = &Workspace{}
	if err := fromXml(); err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	t.Logf("Read XML file ok. -> \n%#v", ws)
}
