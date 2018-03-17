package libtsdbpb

// sizer is used to output raw size of series in bit
//
// string are all treated as ASCII and size of it is 8 * length
const characterSize = 8

func (m *Tag) RawSize() int {
	return characterSize * (len(m.K) + len(m.V))
}

func (*PointInt) RawSize() int {
	return 64 + 64
}

func (*PointDouble) RawSize() int {
	return 64 + 64
}

func (m *PointIntTagged) RawSize() int {
	s := characterSize*len(m.Name) + m.Point.RawSize()
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *PointDoubleTagged) RawSize() int {
	s := characterSize*len(m.Name) + m.Point.RawSize()
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
	s := characterSize*len(m.Name) + (64+64)*len(m.Points)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesDoubleTagged) RawSize() int {
	s := characterSize*len(m.Name) + (64+64)*len(m.Points)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesIntTaggedColumnar) RawSize() int {
	s := characterSize*len(m.Name) + (64+64)*len(m.Times)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}

func (m *SeriesDoubleTaggedColumnar) RawSize() int {
	s := characterSize*len(m.Name) + (64+64)*len(m.Times)
	for i := 0; i < len(m.Tags); i++ {
		s += m.Tags[i].RawSize()
	}
	return s
}
