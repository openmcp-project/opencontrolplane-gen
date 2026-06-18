//go:generate opencontrolplane-gen
package testdata

// opencontrolplane-gen:replace TestReconciler=RECONCILER_NAME
type TestReconciler struct {
	namespace string

	//opencontrolplane-gen:if OPTIONAL_FIELDS=include
	conditionalName string
	// test nested command
	//opencontrolplane-gen:replace Replace=FIELD_VALUE
	conditionalReplace int
	//opencontrolplane-gen:fi
}

// opencontrolplane-gen:replace TestReconciler=RECONCILER_NAME
func (r *TestReconciler) Reconciler() {
	//opencontrolplane-gen:replace provider=FIELD_VALUE namespace=FIELD_NAMESPACE
	var provider, namespace string
	//opencontrolplane-gen:replace provider=FIELD_VALUE
	_ = provider
	//opencontrolplane-gen:replace namespace=FIELD_NAMESPACE
	_ = namespace
}
