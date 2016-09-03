package ipfixmessage

import (
	"time"
)

//--- IPFIX Message

type IPFIXMessage struct {
	MessageHeader IPFIXMessageHeader

	Sets []*IPFIXSet
}

type IPFIXMessageHeader struct {
	VersionNumber       uint16    //Must be Version
	Length              uint16    //Total length of the IPFIX Message, measured in octets, including Message Header and Set(s).
	ExportTime          time.Time //Time, in seconds (int64), since 0000 UTC Jan 1, 1970, at which the IPFIX Message Header leaves the Exporter.
	SequenceNumber      uint32    //Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent on this PR-SCTP stream from the current Observation Domain by the Exporting Process.
	ObservationDomainID uint32    //A 32-bit identifier of the Observation Domain that is locally unique to the Exporting Process.
}

//--- IPFIX Set

type IPFIXSet struct {
	SetHeader IPFIXSetHeader //Generic Set Header.
	Records   []*IPFIXRecord //Various types of record, depending on the SetID value (or rather, type of IPFIX set)
}

type IPFIXSetHeader struct {
	SetID  uint16 //Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.  All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
	Length uint16 //Total length of the Set, in octets, including the Set Header, all records, and the optional padding.
}

//--------------------------
//--- IPFIX Records
type IPFIXRecord interface {
	Size() uint16
}

//--- Template Record

type IPFIXTemplateRecord struct {
	//-- Header
	TemplateID uint16 //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	FieldCount uint16 //Number of fields in this Template Record.
	//-- End Header

	FieldSpecifiers []IPFIXFieldSpecifier
}

//--- Options Template Record

type IPFIXOptionsTemplateRecord struct {
	//-- Header
	TemplateID      uint16 //Template ID of this Options Template Record.  This value is greater than 255.
	FieldCount      uint16 //Number of all fields in this Options Template Record, including the Scope Fields.
	ScopeFieldCount uint16 //Number of scope fields in this Options Template Record.  The Scope Fields are normal Fields except that they are interpreted as scope at the Collector.  The Scope Field Count MUST NOT be zero.
	//-- End Header

	FieldSpecifiers []IPFIXFieldSpecifier
}

type IPFIXFieldSpecifier struct {
	E                            bool   //Enterprise bit.  This is the first bit of the Field Specifier.  If this bit is zero, the Information Element Identifier identifies an IETF-specified Information Element, and the four-octet Enterprise Number field MUST NOT be present.  If this bit is one, the Information Element identifier identifies an enterprise-specific Information Element, and the Enterprise Number filed MUST be present.
	InformationElementIdentifier uint16 //A numeric value that represents the type of Information Element. Refer to [RFC5102].
	FieldLength                  uint16 //The length of the corresponding encoded Information Element, in octets. The value 65535 is reserved for variable- length Information Elements (see Section 7).
	EnterpriseNumber             uint32 //IANA enterprise number [PEN] of the authority defining the Information Element identifier in this Template Record.
}

//--- Data Record

type IPFIXDataRecord struct {
	FieldValue []interface{} //Note that Field Values do not necessarily have a length of 16 bits. Field Values are encoded according to their data type specified in [RFC5102].
}
