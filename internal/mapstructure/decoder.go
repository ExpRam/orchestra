package mapstructure

import (
	mapstructurev2 "github.com/go-viper/mapstructure/v2"
)

type MapstructureSpecDecoder struct {
	tag string
}

func NewMapstructureSpecDecoder(tag string) MapstructureSpecDecoder {
	return MapstructureSpecDecoder{tag: tag}
}

func (d MapstructureSpecDecoder) Decode(tree map[string]any, target any) error {
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
