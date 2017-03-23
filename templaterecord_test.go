package ipfix

import (
	"fmt"
	"net"
	"testing"
)

const (
	templaterecordTestPrint = false
)

func TestTemplateRecordMarker(t *testing.T) {
	if templaterecordTestPrint {
		fmt.Printf(testMarkerString, "Template Record")
	}
}

func TestTemplateRecordBasic(t *testing.T) {
	_, err := NewTemplateRecord(10)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error creating New Template Record. Should have gotten error, but got nil.")
	}

	tr, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	fsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp1)
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	fsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp2)
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	binarydata, err := tr.MarshalBinary()
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}
	if templaterecordTestPrint {
		fmt.Println(binarydata)
	}
	tr2, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	err = tr2.UnmarshalBinary(binarydata)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}
	if templaterecordTestPrint {
		fmt.Println(tr2)
	}
}

func TestOptionsTemplateRecordBasic(t *testing.T) {
	_, err := NewOptionsTemplateRecord(10)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error creating New Template Record. Should have gotten error, but got nil.")
	}

	tr, err := NewOptionsTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	fsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp1)
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	fsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp2)
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	binarydata, err := tr.MarshalBinary()
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}

	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	ofsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddScopeSpecifier(ofsp1)
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	ofsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddScopeSpecifier(ofsp2)
	if templaterecordTestPrint {
		fmt.Println(tr)
	}
	binarydata, err = tr.MarshalBinary()
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}

	if templaterecordTestPrint {
		fmt.Println(binarydata)
	}
	tr2, err := NewOptionsTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	err = tr2.UnmarshalBinary(binarydata)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}
	if templaterecordTestPrint {
		fmt.Println(tr2)
	}
}

func TestRegisterTemplate(t *testing.T) {
	type substruct struct {
		port  uint32 `ipfix:"substructplaap"`
		value string
	}
	type simplestruct struct {
		sourceip   net.IP    `ipfix:"e:44913,id:14,len:4,someflag,somemalformed:v:a:lue"`
		sourceport uint16    `ipfix:"e:42,id:12,type:unsigned16,len:2,desc:Some field specific to us"`
		basiclist  []string  `ipfix:"e:101,id:100,type:octetarray,len:65535,desc:dnsresolver"`
		ignored    string    //This should be ignored
		subthing   substruct `ipfix:"e:1020,id:101,type:subtemplatelist,subtemplateid:500,desc:A subtemplatelist"`
	}
	_, err := RegisterTemplateRecord(nil, 10, simplestruct{})
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error registering New Template Record. Should have gotten error, but got nil.")
	}
	tmpl, err := RegisterTemplateRecord(nil, 257, simplestruct{})
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error registering New Template Record: %+v", err)
	}
	if tmpl == nil {
		fmt.Println("ok")
	}
}
