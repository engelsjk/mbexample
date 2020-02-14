package main

import (
	"reflect"
	"testing"

	"github.com/iancoleman/strcase"
)

func TestCommands(t *testing.T) {
	v := reflect.ValueOf(CLI)
	typeOfCLI := v.Type()
	for i := 0; i < v.NumField(); i++ {
		cmd := strcase.ToKebab(typeOfCLI.Field(i).Name)
		_, err := NewMap(cmd)
		if err != nil {
			t.Errorf("cmd: %s, error: %s", cmd, err.Error())
		}
	}
}
