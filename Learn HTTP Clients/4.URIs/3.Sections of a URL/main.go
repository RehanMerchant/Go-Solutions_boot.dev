package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: func() string {
			if password, ok := parsedUrl.User.Password(); ok {
				return password
			}
			return ""
		}(),
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}
