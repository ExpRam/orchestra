package mapstructure

import (
	mapstructurev2 "github.com/go-viper/mapstructure/v2"

	"github.com/expram/orchestra/internal/operation/resolver"
)

type SpecDecoder struct {
	tag string
}

var _ resolver.SpecDecoder = SpecDecoder{}

func NewSpecDecoder(tag string) SpecDecoder {
	return SpecDecoder{tag: tag}
}

func (d SpecDecoder) Decode(tree map[string]any, target any) error {
	decoder, err := mapstructurev2.NewDecoder(&mapstructurev2.DecoderConfig{
		TagName:     d.tag,
		Result:      target,
		ErrorUnused: true,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(tree); err != nil {
		return translate(err)
	}

	return nil
}
