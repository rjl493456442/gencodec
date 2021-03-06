// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package omitempty

import (
	"encoding/json"
)

func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Int replacedInt `json:",omitempty"`
	}
	var enc X
	enc.Int = replacedInt(x.Int)
	return json.Marshal(&enc)
}

func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Int *replacedInt `json:",omitempty"`
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = int(*dec.Int)
	}
	return nil
}

func (x X) MarshalYAML() (interface{}, error) {
	type X struct {
		Int replacedInt `json:",omitempty"`
	}
	var enc X
	enc.Int = replacedInt(x.Int)
	return &enc, nil
}

func (x *X) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type X struct {
		Int *replacedInt `json:",omitempty"`
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = int(*dec.Int)
	}
	return nil
}

func (x X) MarshalTOML() (interface{}, error) {
	type X struct {
		Int replacedInt `json:",omitempty"`
	}
	var enc X
	enc.Int = replacedInt(x.Int)
	return &enc, nil
}

func (x *X) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type X struct {
		Int *replacedInt `json:",omitempty"`
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = int(*dec.Int)
	}
	return nil
}
