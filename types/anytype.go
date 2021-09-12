package types

import "encoding/base64"

// Smart type for converting
type AnyType struct {
	Valid  bool
	String string
}

//MarshalJSON method is called by json.Marshal,
//whenever it is of type NullString
func (x *AnyType) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}

	return []byte(x.String), nil
}

func (x *AnyType) UnmarshalJSON(data []byte) error {
	x.String = "\"" + base64.StdEncoding.EncodeToString(data) + "\""
	x.Valid = true

	return nil
}
