package otp

import (
	"io"
)

type SuperReader struct {
	reader  io.Reader
	cichper io.Reader
}

type SuperWriter struct {
	writer io.Writer
	cipher io.Reader
}

func (sr *SuperReader) Read(p []byte) (int, error) {
	n, err := sr.reader.Read(p)
	if err != nil && err != io.EOF {
		return 0, err
	}
	key := make([]byte, n)
	sr.cichper.Read(key)
	for i := 0; i < n; i++ {
		p[i] ^= key[i]
	}
	return n, err
}

func NewReader(r io.Reader, prng io.Reader) io.Reader {
	sr := new(SuperReader)
	sr.reader = r
	sr.cichper = prng
	return sr
}

func (sw *SuperWriter) Write(p []byte) (int, error) {
	key := make([]byte, len(p))
	sw.cipher.Read(key)
	new := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		new[i] = p[i] ^ key[i]
	}
	n, err := sw.writer.Write(new)
	return n, err
}

func NewWriter(w io.Writer, prng io.Reader) io.Writer {
	sw := new(SuperWriter)
	sw.writer = w
	sw.cipher = prng
	return sw
}
