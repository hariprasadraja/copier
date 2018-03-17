package copier

import (
	"reflect"
	"testing"
)

func TestCopier_CopyStructToMap(t *testing.T) {
	type fields struct {
		Source          interface{}
		SourceType      reflect.Type
		Destination     interface{}
		DestinationType reflect.Type
		Configuration   Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	   {
			name: "convert Stuct to Map",
			fields: fields{
				Source: struct {
					Name string
				    Email string
				}{
					"Hari",
					"",
				},
				Destination: make(map[string]interface{}),
				Configuration: Config{
					OmitEmpty: true,
				},
			},
             wantErr: false,
	   },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copier := &Copier{
				Source:          tt.fields.Source,
				SourceType:      tt.fields.SourceType,
				Destination:     tt.fields.Destination,
				DestinationType: tt.fields.DestinationType,
				Configuration:   tt.fields.Configuration,
			}
			if err := copier.CopyStructToMap(); (err != nil) != tt.wantErr {
				t.Errorf("Copier.CopyStructToMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
