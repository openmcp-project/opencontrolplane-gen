package testdata

type Example struct {
	namespace string

	conditionalName string
	// test nested command
	conditionalValue int
}

func (r *Example) Reconciler() {
	var Value, myNamespace string
	_ = Value
	_ = myNamespace
}
