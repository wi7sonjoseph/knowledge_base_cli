package util

import (
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// GetTrimmedInterfaceString will trim the interface string
func GetTrimmedInterfaceString(val interface{}, trimVal string) string {
	strVal, _ := val.(string)
	strVal = strings.TrimSpace(strVal)
	return strings.ReplaceAll(strVal, trimVal, "")
}
