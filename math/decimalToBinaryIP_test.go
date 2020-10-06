package math

import "testing"

type ipTest struct {
	ipDec string
	ipBin string
}

type decBinTest struct {
	dec int
	bin string
}

func TestDecimalToBinaryID(t *testing.T) {
	ips := []ipTest{
		{ipDec: "192.168.0.1",
			ipBin: "11000000.10101000.00000000.00000001"},
		{ipDec: "207.255.64.110",
			ipBin: "11001111.11111111.01000000.01101110"},
		{ipDec: "123.213.132.231",
			ipBin: "01111011.11010101.10000100.11100111"},
	}
	for _, ip := range ips {
		if ipBin := DecimalToBinaryIP(ip.ipDec); ipBin != ip.ipBin {
			t.Errorf("Resultado incorreto.\nDeve ser: %s\nObteve-se: %s",
				ip.ipBin, ipBin)
		}
	}
}

func TestDecimalToBinary(t *testing.T) {
	testCases := []decBinTest{
		{dec: 192, bin: "11000000"},
		{dec: 168, bin: "10101000"},
		{dec: 213, bin: "11010101"},
	}
	for _, test := range testCases {
		if got := DecimalToBinary(test.dec); got != test.bin {
			t.Errorf("Resultado incorreto.\nDeve ser: %s\nObteve-se: %s",
				test.bin, got)
		}
	}
}
