package dto

type UnresolvedManifest struct {
	Kind       string
	APIVersion string
	Name       string
	Metadata   map[string]any
	Spec       map[string]any
}
