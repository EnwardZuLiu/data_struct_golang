package list

import (
	"fmt"
	"testing"
)

func TestConversionNumber(t *testing.T) {
	conversion(100, 16)
	fmt.Println("\n--")
	conversion(100, 8)
}
