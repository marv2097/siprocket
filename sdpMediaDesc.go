package siprocket

/*
RFC4566 - https://tools.ietf.org/html/rfc4566#section-5.14

5.14.  Media Descriptions ("m=")

  m=<media> <port> <proto> <fmt> ...

A session description may contain a number of media descriptions.
Each media description starts with an "m=" field and is terminated by
either the next "m=" field or by the end of the session description.

eg:
m=audio 24414 RTP/AVP 8 18 101

*/

type sdpMediaDesc struct {
	MediaType []byte // Named portion of URI
	Port      []byte // Port number
	Proto     []byte // Protocol
	Fmt       []byte // Fmt
	Src       []byte // Full source if needed
}

func parseSdpMediaDesc(v []byte, out *sdpMediaDesc) {

	pos := 0
	state := FIELD_MEDIA

	// Init the output area
	out.MediaType = nil
	out.Port = nil
	out.Proto = nil
	out.Fmt = nil
	out.Src = nil

	// Keep the source line if needed
	if keep_src {
		out.Src = v
	}

	// Loop through the bytes making up the line
	for pos < len(v) {
		// FSM
		switch state {
		case FIELD_MEDIA:
			if v[pos] == ' ' {
				state = FIELD_PORT
				pos++
				continue
			}
			out.MediaType = append(out.MediaType, v[pos])

		case FIELD_PORT:
			if v[pos] == ' ' {
				state = FIELD_PROTO
				pos++
				continue
			}
			out.Port = append(out.Port, v[pos])

		case FIELD_PROTO:
			if v[pos] == ' ' {
				state = FIELD_FMT
				pos++
				continue
			}
			out.Proto = append(out.Proto, v[pos])

		case FIELD_FMT:
			out.Fmt = append(out.Fmt, v[pos])
		}
		pos++
	}
}
