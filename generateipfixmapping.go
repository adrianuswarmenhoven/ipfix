// +build ignore

// This program generates the output for the byterootgeneratedtriesearch.go file to prevent accidents it needs to be redirected.
//It can be invoked by running go generate
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	SourceMarker = "//***GENERATEMARKER***"
)

type fieldvalueelement struct {
	Name      string
	DataType  string
	ElementID int

	GoFieldValue  string //The Field Value as implemented in this package
	GoFieldLength string //The length of the Field
}

type templatevariables struct {
	TimeStamp       time.Time
	TemplateFile    string
	Elements        map[int]map[int]fieldvalueelement
	EnterpriseOrder sort.IntSlice
	ElementsOrder   map[int]sort.IntSlice
	Sources         map[int]string
}

var (
	sources        = make(map[int]string)
	elementsmap    = make(map[int]map[int]fieldvalueelement)
	sourcefetchers = make([]func(), 0, 0)
)

func main() {
	input := flag.String("i", "", "Input template")
	output := flag.String("o", "", "Output")
	flag.Usage = func() {
		fmt.Println(`Usage: generateipfixmapping -i <template.go> [-o <output.go>]

-i  Input template
-o  Optional output filename
`)
	}
	flag.Parse()
	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	for _, sourcefetcher := range sourcefetchers {
		sourcefetcher()
	}
	sourcetemplate, err := template.ParseFiles(*input)
	if err == nil {
		templatedata := templatevariables{
			TemplateFile:  *input,
			TimeStamp:     time.Now(),
			Elements:      elementsmap,
			ElementsOrder: make(map[int]sort.IntSlice),
			Sources:       sources,
		}
		sortedentid := make(sort.IntSlice, 0, 0)
		for entid := range elementsmap {
			sortedentid = append(sortedentid, entid)
		}
		sort.Sort(sortedentid)
		templatedata.EnterpriseOrder = sortedentid
		for _, entid := range sortedentid {
			fields := elementsmap[entid]
			sortedfieldid := make(sort.IntSlice, 0, 0)
			for fieldid := range fields {
				sortedfieldid = append(sortedfieldid, fieldid)
			}
			sort.Sort(sortedfieldid)
			templatedata.ElementsOrder[entid] = sortedfieldid
			for _, elid := range sortedfieldid {
				el := fields[elid]
				if strings.TrimSpace(el.Name) != "" &&
					strings.TrimSpace(el.DataType) != "" &&
					el.ElementID != 0 {
					GoRetval := "FieldValueOctetArray" //If we simply do not know, we can always pass it on as a bunch of bytes
					GoFieldLen := "65535"
					switch el.DataType {
					case "unsigned8":
						GoRetval = "FieldValueUnsigned8"
						GoFieldLen = "1"
					case "unsigned16":
						GoRetval = "FieldValueUnsigned16"
						GoFieldLen = "2"
					case "unsigned32":
						GoRetval = "FieldValueUnsigned32"
						GoFieldLen = "4"
					case "unsigned64":
						GoRetval = "FieldValueUnsigned64"
						GoFieldLen = "8"

					case "signed8":
						GoRetval = "FieldValueSigned8"
						GoFieldLen = "1"
					case "signed16":
						GoRetval = "FieldValueSigned16"
						GoFieldLen = "2"
					case "signed32":
						GoRetval = "FieldValueSigned32"
						GoFieldLen = "4"
					case "signed64":
						GoRetval = "FieldValueSigned64"
						GoFieldLen = "8"

					case "float32":
						GoRetval = "FieldValueFloat32"
						GoFieldLen = "4"
					case "float64":
						GoRetval = "FieldValueFloat64"
						GoFieldLen = "8"

					case "boolean":
						GoRetval = "FieldValueBoolean"
						GoFieldLen = "1"

					case "macAddress":
						GoRetval = "FieldValueMacAddress"
						GoFieldLen = "6"

					case "octetArray":
						GoRetval = "FieldValueOctetArray"
						GoFieldLen = "65535"

					case "string":
						GoRetval = "FieldValueString"
						GoFieldLen = "65535"

					case "dateTimeSeconds":
						GoRetval = "FieldValueDateTimeSeconds"
						GoFieldLen = "4"
					case "dateTimeMilliseconds":
						GoRetval = "FieldValueDateTimeMilliseconds"
						GoFieldLen = "4"
					case "dateTimeMicroseconds":
						GoRetval = "FieldValueDateTimeMicroseconds"
						GoFieldLen = "4"
					case "dateTimeNanoseconds":
						GoRetval = "FieldValueDateTimeNanoseconds"
						GoFieldLen = "4"

					case "ipv4Address":
						GoRetval = "FieldValueIPv4Address"
						GoFieldLen = "4"
					case "ipv6Address":
						GoRetval = "FieldValueIPv6Address"
						GoFieldLen = "16"

					case "basicList":
						GoRetval = "FieldValueBasicList"
						GoFieldLen = "65535"

					case "subTemplateList":
						GoRetval = "FieldValueSubTemplateList"
						GoFieldLen = "65535"
					case "subTemplateMultiList":
						GoRetval = "FieldValueSubTemplateMultiList"
						GoFieldLen = "65535"
					}
					tmpelement := templatedata.Elements[entid][elid]
					tmpelement.GoFieldValue = GoRetval
					tmpelement.GoFieldLength = GoFieldLen
					templatedata.Elements[entid][elid] = tmpelement
				}
			}
		}
		source := bytes.Buffer{}
		err := sourcetemplate.Execute(&source, templatedata)
		sourceb, err := format.Source(source.Bytes())
		marker := bytes.Index(sourceb, []byte(SourceMarker))
		if marker > -1 {
			sourceb = sourceb[marker+len(SourceMarker)+1:]
		}
		if err != nil {
			fmt.Println(err, string(source.Bytes()))
		} else {
			if *output == "" {
				fmt.Println(string(sourceb))
			} else {
				ioutil.WriteFile(*output, sourceb, os.ModePerm)
			}
		}
	}
}

