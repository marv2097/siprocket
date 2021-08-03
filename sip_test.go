package siprocket

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_sipParse_Nonsense(t *testing.T) {

	var out, exp SipMsg

	msg := `asdf`
	exp = SipMsg{
		Req: sipReq{
			Method:     []byte(nil),
			UriType:    "",
			StatusCode: []byte(nil),
			StatusDesc: []byte(nil),
			User:       []byte(nil),
			Host:       []byte(nil),
			Port:       []byte(nil),
			UserType:   []byte(nil),
			Src:        []byte("asdf"),
		},
		From: sipFrom{
			UriType: []byte(nil),
			Name:    []byte(nil),
			User:    []byte(nil),
			Host:    []byte(nil),
			Port:    []byte(nil),
			Params:  [][]byte(nil),
			Tag:     []byte(nil),
			Src:     []byte(nil),
		},
		To: sipTo{
			UriType: []byte(nil),
			Name:    []byte(nil),
			User:    []byte(nil),
			Host:    []byte(nil),
			Port:    []byte(nil),
			Params:  [][]byte(nil),
			Tag:     []byte(nil),
			Src:     []byte(nil),
		},
		Contact: sipContact{
			UriType: "",
			Name:    []byte(nil),
			User:    []byte(nil),
			Host:    []byte(nil),
			Port:    []byte(nil),
			Tran:    []byte(nil),
			Qval:    []byte(nil),
			Expires: []byte(nil),
			Src:     []byte(nil),
		},
		Via: []sipVia{},
		Cseq: sipCseq{
			Id:     []byte(nil),
			Method: []byte(nil),
			Src:    []byte(nil),
		},
		Ua: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		Exp: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		MaxFwd: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		CallId: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		ContType: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		ContLen: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		XGammaIP: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},

		Sdp: SdpMsg{
			MediaDesc: sdpMediaDesc{
				MediaType: []byte(nil),
				Port:      []byte(nil),
				Proto:     []byte(nil),
				Fmt:       []byte(nil),
				Src:       []byte(nil),
			},
			Attrib: []sdpAttrib{},
			ConnData: sdpConnData{
				AddrType: []byte(nil),
				ConnAddr: []byte(nil),
				Src:      []byte(nil),
			},
		},
	}
	out = Parse([]byte(msg))
	eq := reflect.DeepEqual(out, exp)
	if !eq {
		t.Errorf("\n%#v \n %#v", exp, out)
	}

}

