package permissions

import "reflect"

func (p *SystemPermissions) Extend(other *SystemPermissions) {
	if p == nil || other == nil {
		return
	}

	targetVo := reflect.ValueOf(p).Elem()
	sourceVo := reflect.ValueOf(other).Elem()
	// typeR := targetVo.Type()

	for i := 0; i < sourceVo.NumField(); i++ {
		// fieldName := typeR.Field(i).Name
		sourceField := sourceVo.Field(i)
		fieldValue := sourceField.Bool()
		if !fieldValue {
			continue
		}
		targetVo.Set(reflect.Value(sourceField))
	}
}
