package fixtures

type name struct {
	first  string
	last   string
	phones []float64
}

type f1 struct {
	pets []interface{}
	prof string
	age  float64
	name name
}
