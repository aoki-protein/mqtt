package packet_test

import (
	"reflect"
	"testing"

	packet "github.com/aoki-r/mqtt"
)

func TestPacketType(t *testing.T) {
	type args struct {
		bs []byte
	}

	tests := []struct {
		name string
		args args
		want packet.FixedHeader
	}{
		{
			"Reserved",
			args{[]byte{0x10, 0x00}},
			packet.FixedHeader{0},
		},
		{
			"CONNECT",
			args{[]byte{0x10, 0x00}},
			packet.FixedHeader{1},
		},
		{
			"CONNECT",
			args{[]byte{0x10, 0x00}},
			packet.FixedHeader{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := packet.ToFixedHeader(tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFixedHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
