//go:build !solution

package hogwarts

var stack []string
var colours map[string]int

func TopologicalSort(node string, prereqs map[string][]string) {
	if colours[node] == 2 {
		return
	}
	if colours[node] == 1 {
		panic("Зацикливание")
	}
	colours[node] = 1
	if _, ok := prereqs[node]; !ok {
		colours[node] = 2
		stack = append(stack, node)
	} else {
		for _, value := range prereqs[node] {
			TopologicalSort(value, prereqs)
		}
		colours[node] = 2
		stack = append(stack, node)
	}
}

func GetCourseList(prereqs map[string][]string) []string {
	colours = make(map[string]int)
	for key := range prereqs {
		colours[key] = 0
	}

	for key := range prereqs {
		TopologicalSort(key, prereqs)
	}

	return stack
}

// func main() {
// 	var naiveScience = map[string][]string{
// 		"здравый смысл":    {},
// 		"русский язык":     {"здравый смысл"},
// 		"литература":       {"здравый смысл"},
// 		"иностранный язык": {"здравый смысл"},
// 		"алгебра":          {"здравый смысл"},
// 		"геометрия":        {"здравый смысл"},
// 		"информатика":      {"здравый смысл"},
// 		"история":          {"здравый смысл"},
// 		"обществознание":   {"здравый смысл"},
// 		"география":        {"здравый смысл"},
// 		"биология":         {"здравый смысл"},
// 		"физика":           {"здравый смысл"},
// 		"химия":            {"здравый смысл"},
// 		"музыка":           {"здравый смысл"},
// 	}
// 	GetCourseList(naiveScience)
// }
