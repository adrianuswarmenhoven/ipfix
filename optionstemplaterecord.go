package ipfixmessage

//--- Options Template Record

type OptionsTemplateRecord struct {
	//-- Header
	TemplateID      uint16 //Template ID of this Options Template Record.  This value is greater than 255.
	FieldCount      uint16 //Number of all fields in this Options Template Record, including the Scope Fields.
	ScopeFieldCount uint16 //Number of scope fields in this Options Template Record.  The Scope Fields are normal Fields except that they are interpreted as scope at the Collector.  The Scope Field Count MUST NOT be zero.
	//-- End Header

	FieldSpecifiers []FieldSpecifier
}
