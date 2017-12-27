package whois

type Whois interface {
	Lookup(domain string) (bool, error)
}
