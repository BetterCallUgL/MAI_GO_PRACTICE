//go:build !change

package externalsort

import (
	"bufio"
	"io"
)

type SuperReader struct {
	reader *bufio.Reader
}

type SuperWriter struct {
	writer *bufio.Writer
}

type LineReader interface {
	ReadLine() (string, error)
}

type LineWriter interface {
	Write(l string) error
}

func NewReader(r io.Reader) LineReader {
	return &SuperReader{bufio.NewReader(r)}
}

func NewWriter(w io.Writer) LineWriter {
	return &SuperWriter{bufio.NewWriter(w)}
}

func (sr *SuperReader) ReadLine() (string, error) {
	s, err := sr.reader.ReadString('\n')
	if err == nil {
		s = s[:len(s)-1]
	}

	return s, err
}

func (sw *SuperWriter) Write(l string) error {
	s := l + "\n"
	_, err := sw.writer.WriteString(s)
	return err
}
