package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// CoinsecMainnetPrivate is the version that is used for
// coinsec mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var CoinsecMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// CoinsecMainnetPublic is the version that is used for
// coinsec mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var CoinsecMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// CoinsecTestnetPrivate is the version that is used for
// coinsec testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var CoinsecTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// CoinsecTestnetPublic is the version that is used for
// coinsec testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var CoinsecTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// CoinsecDevnetPrivate is the version that is used for
// coinsec devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var CoinsecDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// CoinsecDevnetPublic is the version that is used for
// coinsec devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var CoinsecDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// CoinsecSimnetPrivate is the version that is used for
// coinsec simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var CoinsecSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// CoinsecSimnetPublic is the version that is used for
// coinsec simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var CoinsecSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case CoinsecMainnetPrivate:
		return CoinsecMainnetPublic, nil
	case CoinsecTestnetPrivate:
		return CoinsecTestnetPublic, nil
	case CoinsecDevnetPrivate:
		return CoinsecDevnetPublic, nil
	case CoinsecSimnetPrivate:
		return CoinsecSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case CoinsecMainnetPrivate:
		return true
	case CoinsecTestnetPrivate:
		return true
	case CoinsecDevnetPrivate:
		return true
	case CoinsecSimnetPrivate:
		return true
	}

	return false
}
