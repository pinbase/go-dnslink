package dnslink

import "github.com/pinbase/go-dnslink/providers/cloudflare"

// Returns the package version
func Version() string {
	return "0.1.0"
}

// Export all the DNS providers here
var Cloudflare = cloudflare.SetRecord
