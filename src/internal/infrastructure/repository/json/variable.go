package json

const (
	jsonPath = "test/json"

	GET = iota + 1
	POST
)

var (
	methods = map[int]string{
		GET:  "get",
		POST: "post",
	}
)
