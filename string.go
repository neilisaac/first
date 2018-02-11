package first

import "fmt"

// String returns the first non-nil, non-empty argument.
// It supports strings, string pointers, functions returning strings, and stringers.
func String(stringables ...interface{}) string {
	for _, s := range stringables {
		if s == nil {
			continue
		}

		value := ""
		switch s.(type) {
		case string:
			value = s.(string)
		case *string:
			if s.(*string) != nil {
				value = *s.(*string)
			}
		case func() string:
			value = s.(func() string)()
		case fmt.Stringer:
			value = s.(fmt.Stringer).String()
		default:
			panic("parameter with unsupported type")
		}

		if value != "" {
			return value
		}
	}
	return ""
}
