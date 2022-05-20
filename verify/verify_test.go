package verify

import "testing"

func TestIPv4Execute(t *testing.T) {
	verify := &IPv4Verifier{Config: Config{}}
	verify.Execute()
}

func TestIPv6Execute(t *testing.T) {
	verify := &IPv6Verifier{Config: Config{}}
	verify.Execute()
}
