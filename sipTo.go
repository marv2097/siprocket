package siprocket

import (
	"bytes"
	"errors"
)

// Parses a single line that is in the format of a to line, v
// Also requires a pointer to a struct of type sipTo to write output to
// RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt - 8.1.1.2 To

type sipTo struct {
	UriType []byte   // Type of URI sip, sips, tel etc
	Name    []byte   // Named portion of URI
	User    []byte   // User part
	Host    []byte   // Host part
	Port    []byte   // Port number
	Parms   [][]byte // Arrray of URI prams
	Tag     []byte   // Tag
	Src     []byte   // Full source if needed
}

/* Examples
sip:user:password@host:port;uri-parameters?headers

"name"sip:who@where:port;tag=
name <sip:who@where:port>;tag=
sip:who@where:port;tag=
<sip:who@where:port;uri-param=>;tag=
*/

func parseSipTo(v []byte, out *sipTo) error {

	var idx int

	// Init the output area
	out.Name = nil
	out.User = nil
	out.Host = nil
	out.Parms = nil
	out.Port = nil
	out.Tag = nil

	// Keep the source line if needed
	out.Src = v

	// Probably easier to strip any tag from the end
	// ; and = are reserved charactors so this should not be found elsewhere
	if idx = bytes.LastIndex(v, []byte(";tag=")); idx > -1 {
		out.Tag = v[idx+5:]
		v = v[:idx]
	}

	// Next if our uri string uses <> encapsulation our end charactor should be >
	if idx = bytes.LastIndexByte(v, byte('>')); idx == len(v)-1 {
		if idx = bytes.IndexByte(v, byte('<')); idx > -1 {
			out.Name = v[:idx]
			v = v[idx+1 : len(v)-1]
		} else {
			return errors.New("found ending encapsualtion > but not staring <")
		}
	}

	// Split off the params
	for {
		if idx = bytes.LastIndexByte(v, byte(';')); idx == -1 {
			break
		}
		out.Parms = append(out.Parms, v[idx+1:])
		v = v[:idx]
	}

	// Next we'll find that method SIP(S)
	if idx = bytes.Index(v, []byte("sip:")); idx > -1 {
		if idx > 0 {
			out.Name = v[:idx]
		}
		out.UriType = v[idx : idx+3]
		v = v[idx+4:]
	} else if idx = bytes.Index(v, []byte("sips:")); idx > -1 {
		if idx > 0 {
			out.Name = v[:idx]
		}
		out.UriType = v[idx : idx+4]
		v = v[idx+5:]
	} else {
		return errors.New("no UriType found")
	}

	// clean up out.Name
	out.Name = bytes.Trim(out.Name, ` `)
	out.Name = bytes.Trim(out.Name, `"`)

	// Next find if userinfo is present denoted by @ (reserved charactor)
	if idx = bytes.IndexByte(v, byte('@')); idx > -1 {
		out.User = v[:idx]
		v = v[idx+1:]
	}

	// remote any port
	if idx = bytes.IndexByte(v, byte(':')); idx > -1 {
		out.Port = v[idx+1:]
		v = v[:idx]
	}

	// all that is left is the host
	out.Host = v

	return nil
}
