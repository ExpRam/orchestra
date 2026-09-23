package manifest

type Manifest struct {
	Metadata Metadata
	Spec     Spec
}

type Metadata = map[string]any
type Spec = map[string]any
