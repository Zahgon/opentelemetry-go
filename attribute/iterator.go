package attribute

type Iterator struct {
	storage *Set
	idx     int
}

type MergeIterator struct {
	one     oneIterator
	two     oneIterator
	current KeyValue
}

type oneIterator struct {
	iter Iterator
	done bool
	attr KeyValue
}

func (i *Iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (i *Iterator) Label() KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (i *Iterator) Attribute() KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (i *Iterator) IndexedLabel() (int, KeyValue) {
	_ = "STUB: not implemented"
	return 0, *new(KeyValue)
}

func (i *Iterator) IndexedAttribute() (int, KeyValue) {
	_ = "STUB: not implemented"
	return 0, *new(KeyValue)
}

func (i *Iterator) Len() int { _ = "STUB: not implemented"; return 0 }

func (i *Iterator) ToSlice() []KeyValue { _ = "STUB: not implemented"; return nil }

func NewMergeIterator(s1, s2 *Set) MergeIterator {
	_ = "STUB: not implemented"
	return *new(MergeIterator)
}

func makeOne(iter Iterator) oneIterator { _ = "STUB: not implemented"; return *new(oneIterator) }

func (oi *oneIterator) advance() { _ = "STUB: not implemented"; return }

func (m *MergeIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (m *MergeIterator) Label() KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (m *MergeIterator) Attribute() KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }
