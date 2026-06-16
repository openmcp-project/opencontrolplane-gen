package testdata

type Example struct {
	namespace string
	conditionalName string
	// test nested command
	conditionalmyName int
}

func (r *Example) Reconciler() {
	var myName, myNamespace string
	_ = myName
	_ = myNamespace
}
