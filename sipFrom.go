package siprocket

/*
Parses a single line that is in the format of a from line, v
Also requires a pointer to a struct of type sipFrom to write output to

RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt - 8.1.1.3 From

*/

type sipFrom struct {
	UriType  string // Type of URI sip, sips, tel etc
	Name     []byte // Named portion of URI
	User     []byte // User part
	Host     []byte // Host part
	Port     []byte // Port number
	Tag      []byte // Tag
	UserType []byte // User Type
	Src      []byte // Full source if needed
}

func parseSipFrom(v []byte, out *sipFrom) {

	pos := 0
	state := FIELD_BASE

	// Init the output area
	out.UriType = ""
	out.Name = nil
	out.User = nil
	out.Host = nil
	out.Port = nil
	out.Tag = nil
	out.UserType = nil
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
				// Look for a Tag identifier
				if getString(v, pos, pos+4) == "tag=" {
					state = FIELD_TAG
					pos = pos + 4
					continue
				}
				// Look for other identifiers and ignore
				if v[pos] == '=' {
					state = FIELD_IGNORE
					pos = pos + 1
					continue
				}
				// Look for a User Type identifier
				if getString(v, pos, pos+5) == "user=" {
					state = FIELD_USERTYPE
					pos = pos + 5
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

		case FIELD_USERTYPE:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.UserType = append(out.UserType, v[pos])

		case FIELD_TAG:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Tag = append(out.Tag, v[pos])

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
