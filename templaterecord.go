package ipfixmessage

//--- Template Record

type IPFIXTemplateRecord struct {
	//-- Header
	TemplateID uint16 //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	FieldCount uint16 //Number of fields in this Template Record.
	//-- End Header

	FieldSpecifiers []IPFIXFieldSpecifier
}
