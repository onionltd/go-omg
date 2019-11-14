package spec

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
	"regexp"
	"time"
)

type Canary []byte

var ErrExpired = errors.New("canary has expired")

// Validate validates canary message. It looks for existence of mandatory phrases within the signed
// message, a Bitcoin block and a date. It parses the date and checks if the message is valid for given time period.
func (m Canary) Validate(t time.Time) error {
	block, _ := clearsign.Decode(m)
	if block == nil {
		return errors.New("not a clearsigned message")
	}
	if err := findPhrases(block.Plaintext); err != nil {
		return err
	}
	if err := findBitcoinBlockHash(block.Plaintext); err != nil {
		return err
	}
	date, err := getDate(block.Plaintext)
	if err != nil {
		return err
	}
	if t.After(date.Add(14 * 24 * time.Hour)) {
		return ErrExpired
	}
	return nil
}

func getDate(b []byte) (time.Time, error) {
	phrase := "Today is ([0-9]{4}-[0-9]{2}-[0-9]{2})"
	rx, err := regexp.Compile(phrase)
	if err != nil {
		return time.Time{}, nil
	}
	if mg := rx.FindSubmatch(b); mg != nil && len(mg) > 1 {
		b = mg[1]
	} else {
		return time.Time{}, errors.New("date is missing")
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", string(b)))
}

func findBitcoinBlockHash(b []byte) error {
	phrase := "[0]{8}[a-fA-F0-9]{56}"
	matched, err := regexp.Match(phrase, b)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("bitcoin block hash is missing")
	}
	return nil
}

func findDate(b []byte) error {
	phrase := "Today is [0-9]{4}-[0-9]{2}-[0-9]{2}"
	matched, err := regexp.Match(phrase, b)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("date is missing")
	}
	return nil
}

func findPhrases(b []byte) error {
	phrases := []string{
		"I am in control of my PGP key",
		"I will update this canary within 14 days",
	}
	for _, phrase := range phrases {
		matched, err := regexp.Match(phrase, b)
		if err != nil {
			return err
		}
		if !matched {
			return fmt.Errorf("a phrase '%s' missing", phrase)
		}
	}
	return nil
}

// VerifySignature verifies signature of the signed message. If message's signature matches a key in keyRing, it returns
// an identity of the key holder.
func (m Canary) VerifySignature(keyRing openpgp.KeyRing) (*openpgp.Entity, error) {
	block, _ := clearsign.Decode(m)
	if block == nil {
		return nil, errors.New("not a clearsigned message")
	}
	return openpgp.CheckDetachedSignature(keyRing, bytes.NewReader(block.Bytes), block.ArmoredSignature.Body)
}
