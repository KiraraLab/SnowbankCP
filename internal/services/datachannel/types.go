package datachannel

type DataType string

var DataTypes = struct {
	String    DataType
	Integer   DataType
	Float     DataType
	Boolean   DataType
	Timestamp DataType
	JSON      DataType
	Blob      DataType
}{
	String:    "string",
	Integer:   "integer",
	Float:     "float",
	Boolean:   "boolean",
	Timestamp: "timestamp",
	JSON:      "json",
	Blob:      "blob",
}
