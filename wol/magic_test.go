package wol

import (
	"encoding/hex"
	"strings"
	"testing"
)

func craftMagic(MAC string) []byte {
	syncStream := "FFFFFFFFFFFF"
	delimiter := []string{":", "-", "."}
	for _, d := range delimiter {
		MAC = strings.ReplaceAll(MAC, d, "")
	}
	sixteenMAC := strings.Repeat(MAC, 16)
	magicPacket, _ := hex.DecodeString(syncStream + sixteenMAC)

	return magicPacket
}

func TestCreateMagicPacket(t *testing.T) {
	MACOne := "FF:FF:FF:FF:FF:FF"
	MACheck := craftMagic(MACOne)
	retOne, errOne := CreateMagicPacket(MACOne)
	if errOne != nil {
		t.Fail()
	} else if hex.EncodeToString(retOne) != hex.EncodeToString(MACheck) {
		t.Fail()
	}

	MACTwo := "FF-FF-FF-FF-FF-FF"
	MACheck = craftMagic(MACTwo)

	retTwo, errTwo := CreateMagicPacket(MACTwo)
	if errTwo != nil {
		t.Fail()
	} else if hex.EncodeToString(retTwo) != hex.EncodeToString(MACheck) {
		t.Fail()
	}

	MACThree := "FFFF.FFFF.FFFF"
	MACheck = craftMagic(MACTwo)

	retThree, errThree := CreateMagicPacket(MACThree)
	if errThree != nil {
		t.Fail()
	} else if hex.EncodeToString(retThree) != hex.EncodeToString(MACheck) {
		t.Fail()
	}
}
