package siprocket

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

*/

type sipContact struct {
	UriType string // Type of URI sip, sips, tel etc
	Name    []byte // Named portion of URI
	User    []byte // User part
	Host    []byte // Host part
	Port    []byte // Port number
	Tran    []byte // Transport
	Qval    []byte // Q Value
	Expires []byte // Expires
	Src     []byte // Full source if needed
}

func parseSipContact(v []byte, out *sipContact) {

	pos := 0
	state := FIELD_BASE

	// Init the output area
	out.UriType = ""
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

	// Loop through the bytes making up the line
	for pos < len(v) {
		// FSM
		//fmt.Println("POS:", pos, "CHR:", string(v[pos]), "STATE:", state)
		switch state {
		case FIELD_BASE:
			if v[pos] == '"' && out.UriType == "" {
				state = FIELD_NAMEQ
				pos++
				continue
			}
			if v[pos] != ' ' {
				// Not a space so check for uri types
				if getString(v, pos, pos+4) == "sip:" {
					state = FIELD_USER
					pos = pos + 4
					out.UriType = "sip"
					continue
				}
				if getString(v, pos, pos+5) == "sips:" {
					state = FIELD_USER
					pos = pos + 5
					out.UriType = "sips"
					continue
				}
				if getString(v, pos, pos+4) == "tel:" {
					state = FIELD_USER
					pos = pos + 4
					out.UriType = "tel"
					continue
				}
				// Look for a Q identifier
				if getString(v, pos, pos+2) == "q=" {
					state = FIELD_Q
					pos = pos + 2
					continue
				}
				// Look for a Expires identifier
				if getString(v, pos, pos+8) == "expires=" {
					state = FIELD_EXPIRES
					pos = pos + 8
					continue
				}
				// Look for a transport identifier
				if getString(v, pos, pos+10) == "transport=" {
					state = FIELD_TRAN
					pos = pos + 10
					continue
				}
				// Look for other identifiers and ignore
				if v[pos] == '=' {
					state = FIELD_IGNORE
					pos = pos + 1
					continue
				}
				// Check for other chrs
				if v[pos] != '<' && v[pos] != '>' && v[pos] != ';' && out.UriType == "" {
					state = FIELD_NAME
					continue
				}
			}

		case FIELD_NAMEQ:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Name = append(out.Name, v[pos])

		case FIELD_NAME:
			if v[pos] == '<' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Name = append(out.Name, v[pos])

		case FIELD_USER:
			if v[pos] == '@' {
				state = FIELD_HOST
				pos++
				continue
			}
			out.User = append(out.User, v[pos])

		case FIELD_HOST:
			if v[pos] == ':' {
				state = FIELD_PORT
				pos++
				continue
			}
			if v[pos] == ';' || v[pos] == '>' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Host = append(out.Host, v[pos])

		case FIELD_PORT:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Port = append(out.Port, v[pos])

		case FIELD_TRAN:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Tran = append(out.Tran, v[pos])

		case FIELD_Q:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Qval = append(out.Qval, v[pos])

		case FIELD_EXPIRES:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Expires = append(out.Expires, v[pos])

		case FIELD_IGNORE:
			if v[pos] == ';' || v[pos] == '>' {
				state = FIELD_BASE
				pos++
				continue
			}

		}
		pos++
	}
}
