/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) since 2023, v2rayA Organization <team@v2raya.org>
 */

package sniffing

import (
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"testing"
)

var tlsStream, _ = hex.DecodeString("1603010200010001fc0303d90fdf25b0c7a11c3eb968604a065157a149407c139c22ed32f5c6f486ed2c04206c51c32da7f83c3c19766be60d45d264e898c77504e34915c44caa69513c2221003e130213031301c02cc030009fcca9cca8ccaac02bc02f009ec024c028006bc023c0270067c00ac0140039c009c0130033009d009c003d003c0035002f00ff0100017500000013001100000e7777772e676f6f676c652e636f6d000b000403000102000a00160014001d0017001e00190018010001010102010301040010000e000c02683208687474702f312e31001600000017000000310000000d002a0028040305030603080708080809080a080b080408050806040105010601030303010302040205020602002b0009080304030303020301002d00020101003300260024001d00207fe08226bdc4fb1715e477506b6afe8f3abe2d20daa1f8c78c5483f1a90a9b19001500af00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

func TestName(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	sniffer := NewPacketSniffer(tlsStream)
	d, err := sniffer.SniffTls()
	if err != nil {
		t.Fatal(err)
	}
	if d == "" {
		t.Fatal(d)
	}
}