func Test_sipParse_invite(t *testing.T) {

	var out, exp SipMsg

	msg := `INVITE sip:123456789@testcompany.com SIP/2.0
Via: SIP/2.0/WSS testcompany.com;branch=z0GMslasdf
Max-Forwards: 69
To: <sip:123456789@testcompany.com>
From: <sip:PersonA_PC_123456789@testcompany.com>;tag=ujpedsvksh
Call-ID: kasdf023l4qklaansdf02
CSeq: 8918 INVITE
X-gamma-public-ip: 127.0.0.1
Contact: <sip:PersonA_PC_123456789@testcompany.com;ob>
Content-Type: application/sdp
Allow: INVITE,ACK,CANCEL,BYE,UPDATE,MESSAGE,OPTIONS,REFER,INFO,NOTIFY
Supported: ice,replaces,outbound
User-Agent: softphone-desktop
Content-Length: 1245
	
m=audio 51268 RTP/AVP 111 9 8 101
c=IN IP4 127.0.0.1
a=rtpmap:111 opus/48000/2
a=rtpmap:9 G722/8000`
	exp = SipMsg{
		Req: sipReq{
			Method:     []byte("INVITE"),
			UriType:    "sip",
			StatusCode: []byte(nil),
			StatusDesc: []byte(nil),
			User:       []byte("123456789"),
			Host:       []byte("testcompany.com"),
			Port:       []byte(nil),
			UserType:   []byte(nil),
			Src:        []byte("INVITE sip:123456789@testcompany.com SIP/2.0"),
		},
		From: sipFrom{
			UriType: []byte("sip"),
			Name:    []byte(nil),
			User:    []byte("PersonA_PC_123456789"),
			Host:    []byte("testcompany.com"),
			Port:    []byte(nil),
			Params:  [][]byte(nil),
			Tag:     []byte("ujpedsvksh"),
			Src:     []byte("<sip:PersonA_PC_123456789@testcompany.com>;tag=ujpedsvksh"),
		},
		To: sipTo{
			UriType: []byte("sip"),
			Name:    []byte(nil),
			User:    []byte("123456789"),
			Host:    []byte("testcompany.com"),
			Port:    []byte(nil),
			Params:  [][]byte(nil),
			Tag:     []byte(nil),
			Src:     []byte("<sip:123456789@testcompany.com>"),
		},
		Contact: sipContact{
			UriType: "sip",
			Name:    []byte(nil),
			User:    []byte("PersonA_PC_123456789"),
			Host:    []byte("testcompany.com"),
			Port:    []byte(nil),
			Tran:    []byte(nil),
			Qval:    []byte(nil),
			Expires: []byte(nil),
			Src:     []byte("<sip:PersonA_PC_123456789@testcompany.com;ob>"),
		},
		Via: []sipVia{
			{
				Trans:  "wss",
				Host:   []byte("testcompany.com"),
				Port:   []byte(nil),
				Branch: []byte("z0GMslasdf"),
				Rport:  []byte(nil),
				Maddr:  []byte(nil),
				Ttl:    []byte(nil),
				Rcvd:   []byte(nil),
				Src:    []byte("SIP/2.0/WSS testcompany.com;branch=z0GMslasdf"),
			},
		},
		Cseq: sipCseq{
			Id:     []byte("8918"),
			Method: []byte("INVITE"),
			Src:    []byte("8918 INVITE"),
		},
		Ua: sipVal{
			Value: []byte("softphone-desktop"),
			Src:   []byte(nil),
		},
		Exp: sipVal{
			Value: []byte(nil),
			Src:   []byte(nil),
		},
		MaxFwd: sipVal{
			Value: []byte("69"),
			Src:   []byte(nil),
		},
		CallId: sipVal{
			Value: []byte("kasdf023l4qklaansdf02"),
			Src:   []byte(nil),
		},
		ContType: sipVal{
			Value: []byte("application/sdp"),
			Src:   []byte(nil),
		},
		ContLen: sipVal{
			Value: []byte("1245"),
			Src:   []byte(nil),
		},
		XGammaIP: sipVal{
			Value: []byte("127.0.0.1"),
			Src:   []byte(nil),
		},
		Sdp: SdpMsg{
			MediaDesc: sdpMediaDesc{
				MediaType: []byte("audio"),
				Port:      []byte("51268"),
				Proto:     []byte("RTP/AVP"),
				Fmt:       []byte("111 9 8 101"),
				Src:       []byte("audio 51268 RTP/AVP 111 9 8 101"),
			},
			Attrib: []sdpAttrib{
				{
					Cat: []byte("rtpmap"),
					Val: []byte("111 opus/48000/2"),
					Src: []byte("rtpmap:111 opus/48000/2"),
				},
				{
					Cat: []byte("rtpmap"),
					Val: []byte("9 G722/8000"),
					Src: []byte("rtpmap:9 G722/8000"),
				},
			},
			ConnData: sdpConnData{
				AddrType: []byte("IP4"),
				ConnAddr: []byte("127.0.0.1"),
				Src:      []byte("IN IP4 127.0.0.1"),
			},
		},
	}
	out = Parse([]byte(msg))
	eq := reflect.DeepEqual(out, exp)
	if !eq {
		exp, _ := json.Marshal(exp)
		out, _ := json.Marshal(out)
		t.Errorf("\n%s \n %s", exp, out)
	}
}
