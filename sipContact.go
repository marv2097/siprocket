package siprocket

import (
	"bytes"
	"errors"
)

/*

RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt - 8.1.1.8 Contact

   The Contact header field provides a SIP or SIPS URI that can be used
   to contact that specific instance of the UA for subsequent requests.
   The Contact header field MUST be present and contain exactly one SIP
   or SIPS URI in any request that can result in the establishment of a
   dialog.

Examples:

   Contact: "Mr. Watson" <sip:watson@worcester.bell-telephone.com>
      ;q=0.7; expires=3600,
      "Mr. Watson" <mailto:watson@bell-telephone.com> ;q=0.1
   m: <sips:bob@192.0.2.4>;expires=60


    sip:user:password@host:port;header-parameters
    sip:user:password@host:port;uri-parameters?headers-parameters
	<sip:user:password@host:port;uri-parameters>headers-parameters
	display name <user:password@host:port;uri-parameters>headers-parameters
	"display name" <user:password@host:port;uri-parameters>headers-parameters
*/

type sipContact struct {
	UriType []byte // Type of URI sip, sips, tel etc
	Name    []byte // Named portion of URI
	User    []byte // User part
	Host    []byte // Host part
	Port    []byte // Port number
	Tran    []byte // Transport
	Qval    []byte // Q Value
	Expires []byte // Expires
	Src     []byte // Full source if needed
}

func parseSipContact(v []byte, out *sipContact) error {

	var idx int

	// Init the output area
	out.UriType = nil
	out.Name = nil
	out.User = nil
	out.Host = nil
	out.Port = nil
	out.Tran = nil
	out.Qval = nil
	out.Expires = nil
	out.Src = nil

	// Keep the source line if needed
	if keep_src {
		out.Src = v
	}

	// Check if our uri string uses <> encapsulation
	// Although <> is not a reserved charactor so its possible we can go wrong
	// If there is a name string then encapsultion must be used.
	if idx = bytes.LastIndexByte(v, byte('>')); idx > -1 {

		// parse header parameters of the encapulated form
		parseSipContactHeaderParams(v[idx:], out)
		v = v[:idx]

		if idx = bytes.LastIndexByte(v, byte('<')); idx == -1 {
			return errors.New("found ending encapsualtion > but not staring <")
		}

		// Extract the name field
		out.Name = v[:idx]

		// clean up out.Name
		out.Name = bytes.Trim(out.Name, ` `)
		out.Name = bytes.Trim(out.Name, `"`)

		v = v[idx+1:]

		// Extract the URL parameters
		// These can only be located inside the encapsulated form
		for {
			if idx = bytes.LastIndexByte(v, byte(';')); idx == -1 {
				break
			}
			//out.Params = append(out.Params, v[idx+1:])
			v = v[:idx]
		}
	} else {
		// Parse header parameters of the non encapsulated form

		// If its in the query form
		if idx = bytes.LastIndexByte(v, byte('?')); idx > -1 {

			// parse header parameters
			parseSipContactHeaderParams(v[idx:], out)
			v = v[:idx]

			// Extract the URL parameters
			for {
				if idx = bytes.LastIndexByte(v, byte(';')); idx == -1 {
					break
				}
				//	out.Params = append(out.Params, v[idx+1:])
				v = v[:idx]
			}
		} else {
			// Parse header parameters
			if idx = bytes.LastIndexByte(v, byte(';')); idx > -1 {
				parseSipContactHeaderParams(v[idx:], out)
				v = v[:idx]
			}
		}
	}

	// Next we'll find that method SIP(S)
	// Whilse the protocol allows the use 352 URI schema (we are only supporting sip)
	// https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml
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
		return errors.New("unsupport URI-Schema found")
	}

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

func parseSipContactHeaderParams(v []byte, out *sipContact) {
	var idx int

	for {
		if idx = bytes.LastIndexByte(v[idx:], byte(';')); idx == -1 {
			break
		}

		if len(v[idx:]) < 3 {
			v = v[:idx]
			continue
		}

		if string(v[idx:idx+3]) == ";q=" {
			out.Qval = v[idx+3:]
			v = v[:idx]
			continue
		}

		if len(v[idx:]) < 9 {
			v = v[:idx]
			continue
		}

		if string(v[idx:idx+9]) == ";expires=" {
			out.Expires = v[idx+9:]
			v = v[:idx]
			continue
		}
	}
}
