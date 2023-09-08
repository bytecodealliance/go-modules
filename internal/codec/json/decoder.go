package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ydnar/wit-bindgen-go/internal/codec"
)

type Decoder struct {
	dec    *json.Decoder
	codecs codec.Codecs
}

func NewDecoder(r io.Reader, codecs ...codec.Codec) *Decoder {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	return &Decoder{
		dec:    dec,
		codecs: codec.Codecs(codecs),
	}
}

func (dec *Decoder) Decode(v any) error {
	c, err := dec.codecs.Codec(v)
	if err != nil {
		return err
	}
	if c != nil {
		v = c
	}

	tok, err := dec.dec.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	err = dec.decodeToken(v, tok)
	if err != nil {
		return err
	}

	if end, ok := v.(codec.EndDecoder); ok {
		err = end.DecodeEnd()
	}

	return err
}

func (dec *Decoder) decodeToken(v any, tok json.Token) error {
	if tok == nil {
		return codec.DecodeNil(v)
	}

	switch tok := tok.(type) {
	case bool:
		return codec.DecodeBool(v, tok)
	case json.Number:
		return codec.DecodeNumber(v, string(tok))
	case string:
		return codec.DecodeString(v, tok)
	case json.Delim:
		switch tok {
		case '{':
			return dec.decodeObject(v)
		case '[':
			return dec.decodeArray(v)
		default:
			return fmt.Errorf("unexpected JSON token %v at offset %d", tok, dec.dec.InputOffset())
		}
	}

	return nil
}

// decodeObject decodes a JSON object into v.
// It expects that the initial { token has already been decodec.
func (dec *Decoder) decodeObject(o any) error {
	d, ok := o.(codec.FieldDecoder)
	if !ok {
		// TODO: how to handle undecodable objects?
		d = &ignore{}
	}

	for dec.dec.More() {
		name, err := dec.stringToken()
		if err != nil {
			return err
		}
		v, err := d.DecodeField(name)
		if err != nil {
			return err
		}
		err = dec.Decode(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// decodeArray decodes a JSON array into v.
// It expects that the initial [ token has already been decoded.
func (dec *Decoder) decodeArray(v any) error {
	d, ok := v.(codec.ElementDecoder)
	if !ok {
		// TODO: how to handle undecodable arrays?
		d = &ignore{}
	}

	for i := 0; dec.dec.More(); i++ {
		v, err := d.DecodeElement(i)
		if err != nil {
			return err
		}
		err = dec.Decode(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dec *Decoder) stringToken() (string, error) {
	tok, err := dec.dec.Token()
	if err != nil {
		return "", err
	}
	s, ok := tok.(string)
	if !ok {
		return "", fmt.Errorf("unexpected JSON token %v at offset %d", tok, dec.dec.InputOffset())
	}
	return s, nil
}

type ignore struct{}

func (i *ignore) DecodeField(string) (any, error) {
	return i, nil
}

func (i *ignore) DecodeElement(int) (any, error) {
	return i, nil
}

func (i *ignore) UnmarshalJSON([]byte) error {
	return nil
}
