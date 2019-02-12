package dummy

const (
	GET = iota
	POST
)

var (
	methods = map[int]string{
		GET:  "get",
		POST: "post",
	}
)
