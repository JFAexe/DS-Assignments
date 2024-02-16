package coder

import "io"

type PolyalphabeticCoder struct {
	key []byte
}

func New(key string) *PolyalphabeticCoder {
	return &PolyalphabeticCoder{
		key: []byte(key),
	}
}

func (c *PolyalphabeticCoder) Encode(r io.Reader, w io.Writer) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for i := range input {
		input[i] += c.key[i%len(c.key)]
	}

	_, err = w.Write(input)

	return err
}

func (c *PolyalphabeticCoder) Decode(r io.Reader, w io.Writer) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for i := range input {
		input[i] -= c.key[i%len(c.key)]
	}

	_, err = w.Write(input)

	return err
}
