package pdo

import "fmt"

// String - return a string representation of the DataObject
func (pdo *DataObject) String() string {
	switch pdo.State.(type) {
	case string:
		return fmt.Sprintf("%s", pdo.State.(string))
	case int:
		return fmt.Sprintf("%d", pdo.State.(int))
	case int8:
		return fmt.Sprintf("%d", pdo.State.(int8))
	case int16:
		return fmt.Sprintf("%d", pdo.State.(int16))
	case int32:
		return fmt.Sprintf("%d", pdo.State.(int32))
	case int64:
		return fmt.Sprintf("%d", pdo.State.(int64))
	case uint:
		return fmt.Sprintf("%d", pdo.State.(uint))
	case uint8:
		return fmt.Sprintf("%d", pdo.State.(uint8))
	case uint16:
		return fmt.Sprintf("%d", pdo.State.(uint16))
	case uint32:
		return fmt.Sprintf("%d", pdo.State.(uint32))
	case uint64:
		return fmt.Sprintf("%d", pdo.State.(uint64))
	case bool:
		return fmt.Sprintf("%t", pdo.State.(bool))
	case float32:
		return fmt.Sprintf("%f", pdo.State.(float32))
	case float64:
		return fmt.Sprintf("%f", pdo.State.(float64))
	case Array:
		return pdo.State.(Array).String()
	case Enum:
		return pdo.State.(Enum).String()
	case PdoParameterObject:
		return pdo.State.(PdoParameterObject).String()
	default:
		panic(fmt.Sprintf("unhandled default case in DataObject.String(): '%c'", pdo.State))
	}
}
