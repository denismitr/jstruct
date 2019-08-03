package fixtures

type f1 struct {
	age  float64
	name name
	pets []interface{}
	prof string
}

type name struct {
	first  string
	last   string
	phones []float64
}
