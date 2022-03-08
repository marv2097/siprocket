package siprocket

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_sipParseFrom_Nonsense(t *testing.T) {

	var out sipFrom

	msg := "asdf"

	if e := parseSipFrom([]byte(msg), &out); e == nil {
		t.Errorf("failed to generated an error")
	}
}

func Test_sipParse_From_2(t *testing.T) {

	var out, exp sipFrom

	msg := "Bob <sip:bob@test.com>;tag=a6c85cf"
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte("Bob"),
		User:    []byte("bob"),
		Host:    []byte("test.com"),
		Port:    []byte(nil),
		Params:  [][]byte(nil),
		Tag:     []byte("a6c85cf"),
		Src:     []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_3(t *testing.T) {

	var out, exp sipFrom

	msg := `"Board Room" <sip:phone_abc_123@test.com>;tag=ABCD-123-EFG`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte("Board Room"),
		User:    []byte("phone_abc_123"),
		Host:    []byte("test.com"),
		Port:    []byte(nil),
		Params:  [][]byte(nil),
		Tag:     []byte("ABCD-123-EFG"),
		Src:     []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_4(t *testing.T) {

	var out, exp sipFrom

	msg := ` <sip:10.0.0.1:5060;transport=udp;lr>;tag=sip+654321`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(nil),
		User:    []byte(nil),
		Host:    []byte("10.0.0.1"),
		Port:    []byte("5060"),
		Params: [][]byte{
			[]byte("lr"),
			[]byte("transport=udp"),
		},
		Tag: []byte("sip+654321"),
		Src: []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_5(t *testing.T) {

	var out, exp sipFrom

	msg := `sip:10.0.0.1:5060`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(nil),
		User:    []byte(nil),
		Host:    []byte("10.0.0.1"),
		Port:    []byte("5060"),
		Params:  [][]byte(nil),
		Tag:     []byte(nil),
		Src:     []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_6(t *testing.T) {

	var out, exp sipFrom

	msg := `sip:unlimitedsystem.co.uk;tag=12345-6789-`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(nil),
		User:    []byte(nil),
		Host:    []byte("unlimitedsystem.co.uk"),
		Port:    []byte(nil),
		Params:  [][]byte(nil),
		Tag:     []byte("12345-6789-"),
		Src:     []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_7(t *testing.T) {

	var out, exp sipFrom

	msg := `sip:test.system@mydomain.co.uk`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(nil),
		User:    []byte("test.system"),
		Host:    []byte("mydomain.co.uk"),
		Port:    []byte(nil),
		Params:  [][]byte(nil),
		Tag:     []byte(nil),
		Src:     []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_8(t *testing.T) {

	var out, exp sipFrom

	msg := ` <sip:+440800800150@10.0.0.1;user=phone>;tag=1234-4567`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(nil),
		User:    []byte("+440800800150"),
		Host:    []byte("10.0.0.1"),
		Port:    []byte(nil),
		Params: [][]byte{
			[]byte("user=phone"),
		},
		Tag: []byte("1234-4567"),
		Src: []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParse_From_9(t *testing.T) {

	var out, exp sipFrom

	msg := `<sip:9876543521;phone-context=+44@10.0.0.1;user=phone>;tag=sip+6+a100+g333`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(nil),
		User:    []byte(nil),
		Host:    []byte("9876543521"),
		Port:    []byte(nil),
		Params: [][]byte{
			[]byte("user=phone"),
			[]byte("phone-context=+44@10.0.0.1"),
		},
		Tag: []byte("sip+6+a100+g333"),
		Src: []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType '%s' >> '%s'
			Name '%s' >> '%s'
			User '%s' >> '%s'
			Host '%s' >> '%s'
			Port '%s' >> '%s'
			Tag  '%s' >> '%s'
			Src  '%s' >> '%s'
			Params %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Params), len(exp.Params),
			)
			for k, _ := range exp.Params {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Params[k], exp.Params[len(exp.Params)-k-1])
			}
			exp, _ := json.Marshal(exp)
			out, _ := json.Marshal(out)
			t.Errorf("\n%s \n %s", exp, out)
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}
