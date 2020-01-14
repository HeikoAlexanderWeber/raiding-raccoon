package crawler

import (
	"net"
	"net/url"
)

// Selector type.
// Interface for selectors.
type Selector func(*url.URL) bool

// DomainSelector func.
// Only approves URIs that start with the given domain.
func DomainSelector(domain string) Selector {
	return func(url *url.URL) bool {
		host := RefineHostname(url)
		if domain != host {
			return false
		}
		return true
	}
}

// UniqueSelector func.
// Only approves URIs one time.
func UniqueSelector(setIfAbsent func(string) bool) Selector {
	return func(url *url.URL) bool {
		return setIfAbsent(url.String())
	}
}

// RefineHostname func.
// Splits Port from Hostname etc. Only returns the valid plain hostname.
func RefineHostname(url *url.URL) string {
	host := url.Host
	if sHost, _, err := net.SplitHostPort(host); err == nil {
		host = sHost
	}
	return host
}
