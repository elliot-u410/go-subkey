package subkey

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vedhavyas/go-subkey/common"
	"github.com/vedhavyas/go-subkey/ed25519"
	"github.com/vedhavyas/go-subkey/sr25519"
)

func TestDerive(t *testing.T) {
	testsMap := map[Scheme][]struct {
		uri       string
		publicKey string
		ss58Addr  string
		network   uint8
		err       bool
	}{
		sr25519.Scheme{}: {
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap",
				publicKey: "0x88af895626c47cf1235ec3898d238baeb41adca3117b9a77bc2f6b78eca0771b",
				ss58Addr:  "5F9vWoiazEhfxSxCG8nUuDhh5fqNtPnSxp2BrhPsuLqEQASi",
				network:   42,
			},

			{
				uri:       "0x18446f2d685492c3086391aabe8f5e235c3c2e02521985650f0c97052237e717",
				publicKey: "0x88af895626c47cf1235ec3898d238baeb41adca3117b9a77bc2f6b78eca0771b",
				ss58Addr:  "5F9vWoiazEhfxSxCG8nUuDhh5fqNtPnSxp2BrhPsuLqEQASi",
				network:   42,
			},

			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap///password",
				publicKey: "0x5c2d57c4cfa7df7a9d0e9546bb575045f5ec14e9771de8bc907910c84cd5de2a",
				ss58Addr:  "5E9ZjRM9VdqES5JhbABVpvgCstaE7J5x3cE7sTKMGG5TF8tZ",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap/foo",
				publicKey: "0x287061f5973551d070ccc62fb4563a0be2e6324ce183c456850e342aa021f94d",
				ss58Addr:  "5CyjA4yQrQtJBs7jC4D6S672y3Ez4Shd3se6VXB4JBkdGwUZ",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo",
				publicKey: "0x04bd4f94429371e044509d22f8a6d33ab9c336bf54ef6b38eba0cc3a4f125e5a",
				ss58Addr:  "5CAvHXaqNRwbbL4B3MoQJdam8JmotCGAF8kTpgWhR9ahhJYS",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo/bar",
				publicKey: "0x0c6febc87c461f8ddceb295d90c3ba999b1e93c2bdd13145b265512d06729449",
				ss58Addr:  "5CM1gMJkyRoE7txkdHv31y6H4yPMKCALSDpaeaE8BpDVwrht",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap/foo//bar",
				publicKey: "0xe4535b3b8e259badc3c78128bfafe0b50df625862edaff7c9d68999a0811865b",
				ss58Addr:  "5HE5Y6MDZvy9QJsmgjrnJHiSqsYRTrfBLrzLvHQC3f9PM6TR",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo/bar//42/69",
				publicKey: "0x68a5a8f7e29ffcae1d15518b180f6e4f1132b45ffd565cb7953045faf07c8809",
				ss58Addr:  "5ERv3mLP7CX1CViNc6NUQaePBJMkf6BELffpMfXjXjj28SNo",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo/bar//42/69///password",
				publicKey: "0x4055514cd4ddcc7b23024839b68190f3f71bc262eb038145262bfe087bbb5429",
				ss58Addr:  "5DX4GQQm9rSHVcqaG9CgxdZLsj8buBxcRWEYYcHrRXe4epZg",
				network:   42,
			},

			{
				uri:       "bottom drive obey lake curtain smoke basket hold race lonely fit walk",
				publicKey: "0x46ebddef8cd9bb167dc30878d7113b7e168e6f0646beffd77d69d39bad76b47a",
				ss58Addr:  "5DfhGyQdFobKM8NsWvEeAKk5EQQgYe9AydgJ7rMB6E1EqRzV",
				network:   42,
			},
		},
		ed25519.Scheme{}: {
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap",
				publicKey: "0xe4631cda48cb885f3a6d0b521d3278ec3e834dd2e1766f7edb8e1386535cc217",
				ss58Addr:  "5HEADZuqsQzNPxGySd74DGPhfm8vFFPVGaKPWkQigJgtv41f",
				network:   42,
			},

			{
				uri:       "0x18446f2d685492c3086391aabe8f5e235c3c2e02521985650f0c97052237e717",
				publicKey: "0xe4631cda48cb885f3a6d0b521d3278ec3e834dd2e1766f7edb8e1386535cc217",
				ss58Addr:  "5HEADZuqsQzNPxGySd74DGPhfm8vFFPVGaKPWkQigJgtv41f",
				network:   42,
			},

			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap///password",
				publicKey: "0x261a29a2b6f690f394d339dc6e09f7f8fa85a3ed82b7567e2bb2a79c33651eef",
				ss58Addr:  "5CvfSyhefVmXnmQ2c4ff6h4EBuhNqaRpjoEHyMD8JWdnpH7y",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo",
				publicKey: "0x986f6247a100aee1aaaadb215fc681f95a64a86fd1f12d4360514f9be7769f40",
				ss58Addr:  "5FWaDvLD9wuZRiLzCxECXdrc57Xavjh5WMvC54ufMQmvPTxD",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo//42",
				publicKey: "0x7a16bd534b1aab9d420d5ca544927ccff88f76e39b063faee502b63f7a2fb394",
				ss58Addr:  "5EpnTJ2E731sTG9WnHNS2cbcppriXx7RF8nmRSaBHWg5hRSr",
				network:   42,
			},
			{
				uri:       "crowd swamp sniff machine grid pretty client emotion banana cricket flush soap//foo//42///password",
				publicKey: "0x34f7460f79c0c4947dfe1b4176ff8cf974883ed2f2a5c716ed89bd16b11e05dc",
				ss58Addr:  "5DG9oWqVMaxTn7LksujDvYPQEcU19yGiEkgAEHFYoBtYudM9",
				network:   42,
			},
		},
	}

	for scheme, tests := range testsMap {
		for _, c := range tests {
			s, err := Derive(scheme, c.uri)
			if err != nil {
				assert.True(t, c.err)
				continue
			}

			pub := s.Public()
			assert.Equal(t, c.publicKey, common.EncodeHex(pub))
			gotSS58Addr, err := s.SS58Address(common.Network(c.network), common.SS58Checksum)
			assert.NoError(t, err)
			assert.Equal(t, c.ss58Addr, gotSS58Addr)
		}
	}
}

func TestKeyRing_Sign_Verify(t *testing.T) {
	uri := "0xd2dbfa26295528f3893430047b773e5bc5457b02c520c5d80bb83366d42de032"
	kr, err := Derive(sr25519.Scheme{}, uri)
	assert.NoError(t, err)
	msg := []byte("testmessage")
	sig, err := kr.Sign(msg)
	assert.NoError(t, err)
	assert.True(t, kr.Verify(msg, sig))
	fmt.Println(hex.EncodeToString(sig[:]))
}
