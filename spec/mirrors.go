package spec

import (
	"bufio"
	"bytes"
	"errors"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
	"io"
	"strings"
)

type Mirrors []byte

// List returns a list of mirrors found in the signed message.
func (m Mirrors) List() ([]string, error) {
	block, _ := clearsign.Decode(m)
	if block == nil {
		return nil, errors.New("not a clearsigned message")
	}
	mirrors := []string{}
	reader := bufio.NewReader(bytes.NewReader(block.Bytes))
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		// Handle error after string matching the line.
		// If a file doesn't end with a newline, we would never process the last line.
		if strings.HasPrefix(line, "http://") ||
			strings.HasPrefix(line, "https://") ||
			strings.HasSuffix(line, ".onion") ||
			strings.HasSuffix(line, ".onion/") {
			mirrors = append(mirrors, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return mirrors, nil
}

// VerifySignature verifies signature of the signed message. If message's signature matches a key in keyRing, it returns
// an identity of the key holder.
func (m Mirrors) VerifySignature(keyRing openpgp.KeyRing) (*openpgp.Entity, error) {
	block, _ := clearsign.Decode(m)
	if block == nil {
		return nil, errors.New("not a clearsigned message")
	}
	return openpgp.CheckDetachedSignature(keyRing, bytes.NewReader(block.Bytes), block.ArmoredSignature.Body)
}
