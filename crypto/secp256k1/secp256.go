// Tenderly build-compat shim: delegates to go-ethereum's libsecp256k1 binding so
// that libevm and go-ethereum can be linked into one binary without duplicate
// C symbols. Semantics are identical (same upstream C library).
package secp256k1

import (
	"errors"
	"math/big"

	gethsecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var (
	ErrInvalidMsgLen       = errors.New("invalid message length, need 32 bytes")
	ErrInvalidSignatureLen = errors.New("invalid signature length")
	ErrInvalidRecoveryID   = errors.New("invalid signature recovery id")
	ErrInvalidKey          = errors.New("invalid private key")
	ErrInvalidPubkey       = errors.New("invalid public key")
	ErrSignFailed          = errors.New("signing failed")
	ErrRecoverFailed       = errors.New("recovery failed")
)

func Sign(msg []byte, seckey []byte) ([]byte, error) { return gethsecp256k1.Sign(msg, seckey) }

func RecoverPubkey(msg []byte, sig []byte) ([]byte, error) { return gethsecp256k1.RecoverPubkey(msg, sig) }

func VerifySignature(pubkey, msg, signature []byte) bool {
	return gethsecp256k1.VerifySignature(pubkey, msg, signature)
}

func DecompressPubkey(pubkey []byte) (x, y *big.Int) { return gethsecp256k1.DecompressPubkey(pubkey) }

func CompressPubkey(x, y *big.Int) []byte { return gethsecp256k1.CompressPubkey(x, y) }
