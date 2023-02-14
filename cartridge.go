package main

import (
	"fmt"
	"os"
)

type CartridgeHeader struct {
	entry          [4]uint8
	logo           [0x30]uint8
	title          [16]rune
	licenseCode    uint16
	cgbFlag        uint8
	catridgeType   uint8
	romSize        uint8
	ramSize        uint8
	destCode       uint8
	oldLicenceCode uint8
	version        uint8
	checksum       uint8
	globalChecksum uint16
}

type Cartridge struct {
	fname  string
	size   uint32
	data   []uint8
	header *CartridgeHeader
}

var nintendoLogo = [0x30]uint8{
	0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B,
	0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
	0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E,
	0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
	0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC,
	0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
}

var licenseCodes = [0xa5]string{
	0x00: "None",
	0x01: "Nintendo R&D1",
	0x08: "Capcom",
	0x13: "Electronic Arts",
	0x18: "Hudson Soft",
	0x19: "b-ai",
	0x20: "kss",
	0x22: "pow",
	0x24: "PCM Complete",
	0x25: "san-x",
	0x28: "Kemco Japan",
	0x29: "seta",
	0x30: "Viacom",
	0x31: "Nintendo",
	0x32: "Bandai",
	0x33: "Ocean/Acclaim",
	0x34: "Konami",
	0x35: "Hector",
	0x37: "Taito",
	0x38: "Hudson",
	0x39: "Banpresto",
	0x41: "Ubi Soft",
	0x42: "Atlus",
	0x44: "Malibu",
	0x46: "angel",
	0x47: "Bullet-Proof",
	0x49: "irem",
	0x50: "Absolute",
	0x51: "Acclaim",
	0x52: "Activision",
	0x53: "American sammy",
	0x54: "Konami",
	0x55: "Hi tech entertainment",
	0x56: "LJN",
	0x57: "Matchbox",
	0x58: "Mattel",
	0x59: "Milton Bradley",
	0x60: "Titus",
	0x61: "Virgin",
	0x64: "LucasArts",
	0x67: "Ocean",
	0x69: "Electronic Arts",
	0x70: "Infogrames",
	0x71: "Interplay",
	0x72: "Broderbund",
	0x73: "sculptured",
	0x75: "sci",
	0x78: "THQ",
	0x79: "Accolade",
	0x80: "misawa",
	0x83: "lozc",
	0x86: "Tokuma Shoten Intermedia",
	0x87: "Tsukuda Original",
	0x91: "Chunsoft",
	0x92: "Video system",
	0x93: "Ocean/Acclaim",
	0x95: "Varie",
	0x96: "Yonezawa/s’pal",
	0x97: "Kaneko",
	0x99: "Pack in soft",
	0xa4: "Konami (Yu-Gi-Oh!)",
}

var catridgeTypes = [0x100]string{
	0x00: "ROM ONLY",
	0x01: "MBC1",
	0x02: "MBC1+RAM",
	0x03: "MBC1+RAM+BATTERY",
	0x05: "MBC2",
	0x06: "MBC2+BATTERY",
	0x08: "ROM+RAM 1",
	0x09: "ROM+RAM+BATTERY 1",
	0x0b: "MMM01",
	0x0c: "MMM01+RAM",
	0x0d: "MMM01+RAM+BATTERY",
	0x0f: "MBC3+TIMER+BATTERY",
	0x10: "MBC3+TIMER+RAM+BATTERY 2",
	0x11: "MBC3",
	0x12: "MBC3+RAM 2",
	0x13: "MBC3+RAM+BATTERY 2",
	0x19: "MBC5",
	0x1a: "MBC5+RAM",
	0x1b: "MBC5+RAM+BATTERY",
	0x1c: "MBC5+RUMBLE",
	0x1d: "MBC5+RUMBLE+RAM",
	0x1e: "MBC5+RUMBLE+RAM+BATTERY",
	0x20: "MBC6",
	0x22: "MBC7+SENSOR+RUMBLE+RAM+BATTERY",
	0xfc: "POCKET CAMERA",
	0xfd: "BANDAI TAMA5",
	0xfe: "HuC3",
	0xff: "HuC1+RAM+BATTERY",
}

func (c Cartridge) Type() string {
	cartridgeType := catridgeTypes[c.header.catridgeType]

	if cartridgeType != "" {
		return cartridgeType
	}

	return fmt.Sprintf("Unknown cartridge type (0x%d)", c.header.catridgeType)
}

func (c *Cartridge) Load(p string) error {
	d, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	c.size = uint32(len(d))
	c.data = d

	var checksum uint8 = 0x00
	for address := uint16(0x0134); address <= 0x014C; address++ {
		checksum = checksum - d[address] - 1
	}

	c.header.checksum = checksum

	return nil
}
