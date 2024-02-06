//go:build !solution

package externalsort

import (
	"container/heap"
	"io"
	"os"
	"sort"
)

type strHeap []string

func (sh strHeap) Len() int           { return len(sh) }
func (sh strHeap) Less(i, j int) bool { return sh[i] < sh[j] }
func (sh strHeap) Swap(i, j int)      { sh[i], sh[j] = sh[j], sh[i] }

func (sh *strHeap) Push(x interface{}) {
	*sh = append(*sh, x.(string))
}

func (sh *strHeap) Pop() interface{} {
	old := *sh
	n := len(old)
	x := old[n-1]
	*sh = old[0 : n-1]
	return x
}

func Merge(w LineWriter, readers ...LineReader) error {
	sh := &strHeap{}
	heap.Init(sh)
	check := make(map[int]struct{})
	for len(check) != len(readers) {
		for i, r := range readers {
			str, err := r.ReadLine()
			if err == io.EOF && str == "" {
				check[i] = struct{}{}
			} else if err == io.EOF && str != "" || err == nil {
				heap.Push(sh, str)
			} else {
				return err
			}
		}
	}

	for sh.Len() > 0 {
		str := heap.Pop(sh)
		err := w.Write(str.(string))
		if err != nil {
			return err
		}
	}
	return nil
}

func Sort(w io.Writer, in ...string) error {
	for _, name := range in {
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		var lines []string
		sr := NewReader(file)
		for {
			line, err := sr.ReadLine()
			if err == io.EOF && line == "" {
				break
			} else if err != nil && err != io.EOF {
				return err
			}
			lines = append(lines, line)
		}
		sort.Strings(lines)
		file.Close()

		file, err = os.OpenFile(name, os.O_WRONLY, 0600)
		if err != nil {
			return err
		}

		sw := NewWriter(file)
		for _, line := range lines {
			err = sw.Write(line)
			if err != nil {
				return err
			}
		}
	}
	var readers []LineReader
	for _, name := range in {
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		sr := NewReader(file)
		readers = append(readers, sr)
	}
	sw := NewWriter(w)
	err := Merge(sw, readers...)
	if err != nil {
		return err
	}
	return nil
}
