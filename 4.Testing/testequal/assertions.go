package testequal

import "fmt"

func Message(t T, msgAndArgs ...interface{}) string {
	var message string
	if len(msgAndArgs) == 0 {
		message = ""
	} else if len(msgAndArgs) == 1 {
		message = msgAndArgs[0].(string)
	} else {
		message = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}

	return message
}

func Equal(expected, actual interface{}) bool {
	equal := true

	switch val1 := expected.(type) {
	case uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int:
		if actual != expected {
			equal = false
		}
	case string:
		if val2, ok := actual.(string); ok {
			if val1 != val2 {
				equal = false
			}
		} else {
			equal = false
		}
	case map[string]string:
		if val2, ok := actual.(map[string]string); ok {
			if len(val1) != len(val2) || val1 == nil && val2 != nil || val2 == nil && val1 != nil {
				equal = false
			}
			for key := range val1 {
				if val, ok := val2[key]; ok {
					if val != val1[key] {
						equal = false
						break
					}
				} else {
					equal = false
					break
				}
			}
		} else {
			equal = false
		}

	case []int:
		if val2, ok := actual.([]int); ok {
			if len(val1) != len(val2) || val1 == nil && val2 != nil || val2 == nil && val1 != nil {
				equal = false
			} else {
				for i := range val1 {
					if val1[i] != val2[i] {
						equal = false
						break
					}
				}
			}
		} else {
			equal = false
		}
	case []byte:
		if val2, ok := actual.([]byte); ok {
			if len(val1) != len(val2) || val1 == nil && val2 != nil || val2 == nil && val1 != nil {
				equal = false
			} else {
				for i := range val1 {
					if val1[i] != val2[i] {
						equal = false
						break
					}
				}
			}
		} else {
			equal = false
		}
	default:
		equal = false
	}

	return equal
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	message := Message(t, msgAndArgs...)
	equal := Equal(expected, actual)
	t.Helper()
	if !equal {
		t.Errorf(`not equal :
		expected: %v
		actual: %v
		message: %s 
		`, expected, actual, message)
	}

	return equal
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	message := Message(t, msgAndArgs...)
	equal := Equal(expected, actual)
	t.Helper()
	if equal {
		t.Errorf(`equal:
		expected: %v
		actual: %v
		message: %s 
		`, expected, actual, message)
	}

	return !equal
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	message := Message(t, msgAndArgs...)
	equal := Equal(expected, actual)
	t.Helper()
	if !equal {
		t.Errorf(`not equal:
		expected: %v
		actual: %v
		message: %s 
		`, expected, actual, message)
		t.FailNow()
	}
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	message := Message(t, msgAndArgs...)
	equal := Equal(expected, actual)
	t.Helper()
	if equal {
		t.Errorf(`equal:
		expected: %v
		actual: %v
		message: %s 
		`, expected, actual, message)
		t.FailNow()
	}
}
