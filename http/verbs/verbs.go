package verbs

type Verb string

const (
	GET     Verb = "GET"
	HEAD    Verb = "HEAD"
	POST    Verb = "POST"
	PUT     Verb = "PUT"
	DELETE  Verb = "DELETE"
	CONNECT Verb = "CONNECT"
	OPTIONS Verb = "OPTIONS"
	TRACE   Verb = "TRACE"
	PATCH   Verb = "PATCH"
)
