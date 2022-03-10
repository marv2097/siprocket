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

	// Check if our uri string uses <> encapsulation
	// Although <> is not a reserved charactor so its possible we can go wrong
	// If there is a name string then encapsultion must be used.
	if idx = bytes.LastIndexByte(v, byte('>')); idx > -1 {

		// parse header parameters of the encapulated form
		parseSipReqHeaderParams(v[idx:], out)
		v = v[:idx]

		if idx = bytes.LastIndexByte(v, byte('<')); idx == -1 {
			return errors.New("found ending encapsualtion > but not staring <")
		}

		// Extract the name field
		//out.Name = v[:idx]

		// clean up out.Name
		//out.Name = bytes.Trim(out.Name, ` `)
		//out.Name = bytes.Trim(out.Name, `"`)

		v = v[idx+1:]

		// Next we'll find that method SIP(S)
		// Whilse the protocol allows the use 352 URI schema (we are only supporting sip)
		// https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml
		if idx = bytes.Index(v, []byte("sip:")); idx > -1 {
			out.UriType = v[idx : idx+3]
			v = v[idx+4:]
		} else if idx = bytes.Index(v, []byte("sips:")); idx > -1 {
			out.UriType = v[idx : idx+4]
			v = v[idx+5:]
		} else {
			return errors.New("unsupport URI-Schema found")
		}

		// Next find if userinfo is present denoted by @ (reserved charactor)
		if idx = bytes.IndexByte(v, byte('@')); idx > -1 {
			out.User = v[:idx]
			v = v[idx+1:]
		}

		// Trim of the password from the user section
		if idx = bytes.IndexByte(out.User, byte(':')); idx > -1 {
			out.User = out.User[:idx]
		}

		// Apply fix for a non complient ua
		if idx = bytes.IndexByte(out.User, byte(';')); idx > -1 {
			//out.Params = append(out.Params, out.User[idx+1:])
			out.User = out.User[:idx]
		}

		// Extract the URL parameters
		// These can only be located inside the encapsulated form
		for {
			if idx = bytes.LastIndexByte(v, byte(';')); idx == -1 {
				break
			}
			//out.Params = append(out.Params, v[idx+1:])
			v = v[:idx]
		}

		// remote any port
		if idx = bytes.IndexByte(v, byte(':')); idx > -1 {
			out.Port = v[idx+1:]
			v = v[:idx]
		}

		// all that is left is the host
		out.Host = v

	} else {
		// Parse header parameters of the non encapsulated form

		// Next we'll find that method SIP(S)
		// Whilse the protocol allows the use 352 URI schema (we are only supporting sip)
		// https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml
		if idx = bytes.Index(v, []byte("sip:")); idx > -1 {
			out.UriType = v[idx : idx+3]
			v = v[idx+4:]
		} else if idx = bytes.Index(v, []byte("sips:")); idx > -1 {
			out.UriType = v[idx : idx+4]
			v = v[idx+5:]
		} else {
			return errors.New("unsupport URI-Schema found")
		}

		// Next find if userinfo is present denoted by @ (reserved charactor)
		if idx = bytes.IndexByte(v, byte('@')); idx > -1 {
			out.User = v[:idx]
			v = v[idx+1:]
		}

		// Trim of the password from the user section
		if idx = bytes.IndexByte(out.User, byte(':')); idx > -1 {
			out.User = out.User[:idx]
		}
	
		// Apply fix for a non complient ua
		if idx = bytes.IndexByte(out.User, byte(';')); idx > -1 {
			//out.Params = append(out.Params, out.User[idx+1:])
			out.User = out.User[:idx]
		}
	
		// In the non encapsulated the query form is possible
		if idx = bytes.LastIndexByte(v, byte('?')); idx > -1 {
			// parse header parameters
			parseSipReqHeaderParams(v[idx:], out)
			v = v[:idx]
			// Extract the URL parameters
			// only available if the query form is used
			for {
				if idx = bytes.LastIndexByte(v, byte(';')); idx == -1 {
					break
				}
				//out.Params = append(out.Params, v[idx+1:])
				v = v[:idx]
			}
		} else {
			// Parse header parameters
			if idx = bytes.LastIndexByte(v, byte(';')); idx > -1 {
				parseSipReqHeaderParams(v[idx:], out)
				v = v[:idx]
			}
		}

		// remote any port
		if idx = bytes.IndexByte(v, byte(':')); idx > -1 {
			out.Port = v[idx+1:]
			v = v[:idx]
		}
		
		// all that is left is the host
		out.Host = v
	}

	return nil
}

func parseSipResp(v []byte, out *sipReq) error {
	var idx int

	// Get descriptions last bit
	if idx = bytes.LastIndex(v, []byte(" ")); idx > -1 {
		out.StatusDesc = v[idx+1:]
		v = v[:idx]
	}

	// Get statuscode middle bit.
	if idx = bytes.LastIndex(v, []byte(" ")); idx > -1 {
		out.StatusCode = v[idx+1:]
		v = v[:idx]
	}

	return nil
}

func parseSipReqHeaderParams(v []byte, out *sipReq) {
	var idx int

	for {
		if idx = bytes.LastIndexByte(v[idx:], byte(';')); idx == -1 {
			break
		}

		if len(v[idx:]) > 5 {
			if string(v[idx:idx+6]) == ";user=" {
				out.UserType = v[idx+6:]
				return
			}
		}
		v = v[:idx]
	}
}
