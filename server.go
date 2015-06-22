package ahsay

type Scheme int

const (
	HTTP  Scheme = iota + 1
	HTTPS Scheme = iota + 1
)

func (s Scheme) String() string {
	switch s {
	case HTTP:
		return "http"
	case HTTPS:
		return "https"
	default:
		return "scheme not set"
	}
}

type Server struct {
	Scheme   Scheme
	Hostname string
	Username string
	Password string
}
