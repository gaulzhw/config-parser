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


func TestOptionHTTPLogNormal0(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
	line := strings.TrimSpace("option httplog")
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
func TestOptionHTTPLogNormal1(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
	line := strings.TrimSpace("no option httplog")
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
func TestOptionHTTPLogNormal2(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
	line := strings.TrimSpace("option httplog clf")
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
func TestOptionHTTPLogNormal3(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
	line := strings.TrimSpace("option httplog # comment")
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
func TestOptionHTTPLogNormal4(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
	line := strings.TrimSpace("option httplog clf # comment")
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
func TestOptionHTTPLogFail0(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
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
func TestOptionHTTPLogFail1(t *testing.T) {
	parser := &parsers.OptionHTTPLog{}
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
