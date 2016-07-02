// Package jsons converts JSON arrays into streams of objects.
package jsons

import "io"

// The Reader is an io.Reader that wraps another reader for the purposes
// of modifying a JSON stream. When the underlying reader provides a JSON
// array object ([obj, obj, ...]) this will be converted into a stream of
// newline terminated objects (obj\n obj\n obj\n) suitable for consumption
// by a bufio.Scanner or json.Decoder.
type Reader struct {
	level    int
	inString bool
	escaped  bool
	r        io.Reader
}

func New(r io.Reader) *Reader {
	return &Reader{
		r: r,
	}
}

func (s *Reader) Read(bs []byte) (int, error) {
	n, err := s.r.Read(bs)
	if n == 0 {
		return n, err
	}
	for i, c := range bs[:n] {
		if s.inString {
			if s.escaped {
				s.escaped = false
			} else if c == '"' {
				s.inString = false
			} else if c == '\\' {
				s.escaped = true
			}
			continue
		}

		switch c {
		case '"':
			s.inString = true

		case '[':
			if s.level == 0 {
				bs[i] = ' '
			}
			s.level++
		case ']':
			if s.level == 1 {
				bs[i] = '\n'
			}
			s.level--
		case ',':
			if s.level == 1 {
				bs[i] = '\n'
			}

		case '{':
			s.level++
		case '}':
			s.level--
		}
	}
	return n, err
}
