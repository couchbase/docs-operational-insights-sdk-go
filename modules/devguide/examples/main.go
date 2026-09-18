package main

const (
	connStr  = "couchbases://..."
	username = "..."
	password = "..."
)

func main() {}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
