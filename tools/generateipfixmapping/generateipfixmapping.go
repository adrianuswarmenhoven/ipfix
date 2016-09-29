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
)

const (
	templatefile = "ipfixmessage.go"
)

type fieldvalueelement struct {
	Name      string
	DataType  string
	ElementID int

	GoFieldValue string //The Field Value as implemented in this package
}

type templatevariables struct {
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
					switch el.DataType {
					case "unsigned8":
						GoRetval = "FieldValueUnsigned8"
					case "unsigned16":
						GoRetval = "FieldValueUnsigned16"
					case "unsigned32":
						GoRetval = "FieldValueUnsigned32"
					case "unsigned64":
						GoRetval = "FieldValueUnsigned64"

					case "signed8":
						GoRetval = "FieldValueSigned8"
					case "signed16":
						GoRetval = "FieldValueSigned16"
					case "signed32":
						GoRetval = "FieldValueSigned32"
					case "signed64":
						GoRetval = "FieldValueSigned64"

					case "float32":
						GoRetval = "FieldValueFloat32"
					case "float64":
						GoRetval = "FieldValueFloat64"

					case "boolean":
						GoRetval = "FieldValueBoolean"

					case "macAddress":
						GoRetval = "FieldValueMacAddress"

					case "octetArray":
						GoRetval = "FieldValueOctetArray"

					case "string":
						GoRetval = "FieldValueString"

					case "dateTimeSeconds":
						GoRetval = "FieldValueDateTimeSeconds"
					case "dateTimeMilliseconds":
						GoRetval = "FieldValueDateTimeMilliseconds"
					case "dateTimeMicroseconds":
						GoRetval = "FieldValueDateTimeMicroseconds"
					case "dateTimeNanoseconds":
						GoRetval = "FieldValueDateTimeNanoseconds"

					case "ipv4Address":
						GoRetval = "FieldValueIPv4Address"
					case "ipv6Address":
						GoRetval = "FieldValueIPv6Address"

					case "basicList":
						GoRetval = "FieldValueBasicList"
						//	case "subTemplateList":
						//		source += fmt.Sprintf("case %d: return &FieldValueSubTemplateList{},nil // %s\n", el.ElementID, el.Name)
						//	case "subTemplateMultiList":
						//		source += fmt.Sprintf("case %d: return &FieldValueSubTemplateMultiList{},nil // %s\n", el.ElementID, el.Name)

					}
					tmpelement := templatedata.Elements[entid][elid]
					tmpelement.GoFieldValue = GoRetval
					templatedata.Elements[entid][elid] = tmpelement
				}
			}
		}
		source := bytes.Buffer{}
		err := sourcetemplate.Execute(&source, templatedata)
		sourceb, err := format.Source(source.Bytes())
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
	sourcefetchers = append(sourcefetchers, FetchSecDorks)
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

// FetchSecDorks gets it's data from https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml
func FetchSecDorks() {
	const (
		SecDorksIPFIXAssignments = "https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml"
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

	elements := IPFixElementsList{}
	sanitizestring := FetchURL(SecDorksIPFIXAssignments)
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
						sources[int(EnterpriseID)] = "SecDorks - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml"
					}
					elementsmap[int(EnterpriseID)][int(ElementID)] = fieldvalueelement{Name: el.Name, DataType: el.DataType, ElementID: int(ElementID)}
				}
			}
		}
	}
}
