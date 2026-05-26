// Package libstring provides various string related functions.
package libstring

import (
	"net/http"

	"github.com/didip/tollbooth/v8/limiter"
)

// StringInSlice finds needle in a slice of strings.
func StringInSlice(sliceString []string, needle string) bool {
	_ = "STUB: not implemented"
	return false
}

// RemoteIPFromIPLookup picks an ip address explicitly from limiter.IPLookup criteria.
// This function is intended to replace RemoteIP function.
func RemoteIPFromIPLookup(ipLookup limiter.IPLookup, r *http.Request) string {
	_ = "STUB: not implemented"
	return ""
}

// 1. Cover the basic use cases for both ipv4 and ipv6

// 2. Upon error, just return the remote addr.

// CanonicalizeIP returns a form of ip suitable for comparison to other IPs.
// For IPv4 addresses, this is simply the whole string.
// For IPv6 addresses, this is the /64 prefix.
func CanonicalizeIP(ip string) string {
	_ = "STUB: not implemented"

	// This is how net.ParseIP decides if an address is IPv6
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.7:src/net/ip.go;l=704
	return ""
}

// IPv4

// IPv6

// Not an IP address at all

// By default, the string representation of a net.IPNet (masked IP address) is just
// "full_address/mask_bits". But using that will result in different addresses with
// the same /64 prefix comparing differently. So we need to zero out the last 64 bits
// so that all IPs in the same prefix will be the same.
//
// Note: When 1.18 is the minimum Go version, this can be written more cleanly like:
// netip.PrefixFrom(netip.MustParseAddr(ipv6), 64).Masked().Addr().String()
// (With appropriate error checking.)

// Note that this doesn't have the "/64" suffix customary with a CIDR representation,
// but those three bytes add nothing for us.
