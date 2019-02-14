package dummy

const (
	jsonPath = "test/json"

	GET = iota
	POST
)

var (
	methods = map[int]string{
		GET:  "get",
		POST: "post",
	}
)
