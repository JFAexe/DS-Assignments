package coder

import "io"

type GammingCoder struct {
	key []byte
}

func New(key string) *GammingCoder {
	return &GammingCoder{
		key: []byte(key),
	}
}

func (c *GammingCoder) Encode(r io.Reader, w io.Writer) error {
	msg, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for i := 0; i < len(msg); i++ {
		msg[i] ^= c.key[i%len(c.key)]
	}

	_, err = w.Write(msg)

	return err
}
