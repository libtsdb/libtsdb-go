package libtsdbpb

// Sizer is used to output raw size of series in byte
type Sizer interface {
	// RawSize counts series identifier and points without marshal
	RawSize() int
	// RawMetaSize only counts series identifier, name + tags
	RawMetaSize() int
}

var _ Sizer = (*Tag)(nil)
var _ Sizer = (*PointInt)(nil)
var _ Sizer = (*PointDouble)(nil)
var _ Sizer = (*PointIntTagged)(nil)
var _ Sizer = (*PointDoubleTagged)(nil)
var _ Sizer = (*SeriesIntTagged)(nil)
var _ Sizer = (*SeriesDoubleTagged)(nil)
var _ Sizer = (*SeriesIntTaggedColumnar)(nil)
var _ Sizer = (*SeriesDoubleTaggedColumnar)(nil)

// string are all treated as ASCII and size of it is 1 byte * length
const characterSize = 1

// point are all 128 bit, 16 byte
const pointSize = 16 // int64 + float64|int64

func (m *Tag) RawSize() int {
	return m.RawMetaSize()
}

func (m *Tag) RawMetaSize() int {
	return characterSize * (len(m.K) + len(m.V))
}

func (*PointInt) RawSize() int {
	return pointSize
}

func (*PointInt) RawMetaSize() int {
	return 0
}

func (*PointDouble) RawSize() int {
	return pointSize
}

func (*PointDouble) RawMetaSize() int {
	return 0
}

func (m *PointIntTagged) RawSize() int {
	return m.RawMetaSize() + pointSize
}

func (m *PointIntTagged) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *PointDoubleTagged) RawSize() int {
	return m.RawMetaSize() + pointSize
}

func (m *PointDoubleTagged) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *EmptySeries) RawSize() int {
	// empty series is just meta
	return m.RawMetaSize()
}

func (m *EmptySeries) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesIntTagged) RawSize() int {
	return m.RawMetaSize() + pointSize*len(m.Points)
}

func (m *SeriesIntTagged) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesDoubleTagged) RawSize() int {
	return m.RawMetaSize() + pointSize*len(m.Points)
}

func (m *SeriesDoubleTagged) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesIntTaggedColumnar) RawSize() int {
	return m.RawMetaSize() + pointSize*len(m.Times)
}

func (m *SeriesIntTaggedColumnar) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesDoubleTaggedColumnar) RawSize() int {
	return m.RawMetaSize() + pointSize*len(m.Times)
}

func (m *SeriesDoubleTaggedColumnar) RawMetaSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}
