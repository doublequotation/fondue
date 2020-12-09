package parent

type Child struct {
	url  string
	port string
}

type Host struct {
	url        string
	port       string
	auth       string
	host       string
	childNodes string
}
