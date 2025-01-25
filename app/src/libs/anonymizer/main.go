package anonymizer

func anonymizeField(field *string, method AnonimizationType, mask string) {
	// Anonymize the field based on the method
	switch method {
	case TypeReplace:
		replaceField(field, mask)
	case TypeMask:
		maskField(field, mask)
	case TypeHide:
		hideField(field)
	case TypePartialHide:
		partiallyHideField(field)
	}
}

// AnonymizeRecord anonymizes the fields of a given record based on the
// attributes provided. The record must be a pointer to a struct.
//
// Example:
//
// AnonymizeRecord(&record, []Attribute{})
func AnonymizeRecord(record *map[string]interface{}, attributes []Attribute, uniqueMap *UniqueAttributes) error {
	// Loop through the attributes and anonymize the fields
	for _, attribute := range attributes {
		// Gets the reflect value of the field
		interfaceField, ok := (*record)[attribute.Name]
		if !ok {
			return ErrFieldNotFound(attribute.Name)
		}

		if interfaceField == nil {
			(*record)[attribute.Name] = nil
			continue
		}

		field, ok := interfaceField.(string)
		if !ok {
			return ErrFieldNotString(attribute.Name)
		}

		anonymizeField(&field, attribute.Method, attribute.Mask)

		if uniqueMap != nil && attribute.Unique {
			isValid := false
			for !isValid {
				_, alreadyExists := (*uniqueMap)[attribute.Name][field]
				if alreadyExists {
					anonymizeField(&field, attribute.Method, attribute.Mask)
				} else {
					isValid = true

					if (*uniqueMap)[attribute.Name] == nil {
						(*uniqueMap)[attribute.Name] = make(UniqueValues)
					}

					(*uniqueMap)[attribute.Name][field] = true
				}
			}
		}

		// Update the field value
		(*record)[attribute.Name] = field
	}

	return nil
}
