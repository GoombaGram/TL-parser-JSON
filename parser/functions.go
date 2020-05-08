package parser

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ErikPelli/TL-parser-JSON/parser/jsonStruct"
	"strconv"
	"strings"
)

// string utils (thanks to https://www.dotnetperls.com/between-before-after-go)
func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}



// parser functions

func getID(line string) string {
	// Get text after first # (CRC)
	id := between(line,"#", " ")

	// If CRC length is less than 8, add the padding
	if len(id) < 8 {
		for i := len(id); i < 8; i++ {
			id = "0" + id
		}
	}

	// Parse string to byte slice
	idBytes, _ := hex.DecodeString(id)
	// Parse byte slice to int and convert to string representation
	return strconv.FormatInt(int64(int32(binary.BigEndian.Uint32(idBytes))), 10)
}

func getName(line string) string {
	// Get text before first # (name of constructor/method)
	return before(line, "#")
}

func getParams(line string) []*jsonStruct.Param {
	params := make([]*jsonStruct.Param, 0)

	// Get string between first space and " ="
	paramString := between(line, " ", " =")

	// If there aren't params, skip and return an empty slice
	if paramString != "" && !strings.Contains(paramString, "{t:Type}") {
		paramsArray := strings.Split(paramString, " ")
		for _, singleParam := range paramsArray {
			singleParamArray := strings.Split(singleParam, ":")

			param := new(jsonStruct.Param)
			param.Name = singleParamArray[0]
			param.Type = singleParamArray[1]

			params = append(params, param)
		}
	}

	return params
}

func getType(line string) string {
	// Get text after "= " and before ";"
	return between(line, "= ", ";")
}