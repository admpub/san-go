package san

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type test1 struct {
	A string
	B int64 `san:"c"`
}

func TestSimpleStruct(t *testing.T) {
	str := `
a = "test"
c = 1
`
	var s test1

	err := Unmarshal([]byte(str), &s)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "test", s.A)
	assert.Equal(t, int64(1), s.B)

}
