package siprocket

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_sipParseRequestLine_Nonsense(t *testing.T) {

	var out sipReq

	msg := "asdf"

	if e := parseSipReq([]byte(msg), &out); e == nil {
		t.Errorf("failed to generated an error")
	}
}

func Test_sipParse_RequestLine_REGISTER(t *testing.T) {

	var out, exp sipReq

	msg := "REGISTER sip:0800800140@test.com:5060 SIP"
	exp = sipReq{
		Method: 	[]byte("REGISTER"),
		UriType:    []byte("sip"),
		StatusCode: []byte(nil),
		StatusDesc: []byte(nil),
		User:    	[]byte("0800800140"),
		Host:  		[]byte("test.com"),
		Port:     	[]byte("5060"),
		UserType:   []byte(nil),
		Src: 		[]byte(msg),
	}
	if e := parseSipReq([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			Method '%s' >> '%s'
			UriType '%s' >> '%s'
			StatusCode '%s' >> '%s'
			StatusDesc '%s' >> '%s'
			User '%s' >> '%s'
			Host  '%s' >> '%s'
			Port  '%s' >> '%s'
			UserType '%v'  >>  '%v'
			Src '%v'
			    '%v'
			`, out.Method, exp.Method,
				out.UriType, exp.UriType,
				out.StatusCode, exp.StatusCode,
				out.StatusDesc, exp.StatusDesc,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.UserType, exp.UserType,
				out.Src, exp.Src,
			)
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_RequestLine_INVITE(t *testing.T) {

	var out, exp sipReq

	msg := "INVITE sips:8508000123456;phone-context=+44@10.0.0.1;user=phone SIP/2.0"
	exp = sipReq{
		Method: 	[]byte("INVITE"),
		UriType:    []byte("sips"),
		StatusCode: []byte(nil),
		StatusDesc: []byte(nil),
		User:    	[]byte("8508000123456"),
		Host:  		[]byte("10.0.0.1"),
		Port:     	[]byte(nil),
		UserType:   []byte("phone"),
		Src: 		[]byte(msg),
	}
	if e := parseSipReq([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			Method '%s' >> '%s'
			UriType '%s' >> '%s'
			StatusCode '%s' >> '%s'
			StatusDesc '%s' >> '%s'
			User '%s' >> '%s'
			Host  '%s' >> '%s'
			Port  '%s' >> '%s'
			UserType '%v'  >>  '%v'
			Src '%v'
			    '%v'
			`, out.Method, exp.Method,
				out.UriType, exp.UriType,
				out.StatusCode, exp.StatusCode,
				out.StatusDesc, exp.StatusDesc,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.UserType, exp.UserType,
				out.Src, exp.Src,
			)
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}