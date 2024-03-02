package coder

import "io"

type GammingCoder struct {
	key []byte
}

func New(key []byte) *GammingCoder {
	return &GammingCoder{
		key: key,
	}
}

func (c *GammingCoder) Encode(r io.Reader, w io.Writer) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for i := 0; i < len(input); i++ {
		input[i] ^= c.key[i%len(c.key)]
	}

	_, err = w.Write(input)

	return err
}
