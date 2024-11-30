package convertors

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"grpc/pkg/proto_gen/grpc"
	"reflect"
)

type RequestConvertorInterface interface {
	Convert(src interface{}, dest interface{}) (interface{}, error)
}

type RequestConvertor struct {
	Request    interface{}
	DestStruct interface{}
}

func (c RequestConvertor) Convert(src interface{}, dest interface{}) (interface{}, error) {
	s := reflect.Indirect(reflect.ValueOf(src))
	d := reflect.Indirect(reflect.ValueOf(dest))

	if s.Kind() != reflect.Struct || d.Kind() != reflect.Struct {
		return nil, errors.New("request convert. request is not a structure")
	}

	for i := 0; i < d.NumField(); i++ {
		//fieldName := s.Type().Field(i).Name
		//destType := s.Type().Field(i).Tag.Get("convertType")
		field := d.Type().Field(i)
		requestField := s.FieldByName(field.Name)
		convertable := ""

		if v, ok := s.Type().FieldByName(field.Name); ok {
			convertable = v.Tag.Get("convertable")
		}

		if convertable != "" && convertable != "false" {
			requestFieldKind := requestField.Kind()
			requestFieldType := requestField.Type()
			fieldVal := d.Field(i)
			if fieldVal.CanSet() && field.Type.Kind() == requestFieldKind && requestFieldType == field.Type {
				fieldVal.Set(requestField)
			} else if fieldVal.CanSet() {
				val := requestField.Elem()

				switch field.Name {
				case "Sort":
					fieldVal.SetString(c.convertSort(val))
				case "Page":
					fieldVal.Set(reflect.ValueOf(c.convertPage(val)))
				case "Status":
					fieldVal.Set(reflect.ValueOf(c.convertStatus(val)))
				}

				switch field.Type.String() {
				case "uuid.UUID":
					fieldVal.Set(reflect.ValueOf(c.convertUuid(val)))
				}
			}

			fmt.Println("converted request: ", dest)
		}
	}

	return dest, nil
}

func (c RequestConvertor) convertSort(v reflect.Value) string {
	if v.Type().String() == "grpc.Sort" {
		convertedVal := v.MethodByName("String").Call(nil)
		if len(convertedVal) > 0 {
			return convertedVal[0].Interface().(string)
		}
	}

	return "ASC"
}

func (c RequestConvertor) convertPage(v reflect.Value) int32 {
	if v.Type().String() == "grpc.Page" && v.CanInt() && v.Int() > 0 {
		return int32(v.Int())
	}

	return 1
}

func (c RequestConvertor) convertStatus(v reflect.Value) int32 {
	if v.Type().String() == "grpc.Page" && v.CanInt() && v.Int() > 0 {
		return int32(v.Int())
	}

	return 1
}

func (c RequestConvertor) convertUuid(v reflect.Value) uuid.UUID {
	if v.Type() == reflect.TypeOf(grpc.UUID{}) {
		id := v.FieldByName("Id").String()
		uid, err := uuid.Parse(id)
		if err == nil {
			return uid
		}

	}

	return uuid.Nil
}
