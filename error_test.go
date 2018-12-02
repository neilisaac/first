package first

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	var nilError error
	var someError = errors.New("real error")

	if err := Error(); err != nil {
		t.Error("expected nil when passing no arguments: ", err)
	}

	if err := Error(nil, nilError, someError, nilError, nil); err != someError {
		t.Error("incorrect error returned: ", err)
	}

	if err := Error(nil, nilError, nil, nilError); err != nil {
		t.Error("expected nil: ", err)
	}
}
