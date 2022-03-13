package cloudflare

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/ipfs/go-cid"
)

func SetRecord(key string, email string, domain string, cid cid.Cid) error {
	var api *cloudflare.API
	var err error
	if email != "" { // If email and key is provided
		api, err = cloudflare.New(key, email)
	} else { // If only token is provided
		api, err = cloudflare.NewWithAPIToken(key)
	}
	if err != nil {
		return err
	}

	// Get the zone ID
	var zone string
	zoneDomain := domain
	for { // Loop until at root domain
		zone, err = api.ZoneIDByName(zoneDomain)
		if zone != "" {
			if err != nil {
				return err
			}
			break
		}

		// Remove each subdomain (split ".", remove first, join, repeat)
		domainArr := strings.Split(zoneDomain, ".")
		if len(domainArr) == 2 {
			return fmt.Errorf("cannot find zone, are your credentials correct?")
		}
		domainArr = domainArr[1:]
		zoneDomain = strings.Join(domainArr, ".")
	}

	// Get DNSLink records for this domain
	ctx := context.Background()
	dnslinkRecords, err := api.DNSRecords(ctx, zone, cloudflare.DNSRecord{Name: fmt.Sprintf("_dnslink.%s", domain), Type: "TXT"})
	if err != nil {
		return err
	}

	// Decide whether to set or update record
	switch len(dnslinkRecords) {
	case 1: // Update existing record
		err := api.UpdateDNSRecord(ctx, zone, dnslinkRecords[0].ID, cloudflare.DNSRecord{Name: fmt.Sprintf("_dnslink.%s", domain), Type: "TXT", Content: fmt.Sprintf("dnslink=/ipfs/%s", cid.String())})
		if err != nil {
			return err
		}
	case 0: // Create record
		_, err := api.CreateDNSRecord(ctx, zone, cloudflare.DNSRecord{Name: fmt.Sprintf("_dnslink.%s", domain), Type: "TXT", Content: fmt.Sprintf("dnslink=/ipfs/%s", cid.String())})
		if err != nil {
			return err
		}
	default: // Uh-oh
		return fmt.Errorf("too many dns records")
	}

	// No errors, record set successfully!
	return nil
}
