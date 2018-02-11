package first

import "testing"

type structWithStringer struct{ result string }

func (s structWithStringer) String() string { return s.result }

func TestString(t *testing.T) {
	var nilStringPointer *string
	happyString := "happy"
	emptyString := ""

	testCases := []struct {
		Name     string
		Input    []interface{}
		Expected string
	}{
		{"noargs", []interface{}{}, ""},
		{"basic", []interface{}{"hello", "x"}, "hello"},
		{"emptyfirst", []interface{}{"", "foo"}, "foo"},
		{"allempty", []interface{}{"", "", ""}, ""},

		{"nil", []interface{}{nil}, ""},
		{"nilfirst", []interface{}{nil, "x"}, "x"},

		{"pointer", []interface{}{&happyString, "x"}, "happy"},
		{"nilstringpointer", []interface{}{nilStringPointer, "hello"}, "hello"},
		{"emptypointer", []interface{}{&emptyString, "world"}, "world"},

		{"stringfunc", []interface{}{func() string { return "hi" }, "x"}, "hi"},
		{"emptystringfunc", []interface{}{func() string { return "" }, "foo"}, "foo"},

		{"stringer", []interface{}{structWithStringer{"good"}, "x"}, "good"},
		{"emptystringer", []interface{}{structWithStringer{""}, "good"}, "good"},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			value := String(tc.Input...)
			if value != tc.Expected {
				t.Errorf("expected %s but got %s", tc.Expected, value)
			}
		})
	}
}
