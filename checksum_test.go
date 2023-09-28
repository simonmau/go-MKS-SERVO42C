package gomksservo42c

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	var tests = []struct {
		a    []uint8
		want bool
	}{
		{[]uint8{0xE0, 0x30, 0x10}, true},
		{[]uint8{0xE0, 0x30, 0x11}, false},
		{[]uint8{0xE0, 0x30, 0x09}, false},
		{[]uint8{0xE0, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x20}, true},
		{[]uint8{0xE0, 0x33, 0x13}, true},
		{[]uint8{0xE0, 0x3d, 0x1d}, true},
		{[]uint8{0xE0, 0x80, 0x00, 0x60}, true},
		{[]uint8{0xE0, 0x81, 0x01, 0x62}, true},
		{[]uint8{0xE0, 0x83, 0x06, 0x69}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%t", tt.a, tt.want)

		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tt.want, CheckChecksum(tt.a))
		})
	}
}
