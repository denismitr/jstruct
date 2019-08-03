package fixtures

type f3 struct {
	user user
}

type user struct {
	address address
	name    string
	surname string
}

type address struct {
	country country
	house   float64
	street  string
}

type country struct {
	name     string
	timezone string
}
