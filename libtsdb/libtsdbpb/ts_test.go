package libtsdbpb

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestEmptySeries_Size(t *testing.T) {
	assert := asst.New(t)

	s := EmptySeries{
		Name: "cpu",
		Tags: []Tag{
			{K: "region", V: "en-us"},
		},
	}
	assert.Equal(22, s.Size())
}
