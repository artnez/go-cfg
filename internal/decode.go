package internal

import "github.com/mitchellh/mapstructure"

func Decode(data interface{}, structPtr interface{}) {
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           structPtr,
		ZeroFields:       true,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err == nil {
		decoder.Decode(data)
	}
}
