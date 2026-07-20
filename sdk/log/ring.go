package log

type ring struct {
	next, prev *ring
	Value      Record
}

func (r *ring) init() *ring {
	r.next = r
	r.prev = r
	return r
}

func (r *ring) Next() *ring { _ = "STUB: not implemented"; return nil }

func (r *ring) Prev() *ring { _ = "STUB: not implemented"; return nil }

func newRing(n int) *ring { _ = "STUB: not implemented"; return nil }

func (r *ring) Len() int { _ = "STUB: not implemented"; return 0 }

func (r *ring) Do(f func(Record)) { _ = "STUB: not implemented"; return }
