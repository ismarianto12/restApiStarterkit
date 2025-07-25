package utils

import "crypto/rsa"

func LevelAcces(token string) (string, error) {
	levelid, err := DecodeToken(token)
	if err != nil {

	}
	return levelid.LevelID, nil
}

func ParseRSAPublicKeyFromPEM(key []byte) (rsa rsa.PublicKey) {
	// var rsa rsa.PublicKey
	return rsa
}
