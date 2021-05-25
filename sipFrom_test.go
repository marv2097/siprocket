package siprocket

import (
	"reflect"
	"testing"
)

func Test_sipParseFrom1(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = "asdf"
	exp = sipFrom{
		UriType: []byte(""),
		Name:    []byte(""),
		User:    []byte(""),
		Host:    []byte(""),
		Port:    []byte(""),
		Parms: [][]byte{
			[]byte("")},
		Tag: []byte(""),
		Src: []byte(msg),
	}
	if e := parseSipFrom([]byte(msg), &out); e == nil {
		eq := reflect.DeepEqual(out, exp)
		if !eq {
			t.Errorf(`
			UriType %s >> %s
			Name %s >> %s
			User %s >> %s
			Host %s >> %s
			Port %s >> %s
			Tag  %s >> %s
			Src  %s >> %s
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom2(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = "Bob <sip:bob@test.com>;tag=a6c85cf"
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte("Bob"),
		User:    []byte("bob"),
		Host:    []byte("test.com"),
		Port:    []byte(""),
		Parms:   [][]byte{},
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom3(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = `"Board Room"sip:phone_abc_123@test.com;tag=ABCD-123-EFG`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte("Board Room"),
		User:    []byte("phone_abc_123"),
		Host:    []byte("test.com"),
		Port:    []byte(""),
		Parms:   [][]byte{},
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom4(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = ` <sip:10.0.0.1:5060;transport=udp;lr>;tag=sip+654321`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(""),
		User:    []byte(""),
		Host:    []byte("10.0.0.1"),
		Port:    []byte("5060"),
		Parms: [][]byte{
			[]byte("transport=udp"),
			[]byte("lr"),
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom5(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = `sip:10.0.0.1:5060`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(""),
		User:    []byte(""),
		Host:    []byte("10.0.0.1"),
		Port:    []byte("5060"),
		Parms:   [][]byte{},
		Tag:     []byte(""),
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom6(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = `sip:unlimitedsystem.co.uk;tag=12345-6789-`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(""),
		User:    []byte(""),
		Host:    []byte("unlimitedsystem.co.uk"),
		Port:    []byte(""),
		Parms:   [][]byte{},
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom7(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = `sip:test.system@mydomain.co.uk`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(""),
		User:    []byte("test.system"),
		Host:    []byte("mydomain.co.uk"),
		Port:    []byte(""),
		Parms:   [][]byte{},
		Tag:     []byte(""),
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom8(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = ` <sip:+440800800150@10.0.0.1;user=phone>;tag=1234-4567`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(""),
		User:    []byte("+440800800150"),
		Host:    []byte("10.0.0.1"),
		Port:    []byte(""),
		Parms: [][]byte{
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}

func Test_sipParseFrom9(t *testing.T) {

	var out, exp sipFrom
	var msg string

	msg = ` <sip:9876543521;phone-context=+44@10.0.0.1;user=phone>;tag=sip+6+a100+g333`
	exp = sipFrom{
		UriType: []byte("sip"),
		Name:    []byte(""),
		User:    []byte("987654352"),
		Host:    []byte(""),
		Port:    []byte(""),
		Parms: [][]byte{
			[]byte("phone-context=+44@10.0.0.1"),
			[]byte("user=phone"),
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
			Parms %v  >>  %v
			`, out.UriType, exp.UriType,
				out.Name, exp.Name,
				out.User, exp.User,
				out.Host, exp.Host,
				out.Port, exp.Port,
				out.Tag, exp.Tag,
				out.Src, exp.Src,
				len(out.Parms), len(exp.Parms),
			)
			for k, _ := range exp.Parms {
				t.Errorf(`param[%v] '%s' >> '%s'`, k, out.Parms[k], exp.Parms[len(exp.Parms)-k-1])
			}
		}
	} else {
		t.Errorf("example %s generated the error %s\n", msg, e)
	}
}
