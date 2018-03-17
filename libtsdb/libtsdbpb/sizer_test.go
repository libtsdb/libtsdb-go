package libtsdbpb

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
)

func TestTag_RawSize(t *testing.T) {
	assert := asst.New(t)

	tag := Tag{K: "ascii", V: "value"}
	assert.Equal(80, tag.RawSize())
	// TODO: I guess the sizer in gogo protobuf means byte?
	assert.Equal(14, tag.Size())
}

func TestPointInt_RawSize(t *testing.T) {
	assert := asst.New(t)

	p := PointInt{T: 10086, V: 123}
	assert.Equal(128, p.RawSize())
}

func TestPointIntTagged_RawSize(t *testing.T) {
	assert := asst.New(t)

	pt := PointIntTagged{
		Point: PointInt{T: 10086, V: 123},
		Tags:  []Tag{{K: "ascii", V: "value"}},
	}
	assert.Equal(208, pt.RawSize())
}

func TestEmptySeries_RawSize(t *testing.T) {
	assert := asst.New(t)

	s := EmptySeries{
		Name: "cpu",
		Tags: []Tag{
			{K: "region", V: "en-us"},
		},
	}
	assert.Equal(112, s.RawSize())
}

func TestSeriesIntTagged_RawSize(t *testing.T) {
	assert := asst.New(t)

	sit := SeriesIntTagged{
		Name: "cpu",
		Tags: []Tag{
			{K: "region", V: "en-us"},
		},
		Points: []PointInt{
			{T: 10086, V: 100},
		},
	}
	assert.Equal(240, sit.RawSize())
}
