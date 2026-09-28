package manifest

type Metadata = map[string]any

type Manifest struct {
	Type     Type
	Name     Name
	Metadata Metadata
	Spec     any
}

func (m Manifest) Identity() Identity {
	return Identity{Kind: m.Type.Kind, Name: m.Name}
}

type Catalog []Manifest
