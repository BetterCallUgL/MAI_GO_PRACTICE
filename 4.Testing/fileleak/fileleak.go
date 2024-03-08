package fileleak

import (
	"io/fs"
	"os"
)

type testingT interface {
	Errorf(msg string, args ...interface{})
	Cleanup(func())
}

func FillMap(dict map[string]int, dir []fs.DirEntry) {
	for _, file := range dir {
		desc, _ := os.Readlink("/proc/self/fd/" + file.Name())
		if _, ok := dict[desc]; !ok {
			dict[desc] = 0
		}
		dict[desc]++
	}
}

func VerifyNone(t testingT) {
	path := "/proc/self/fd"

	dir, _ := os.ReadDir(path)
	old := make(map[string]int)
	FillMap(old, dir)

	t.Cleanup(func() {
		dir, _ = os.ReadDir(path)
		new := make(map[string]int)
		FillMap(new, dir)

		for key, newVal := range new {
			if oldVal, ok := old[key]; ok {
				if newVal > oldVal {
					t.Errorf("Detected unclosed file clone: %v", key)
				}
			} else {
				t.Errorf("Detected unclosed file:%v", key)
			}
		}
	})
}
