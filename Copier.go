package copier

import (
	"encoding/json"
	"reflect"
)

// Config contain configuartion flags for Copier
type Config struct {
	OmitEmpty  bool
	OmitByJSON bool
}

// Copier Contains Details for Copy Operations.
type Copier struct {
	Source     interface{}
	SourceType reflect.Type
	SourceElem reflect.Value
	SourceKind reflect.Kind
	SourceAddr reflect.Type

	Destination     interface{}
	DestinationType reflect.Type
	DestinationElem reflect.Value
	DestinationKind reflect.Kind
	DestinationAddr reflect.Type

	Configuration Config
}

// NewCopier Creates a Copier with Configuration.
func NewCopier(config Config) *Copier {
	copier := Copier{
		Configuration: config,
	}

	return &copier
}

func (copier *Copier) From(v interface{}) *Copier {
	copier.Source = v
	copier.SourceType = reflect.TypeOf(v)
	copier.SourceElem = reflect.Indirect(reflect.ValueOf(v))
	copier.SourceKind = copier.SourceElem.Kind()
	copier.SourceAddr = reflect.PtrTo(copier.SourceType)

	return copier
}

func (copier *Copier) To(v interface{}) *Copier {
	copier.Destination = v
	copier.DestinationType = reflect.TypeOf(v)
	copier.DestinationElem = reflect.Indirect(reflect.ValueOf(v))
	copier.DestinationKind = copier.SourceElem.Kind()
	copier.DestinationAddr = reflect.PtrTo(copier.SourceType)

	if copier.SourceKind == reflect.Struct {
		if copier.DestinationKind == reflect.Map {
			copier.copyStructToMap()
		}

		if copier.DestinationKind == reflect.Struct {
			copier.copyStuctToStruct()
		}
	}

	if copier.SourceKind == reflect.Map {
		if copier.DestinationKind == reflect.Map {
			copier.copyMapToMap()
		}

		if copier.DestinationKind == reflect.Struct {
			copier.copyMapToStruct()
		}
	}

	return copier
}

func (copier *Copier) copyStructToMap() error {

	if copier.Configuration.OmitByJSON {
		bytes, err := json.Marshal(copier.SourceAddr)
		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, copier.DestinationAddr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (copier *Copier) copyStuctToStruct() {

}

func (copier *Copier) copyMapToStruct() {

}

func (copier *Copier) copyMapToMap() {

}
