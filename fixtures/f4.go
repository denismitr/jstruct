package fixtures

type f4 struct {
	admin admin
	guest guest
}

type admin struct {
	expiresAt float64
	roles     []string
}

type guest struct {
	addresses []object2443843572
	name      string
	surname   string
}

type object2443843572 struct {
	house  float64
	state  state
	street string
}

type state struct {
	name     string
	timezone string
}
