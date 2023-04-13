package wol

import (
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strings"
)

// By providing the target MAC address, it crafts a 'Magic Packet' and returns it in a byte array.
// Supports MAC address in three formats:
// 1. FF:FF:FF:FF:FF:FF
// 2. FF-FF-FF-FF-FF-FF
// 3. FFFF.FFFF.FFFF
func CreateMagicPacket(MAC string) ([]byte, error) {
	// Every 'Magic Packet' begins with 6 x 0xFF
	syncStream := "FFFFFFFFFFFF"

	// Check if provided MAC address is valid through a regex
	rxMAC := regexp.MustCompile("(^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$)|(^([0-9A-Fa-f]{4}[.]){2}([0-9A-Fa-f]{4})$)")
	checkMAC := rxMAC.Find([]byte(MAC))
	if checkMAC == nil {
		return nil, fmt.Errorf("%q is not a valid MAC address", MAC)
	}

	// Remove ':', '-' and '.' from MAC address to be eligible to send
	MAC = string(checkMAC)
	delimiter := []string{":", "-", "."}
	for _, d := range delimiter {
		MAC = strings.ReplaceAll(MAC, d, "")
	}

	// Concatenates the MAC address sixteen times
	sixteenMAC := strings.Repeat(MAC, 16)

	// Crafts the 'Magic Packet'
	magicPacket, err := hex.DecodeString(syncStream + sixteenMAC)
	if err != nil {
		return nil, err
	}
	return magicPacket, nil
}

// Send a provided magic packet to a specific broadcast address and port
func SendMagic(magicPacket []byte, bcastAddr string, port int) (err error) {
	ba := net.ParseIP(bcastAddr)
	if ba == nil {
		return fmt.Errorf("%q is not a valid broadcast address", bcastAddr)
	}

	ipport := fmt.Sprintf("%s:%d", ba, port)

	conn, err := net.Dial("udp", ipport)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(magicPacket)
	if err != nil {
		return err
	}

	return nil
}
