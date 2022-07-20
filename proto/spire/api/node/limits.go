package node

const (
	// Max burst values for ratelimiting
	// Requests containing more than this number of
	// operations will always be rejected
	AttestLimit     int = 1
	CSRLimit        int = 2147483647
	JSRLimit        int = 500
	PushJWTKeyLimit int = 500
)
