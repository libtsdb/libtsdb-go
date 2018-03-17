package libtsdbpb

// sizer is used to output raw size of series in byte
//
// string are all treated as ASCII and size of it is 1 byte * length
const characterSize = 1
const pointSize = 16 // int64 + float64|int64

func (m *Tag) RawSize() int {
	return characterSize * (len(m.K) + len(m.V))
}

func (*PointInt) RawSize() int {
	return pointSize
}

func (*PointDouble) RawSize() int {
	return pointSize
}

func (m *PointIntTagged) RawSize() int {
	s := characterSize*len(m.Name) + pointSize
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *PointDoubleTagged) RawSize() int {
	s := characterSize*len(m.Name) + pointSize
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *EmptySeries) RawSize() int {
	s := characterSize * len(m.Name)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesIntTagged) RawSize() int {
	s := characterSize*len(m.Name) + pointSize*len(m.Points)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesDoubleTagged) RawSize() int {
	s := characterSize*len(m.Name) + pointSize*len(m.Points)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesIntTaggedColumnar) RawSize() int {
	s := characterSize*len(m.Name) + pointSize*len(m.Times)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesDoubleTaggedColumnar) RawSize() int {
	s := characterSize*len(m.Name) + pointSize*len(m.Times)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}
