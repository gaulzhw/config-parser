package global

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type NbProc struct {
	Enabled bool
	Value   int64
}

func (n *NbProc) Init() {
	n.Enabled = false
}

func (n *NbProc) GetParserName() string {
	return "nbproc"
}

func (n *NbProc) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "nbproc") {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) != 2 {
			return "", &errors.ParseError{Parser: "NbProc", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(elements[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "NbProc", Line: line, Message: err.Error()}
		} else {
			n.Enabled = true
			n.Value = num
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "nbproc", Line: line}
}

func (n *NbProc) Valid() bool {
	if n.Enabled {
		return true
	}
	return false
}

func (n *NbProc) String() []string {
	if n.Enabled {
		return []string{fmt.Sprintf("  nbproc %d", n.Value)}
	}
	return []string{}
}