// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *OptionHttpchk) Init() {
    p.data = nil
}

func (p *OptionHttpchk) GetParserName() string {
	return "option"
}

func (p *OptionHttpchk) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.OptionHttpchk{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *OptionHttpchk) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *OptionHttpchk) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.OptionHttpchk:
		p.data = newValue
	case types.OptionHttpchk:
		p.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}
