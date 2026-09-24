package manifest

import "fmt"

type Metadata = map[string]any
type Spec = map[string]any

type Type struct {
	Kind       Kind
	APIVersion APIVersion
}

type Identity struct {
	Kind Kind
	Name Name
}

func (i Identity) String() string {
	return fmt.Sprintf("%s %q", i.Kind, i.Name)
}

type Manifest struct {
	Type     Type
	Name     Name
	Metadata Metadata
	Spec     any
}

func (m Manifest) Identity() Identity {
	return Identity{Kind: m.Type.Kind, Name: m.Name}
}
