package siprocket

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt

INVITE sip:01798300765@87.252.61.202;user=phone SIP/2.0
SIP/2.0 200 OK

*/

import (
	"bytes"
	"errors"
)

type sipReq struct {
	Method     []byte // Sip Method eg INVITE etc
	UriType    []byte // Type of URI sip, sips, tel etc
	StatusCode []byte // Status Code eg 100
	StatusDesc []byte // Status Code Description eg trying
	User       []byte // User part
	Host       []byte // Host part
	Port       []byte // Port number
	UserType   []byte // User Type
	Src        []byte // Full source if needed
}

// request or response line
func parseSipReq(v []byte, out *sipReq) error {

	var idx int

	// Init the output area
	out.Method = nil
	out.UriType = nil
	out.StatusCode = nil
	out.StatusDesc = nil
	out.User = nil
	out.Host = nil
	out.Port = nil
	out.UserType = nil

	// Keep the source line if needed
	out.Src = v

	// Don't process impossibly valid headers
	if len(v) < 10 {
		return errors.New("too short to be valid")
	}

	// Redirect response headsers, these must start SIP always
	if string(v[:3]) == "SIP" {
		return parseSipResp(v, out)
	}

	// Strip of request and prototol
	// These can be disguished from the URI as being the first and last segment seporated by a single space charactor
	if idx = bytes.LastIndex(v, []byte(" ")); idx > -1 {
		v = v[:idx]
	}
	if idx = bytes.Index(v, []byte(" ")); idx > -1 {
		out.Method = v[:idx]
		if len(v) > idx {
			v = v[idx+1:]
		}
	}

	// Next if our uri string uses <> encapsulation our end charactor should be >
	if idx = bytes.LastIndexByte(v, byte('>')); idx == len(v)-1 {
		if idx = bytes.IndexByte(v, byte('<')); idx > -1 {
			//		out.Name = v[:idx]
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
		//	out.Params = append(out.Params, v[idx+1:]

		// recover user parameter if it exist
		if len(v[idx+1:]) > 5 {
			if string(v[idx+1:idx+6]) == "user=" {
				out.UserType = v[idx+6:]
			}
		}
		v = v[:idx]
	}

	// Next we'll find that method SIP(S)
	// Whilse the protocol allows the use 352 URI schema (we are only supporting sip)
	// https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml
	if idx = bytes.Index(v, []byte("sip:")); idx > -1 {
		if idx > 0 {
			//		out.Name = v[:idx]
		}
		out.UriType = v[idx : idx+3]
		v = v[idx+4:]
	} else if idx = bytes.Index(v, []byte("sips:")); idx > -1 {
		if idx > 0 {
			//		out.Name = v[:idx]
		}
		out.UriType = v[idx : idx+4]
		v = v[idx+5:]
	} else {
		return errors.New("no UriType found")
	}

	// clean up out.Name
	//out.Name = bytes.Trim(out.Name, ` `)
	//out.Name = bytes.Trim(out.Name, `"`)

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

func parseSipResp(v []byte, out *sipReq) error {
	var idx int

	// Get descriptions last bit
	if idx = bytes.LastIndex(v, []byte(" ")); idx > -1 {
		out.StatusDesc = v[idx+1:]
		v = v[:idx]
	}

	// Get statuscode middile bit.
	if idx = bytes.LastIndex(v, []byte(" ")); idx > -1 {
		out.StatusCode = v[idx+1:]
		v = v[:idx]
	}

	return nil
}