func init() {
	sourcefetchers = append(sourcefetchers, FetchIANA)
	sourcefetchers = append(sourcefetchers, FetchIPFIXColStyle)
}

func FetchURL(fetchurl string) string {
	data := bytes.NewBuffer([]byte{})
	response, err := http.Get(fetchurl)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(data, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	return string(data.Bytes())
}

// FetchIANA gets it's data from https://www.ietf.org/assignments/ipfix/ipfix.xml
func FetchIANA() {
	const (
		IANAIPFIXAssignments = "https://www.ietf.org/assignments/ipfix/ipfix.xml"
	)
	type Element struct {
		XMLName xml.Name `xml:"record"`

		Name      string `xml:"name"`
		DataType  string `xml:"dataType"`
		ElementID string `xml:"elementId"`
	}

	type IPFixElementsList struct {
		XMLName xml.Name `xml:"registry"`

		Element []Element `xml:"record"`
	}

	type MainRegistry struct {
		XMLName xml.Name `xml:"registry"`

		Registry IPFixElementsList `xml:"registry"`
	}

	elementsmap[0] = make(map[int]fieldvalueelement)
	sources[0] = "IANA - https://www.ietf.org/assignments/ipfix/ipfix.xml"
	registry := MainRegistry{}
	sanitizestring := FetchURL(IANAIPFIXAssignments)
	sanitizestring = strings.Map(func(check rune) rune {
		switch check {
		case '\n', '\r':
			return -1
		}
		return check
	}, sanitizestring)
	err := xml.Unmarshal([]byte(sanitizestring), &registry)
	if err != nil {
		fmt.Printf("Malformed XM: %#v\n", err)
		fmt.Println(err)
	} else {
		for _, el := range registry.Registry.Element {
			ElementID, err := strconv.ParseInt(el.ElementID, 10, 32)
			if err == nil && strings.TrimSpace(el.Name) != "" &&
				strings.TrimSpace(el.DataType) != "" &&
				ElementID != 0 {
				elementsmap[0][int(ElementID)] = fieldvalueelement{Name: el.Name, DataType: el.DataType, ElementID: int(ElementID)}
			}
		}
	}
}

// FetchIPFIXColStyle gets it's data from https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml and others
func FetchIPFIXColStyle() {
	var (
		IPFIXColStyleXML = []string{
			0: "https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml",
			1: "https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml",
		}
	)
	type Element struct {
		XMLName xml.Name `xml:"element"`

		Name         string `xml:"name"`
		DataType     string `xml:"dataType"`
		ElementID    string `xml:"id"`
		EnterpriseID string `xml:"enterprise"`
	}

	type IPFixElementsList struct {
		XMLName xml.Name `xml:"ipfix-elements"`

		Element []Element `xml:"element"`
	}
	for _, fetchurl := range IPFIXColStyleXML {
		elements := IPFixElementsList{}
		sanitizestring := FetchURL(fetchurl)
		sanitizestring = strings.Map(func(check rune) rune {
			switch check {
			case '\n', '\r':
				return -1
			}
			return check
		}, sanitizestring)
		err := xml.Unmarshal([]byte(sanitizestring), &elements)
		if err != nil {
			fmt.Printf("Malformed XM: %#v\n", err)
			fmt.Println(err)
		} else {
			for _, el := range elements.Element {
				ElementID, err := strconv.ParseInt(el.ElementID, 10, 32)
				if err == nil {
					EnterpriseID, err := strconv.ParseInt(el.EnterpriseID, 10, 32)
					if err == nil && strings.TrimSpace(el.Name) != "" &&
						strings.TrimSpace(el.DataType) != "" &&
						ElementID != 0 && EnterpriseID != 0 {
						if _, exists := elementsmap[int(EnterpriseID)]; !exists {
							elementsmap[int(EnterpriseID)] = make(map[int]fieldvalueelement)
							sources[int(EnterpriseID)] = "IPFIXColStyle - " + fetchurl
						}
						elementsmap[int(EnterpriseID)][int(ElementID)] = fieldvalueelement{Name: el.Name, DataType: el.DataType, ElementID: int(ElementID)}
					}
				}
			}
		}
	}
}
