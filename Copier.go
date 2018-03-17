package copier

import (
	"encoding/json"
	"log"
	"reflect"
)

// Config contain configuartion flags for Copier
type Config struct {
	OmitEmpty  bool
	OmitByJSON bool
}

type Copier struct {
	Source          interface{}
	SourceType      reflect.Type
	Destination     interface{}
	DestinationType reflect.Type
	Configuration   Config
}

func NewCopier(config Config) *Copier {
	copier := Copier{
		Configuration: config,
	}

	return &copier
}

func (copier *Copier) From(v interface{}) *Copier {

	copier.Source = v
	copier.SourceType = reflect.TypeOf(v)
	return copier
}

func (copier *Copier) To(v interface{}) *Copier {

	copier.Destination = v
	copier.DestinationType = reflect.TypeOf(v)
	return copier
}

func (copier *Copier) CopyStructToMap() error {

	if copier.Configuration.OmitEmpty {
		// fromValue := reflect.ValueOf(copier.Source)
		// toValue := reflect.ValueOf(copier.Destination)
		log.Println(reflect.TypeOf(copier.Source))

		// copy data by json.
		if copier.Configuration.OmitByJSON {
			bytes, err := json.Marshal(copier.Source)
			if err != nil {
				return err
			}

			err = json.Unmarshal(bytes, &copier.Destination)
			if err != nil {
				return err
			}
		}
	}
	log.Println("destination", copier.Destination)
	return nil
}

func (copier *Copier) copyStuctToStruct() {

}

func (copier *Copier) copyMapToStruct() {

}

func (copier *Copier) copyMapToMap() {

}
