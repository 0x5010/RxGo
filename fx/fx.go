package fx

type (

	// EmittableFunc defines a function that should be used with Start operator.
	EmittableFunc func() interface{}

	// MappableFunc defines a function that acts as a predicate to the Map operator.
	MappableFunc func(interface{}) interface{}

	// ScannableFunc defines a function that acts as a predicate to the Scan operator.
	ScannableFunc func(interface{}, interface{}) interface{}

	// FilterableFunc defines a func that should be passed to the Filter operator.
	FilterableFunc func(interface{}) bool
)