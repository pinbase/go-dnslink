# go-dnslink 🔗

[![](https://img.shields.io/badge/made%20by-Pinbase-blue.svg?style=flat-square)](https://pinbase.io)
[![](https://img.shields.io/github/license/pinbase/go-dnslink)](https://github.com/pinbase/go-dnslink)
[![](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

This is a module to programmatically create and update DNSLink 🔗 records from various DNS providers.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-dnslink` is a standard Go module which can be installed with:

```sh
go get github.com/pinbase/go-dnslink
```

## Usage

Import the module as follows

```go
import (
	...
	"github.com/pinbase/go-dnslink"
)
```

### Cloudflare

The `"DOMAIN"` value should be the full domain you want to assign a DNSLink record to, e.g. "pinbase.io".

The `CID` value should be a CID from [go-cid](https://github.com/ipfs/go-cid).

#### API token

```go
err := dnslink.Cloudflare("API_TOKEN", "", "DOMAIN", CID)
```

You can generate an API token in the Cloudflare dashboard [https://dash.cloudflare.com/profile/api-tokens](https://dash.cloudflare.com/profile/api-tokens), using the `Edit zone DNS` template.

#### API key and email

```go
err := dnslink.Cloudflare("API_KEY", "EMAIL_ADDRESS", "DOMAIN", CID)
```

Or you can provide your account's API key and email address.

## Contribute

PRs are welcome!

Especially good PRs would be ones implementing more DNS providers.

If updating the README, please stick to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © [Pinbase](https://pinbase.io)
