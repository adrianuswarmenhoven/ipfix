package ipfixmessage

//This file only has generate and documentation

//Retrieve ipfix semantic mappings from various sources and create go code
//Current sources are:
//https://www.ietf.org/assignments/ipfix/ipfix.xml
//https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml
//https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml
//go:generate go run generateipfixmapping.go -i ipfixidmapping_template.gotpl -o ipfixidmapping.go
