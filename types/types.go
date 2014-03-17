package strowger

type RouteType uint8

const (
	RouteHTTP RouteType = iota
	RouteTCP
)

type Config struct {
	Service string
	Type    RouteType

	// HTTP
	HTTPDomain string
	HTTPSCert  []byte // DER encoded certificate
	HTTPSKey   []byte // DER encoded private key
}
