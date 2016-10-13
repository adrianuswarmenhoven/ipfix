package ipfixmessage

import (
	"fmt"
	"testing"
)

func TestTemplateListMarker(t *testing.T) {
	fmt.Printf(testMarkerString, "Template List")
}

func TestTemplateListBasic(t *testing.T) {
	tpls := NewActiveTemplateList()
	if tpls == nil {
		t.Errorf(errorPrefixMarker + "Error creating New Active Templates List. Should have gotten empty list, but got nil.")
	}
	dummytpl, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating template record: %#v", err)
	}
	err = tpls.Set(0, dummytpl)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error setting id to 0. Should have gotten error!")
	}
	err = tpls.Set(dummytpl.TemplateID, nil)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error setting nil template. Should have gotten error!")
	}

	err = tpls.Set(dummytpl.TemplateID, dummytpl)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error setting template %d. Got error %#v", dummytpl.TemplateID, err)
	}

	_, err = tpls.Get(0)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error getting template 0. Should have gotten error!")
	}

	dummytpl2, err := tpls.Get(dummytpl.TemplateID)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error getting template %d. Got error %#v", dummytpl.TemplateID, err)
	}
	if dummytpl2 == nil {
		t.Errorf(errorPrefixMarker + "Returned a NIL template.")
	}

}
