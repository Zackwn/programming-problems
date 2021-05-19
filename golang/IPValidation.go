package main

import "regexp"

func Is_valid_ip(ip string) bool {
	var validIP = regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$`)
	return validIP.MatchString(ip)
}
