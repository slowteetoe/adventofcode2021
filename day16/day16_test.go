package day16

import (
	"testing"
)

func TestExample1(t *testing.T) {
	actual := &[]Packet{}
	parsePackets(expandToBinary("38006F45291200"), 0, actual)

	if len(*actual) != 1 {
		t.Fail()
	}

	o, ok := (*actual)[0].(OperatorPacket)
	if !ok {
		t.Fail()
	}

	if len(o.packets) != 2 {
		t.Fail()
	}

	if (o.packets[0] != Type4Packet{6, 10}) {
		t.Fail()
	}
	if (o.packets[1] != Type4Packet{2, 20}) {
		t.Fail()
	}
}

func TestExpansion(t *testing.T) {
	if expandToBinary("D2FE28") != "110100101111111000101000" {
		t.Fail()
	}
}

func TestType4Packet(t *testing.T) {
	actual := &[]Packet{}
	parsePackets(expandToBinary("D2FE28"), 0, actual)

	if ((*actual)[0] != Type4Packet{6, 2021}) {
		t.Fail()
	}
	if (*actual)[0].Value() != 2021 {
		t.Fail()
	}
}
