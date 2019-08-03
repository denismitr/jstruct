package fixtures

type country struct {
	name     string
	timezone string
}

type address struct {
	street  string
	house   float64
	country country
}

type user struct {
	name    string
	surname string
	address address
}

type f3 struct {
	user user
}
