// Copyright 2022 The forwarder Authors. All rights reserved.
// Use of this source code is governed by a MPL
// license that can be found in the LICENSE file.

package fileurl

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	uncEmptyAuthorityRegex = regexp.MustCompile(`^file:/{4,}([^/])`)
	windowsVolumeRegex     = regexp.MustCompile(`^/?([a-zA-Z])[:\|]/`)
)

// ParseFilePathOrURL extends url.Parse with the ability to parse file paths
// and adds extended support for URL file scheme as described in RFC 8089.
// If there is no scheme, it will be set to "file".
// If value equals "-", it will be set to "file://-" meaning stdin.
// See: https://datatracker.ietf.org/doc/html/rfc8089
func ParseFilePathOrURL(val string) (*url.URL, error) {
	// Handle stdin.
	if val == "-" {
		return &url.URL{Scheme: "file", Path: "-"}, nil
	}

	val = strings.ReplaceAll(val, "\\", "/")

	// Handle UNC paths.
	if strings.HasPrefix(val, "//") {
		val = "file:" + val
	}
	if m := uncEmptyAuthorityRegex.FindStringSubmatch(val); m != nil {
		val = "file://" + m[1] + val[len(m[0]):]
	}

	u, err := url.Parse(val)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = "file"
	}
	if u.Scheme != "file" {
		return u, nil
	}

	// Handle Windows paths.
	if u.Path == "" && u.Opaque != "" {
		u.Path, u.Opaque = u.Opaque, u.Path
	}
	if m := windowsVolumeRegex.FindStringSubmatch(u.Path); m != nil {
		u.Path = m[1] + ":/" + u.Path[len(m[0]):]
	}

	u.OmitHost = false // include host in the output
	return u, nil
}
