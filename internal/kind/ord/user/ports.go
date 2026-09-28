package user

type SchemaValidator interface {
	Validate(schema map[string]any, document map[string]any) error
}
