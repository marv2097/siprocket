# siprocket
Fast SIP and SDP Parser

![Alt](https://travis-ci.org/marv2097/siprocket.svg?branch=master "Travis Build")


siprocket is intended for Monitoring applications that need to parse SIP messages on the fly. It allows fast and structured access to the most commonly needed fields from both the SIP header and SDP payload. While intended for use in packet capture systems it could also be adapted to SIP Client and Server tasks.

### Performance:

Without concurrency siprocket can parse approx 100k messages per second on a average xeon CPU. Depending on your application you may be able to parse concurrently which would greatly increase the throughput. The size and complexity of the SIP messages you have to parse will also influence performance. 

### Install:

Install using `go get -u github.com/marv2097/siprocket`
Add the Import `"github.com/marv2097/siprocket"` in your code.

### Usage:

siprocket expects data from a capture interface, socket or file as a slice of bytes. In most applications you will have already parsed the lower layers and pass the SIP payload to siprocket to be parsed. The output from the Parse function is a stuct that contains all of the top level items found in the SIP message. These can then be accessed using a simple dot notation. Most outputs are also slices of bytes with the exception of a few fields. 

In this simple example we will use a SIP message defined directly in our code below. It will parse the message and print out the user-part of the 'From' and 'To' header fields.

```go
// Load up a test message
raw := []byte("SIP/2.0 200 OK\r\n" +
              "Via: SIP/2.0/UDP 192.168.2.242:5060;received=22.23.24.25;branch=z9hG4bK5ea22bdd74d079b9;alias;rport=5060\r\n" +
              "To: <sip:JohnSmith@mycompany.com>;tag=aprqu3hicnhaiag03-2s7kdq2000ob4\r\n" +
              "From: sip:HarryJones@mycompany.com;tag=89ddf2f1700666f272fb861443003888\r\n" +
              "CSeq: 57413 REGISTER\r\n" +
              "Call-ID: b5deab6380c4e57fa20486e493c68324\r\n" +
              "Contact: <sip:JohnSmith@192.168.2.242:5060>;expires=192\r\n\r\n")

// Parse the sip data
sip := siprocket.Parse(raw)

// Print out the To User-part.
fmt.Print("From: ", string(sip.From.User), " To: ", string(sip.To.User))

```
Will Print `From: HarryJones To: JohnSmith`

### Output Data Structure

Many of the SIP headers are in simple key value pairs. For example the Call-ID field, these kinds of fields all share the same format used to store them. It has a slice of bytes for the value, and an optional source variable.

```go
type sipVal struct {
	Value []byte // Sip Value
	Src   []byte // Full source if needed
}
```

To access them we just reference the field name followed by the word `Value`. So if we wanted to get the Call-ID from the example above we would just reference `sip.CallId.Value`. Other fields that support this format are:

Header Field | Reference
--- | ---
User-Agent | `Ua`
Expires | `Exp`
Max-Forwards | `MaxFwd`
Call-Id | `CallId`
Content-Type | `ContType`
Content-Length | `ContLen`

More complicated fields have specific structs to hold the data they contain, for example the From and To header field has each section broken out and are identical:

```go
type sipTo struct {
	UriType  string // Type of URI sip, sips, tel etc
	Name     []byte // Named portion of URI
	User     []byte // User part
	Host     []byte // Host part
	Port     []byte // Port number
	Tag      []byte // Tag
	UserType []byte // User Type
	Src      []byte // Full source if needed
}
```

#### Multiple values

When a field is present in the SIP header multiple times we can use a slice of its struct to hold the multiple values. These can then be itterated over with the `range` keyword or their size checked with the `len` keyword. An example of this is the Via header field. There can be multiple Via's in a SIP message and they are kept in order as the message is parsed from the first line to the last.

For a Via, The key part of the `SipMsg` struct is shown along with the `sipVia` struct:

```go
type SipMsg struct
    Via []sipVia    // Slice of SIP Vias
    

type sipVia struct {
	Trans  string // Type of Transport udp, tcp, tls, sctp etc
	Host   []byte // Host part
	Port   []byte // Port number
	Branch []byte //
	Rport  []byte //
	Maddr  []byte //
	Ttl    []byte //
	Rcvd   []byte //
	Src    []byte // Full source if needed
}
```

The SDP Attributes field also supports multiple entries.

#### SDP

If SDP is found within a SIP message then it will be parsed too. Media Descriptions, Attributes and Connection Data are all available from the SDP payload. If you wanted to get the media port number from an INVITE with SDP and convert it to an integer, you could use something like:

```go
	port, _ := strconv.Atoi(string(sip.Sdp.MediaDesc.Port))
```

### Reading SIP from other sources

In most real world applications you want to read SIP from an external source. This may be a file, network socket or capture device. If you are wanting to capture with pf_ring then you can checkout my cutdown [pf_ring go library](https://github.com/marv2097/gopfring).












