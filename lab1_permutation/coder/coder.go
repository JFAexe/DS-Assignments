package coder

import (
	"bytes"
	"io"
	"slices"
	"sort"
)

type PermutationCoder struct {
	key   []byte
	order []int
}

func New(key []byte) *PermutationCoder {
	coder := &PermutationCoder{
		key: key,
	}

	coder.calculateOrder()

	return coder
}

func (c *PermutationCoder) Encode(r io.Reader, w io.Writer) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	columns := c.setupColumns(input)

	var col, row int

	for i := 0; i < len(input); i++ {
		columns[col][row] = input[i]

		if col = (col + 1) % len(c.key); col == 0 {
			row++
		}
	}

	for _, i := range c.order {
		if _, err = w.Write(columns[i]); err != nil {
			return err
		}
	}

	return nil
}

func (c *PermutationCoder) Decode(r io.Reader, w io.Writer) error {
	input, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	columns := c.setupColumns(input)

	var (
		buf bytes.Buffer

		shift, col, row int
	)

	for _, cur := range c.order {
		columns[cur] = input[shift : shift+len(columns[cur])]

		shift += len(columns[cur])
	}

	for range input {
		if _, col = buf.WriteByte(columns[col][row]), (col+1)%len(c.key); col == 0 {
			row++
		}
	}

	_, err = buf.WriteTo(w)

	return err
}

func (c *PermutationCoder) calculateOrder() {
	bs := make(map[byte]int)

	for _, b := range c.key {
		bs[b] = 0
	}

	ordered := make([]byte, len(c.key))

	copy(ordered, c.key)

	sort.Slice(ordered, func(i, j int) bool { return ordered[i] < ordered[j] })

	for _, b := range ordered {
		var count int

		c.order = append(c.order, slices.IndexFunc(c.key, func(cb byte) bool {
			if b != cb {
				return false
			}

			if count++; count <= bs[b] {
				return false
			}

			bs[b]++

			return true
		}))
	}
}

func (c *PermutationCoder) setupColumns(input []byte) [][]byte {
	var (
		columns [][]byte

		last = len(input) % len(c.key)
		rows = (len(input) - last) / len(c.key)
	)

	for i := 0; i < len(c.key); i++ {
		size := rows

		if i < last {
			size += 1
		}

		columns = append(columns, make([]byte, size))
	}

	return columns
}
