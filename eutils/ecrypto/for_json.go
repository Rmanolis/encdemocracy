package ecrypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
)

func UnmarshalJSON(val string) (*ecdsa.PublicKey, error) {
	pk := new(ecdsa.PublicKey)
	pk.Curve = new(elliptic.CurveParams)
	err := json.Unmarshal([]byte(val), &pk)
	return pk, err
}

func MarshalToJSON(pk *ecdsa.PublicKey) (string, error) {
	out, err := json.Marshal(pk)
	return string(out), err
}
