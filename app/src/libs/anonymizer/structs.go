package anonymizer

type AnonimizationType string

const (
	TypeReplace     AnonimizationType = "replace"
	TypeMask        AnonimizationType = "mask"
	TypeHide        AnonimizationType = "hide"
	TypePartialHide AnonimizationType = "hide_partial"
)

// Attribute is a struct that represents an attribute to be anonymized.
// It contains the name of the field, the anonymization method and the mask
// to be used in the anonymization process. If the anonymization method is
// "replace", the mask will be used as the replacement value. If the method
// is "mask", the mask will be used to generate a random value based on the
// mask pattern. If the method is "hide", the field value will be replaced
// with a sequence of X's. If the method is "hide_partial", all but the first
// 5 characters of the field value will be replaced with X's.
type Attribute struct {
	Name   string
	Mask   string
	Method AnonimizationType
	Unique bool
}

type UniqueValues map[string]interface{}
type UniqueAttributes map[string]UniqueValues
