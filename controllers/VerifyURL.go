package controllers

import "net/url"

// VerifyURL checks to see if the user provided URL contains an https or http prefix and adds
// adds it if not present. Doesn't verify it the link itself is valid or checks for tlds etc.
func VerifyURL(uRequest string) string {
	// ParseRequestURI assumes there are no fragments "#", check behaviour
	_, err := url.ParseRequestURI(uRequest)
	if err != nil {
		uRequest = ("https://" + uRequest)
	}
	return uRequest
}
