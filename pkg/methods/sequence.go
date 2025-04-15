package methods

type (
	Sequence []SequenceItem

	SequenceItem struct {
		Method Method
		Args   []any
		Result any
		Err    error
	}
)
