// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestOptionHttpchkNormal0(t *testing.T) {
	parser := &parsers.OptionHttpchk{}
	line := strings.TrimSpace("option httpchk OPTIONS * HTTP/1.1\\r\\nHost:\\ www")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionHttpchkNormal1(t *testing.T) {
	parser := &parsers.OptionHttpchk{}
	line := strings.TrimSpace("option httpchk <uri>")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionHttpchkNormal2(t *testing.T) {
	parser := &parsers.OptionHttpchk{}
	line := strings.TrimSpace("option httpchk <method> <uri>")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionHttpchkNormal3(t *testing.T) {
	parser := &parsers.OptionHttpchk{}
	line := strings.TrimSpace("option httpchk <method> <uri> <version>")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionHttpchkFail0(t *testing.T) {
	parser := &parsers.OptionHttpchk{}
	line := strings.TrimSpace("---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestOptionHttpchkFail1(t *testing.T) {
	parser := &parsers.OptionHttpchk{}
	line := strings.TrimSpace("--- ---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
