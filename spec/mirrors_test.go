package spec_test

import (
	"bytes"
	"github.com/onionltd/go-omg/spec"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/errors"
	"io/ioutil"
	"testing"
)

var mirrorsMessage = spec.Mirrors(`-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

http://darkfailllnkf4vf.onion
https://dark.fail

# This is an example /mirrors.txt file, a requirement if you want your site listed on https://dark.fail.
#
# Mirrors.txt Rules:
#  - Mirrors must be signed by a PGP key which is in /pgp.txt hosted at all of your URLs.
#  - Any line in this file which begins with “http://“ or “https://“ is an official mirror of your site.
#  - Mirrors must all host the same content. No forums, no link lists. Place those in /related.txt following these same standards.
#  - All valid mirrors must only contain a scheme and domain name, no ports or paths.
#  - /pgp.txt and /mirrors.txt must have the same content on all of your URLs.
#  - Text which is not intended to be parsed as an official mirror must be commented out with a “#” as the first character on the line.
-----BEGIN PGP SIGNATURE-----

iQIzBAEBCAAdFiEEbf2uZtQ/we7OuH584uRp3H2MPaIFAl24tUEACgkQ4uRp3H2M
PaLLfRAAql/RUatr8o2oCf6YdUXh16Y2ODR7a75nVgPOL+n9pnPRKBVu+HPgDeE0
EKZS00si9mmHQrMR2Lmv7wtxenVKKK2HgYGMxYYJdBeDr5pt7sU0y4/KRfVwoQ/D
mmJqDNz/uAHS7+tiwH31CmF8ZqPZvMaxHGuv20qCH0ZOrAx6Pvv9LxYSQaikq7+s
C+k0VCKx5dnWm5hZe72HB6pup+cHu3c8doYmhZKFEgi4sI0Z8SNrk3siyM7rMhcc
qhS4DlNhWXaz0LhV4eE8440V1HFLw0qUyNWFcYkqm0+V2KNQCyDTSnXnP6U9/5hw
QEeXeNCMNqY6CS39BssxhBB+gvPT+oOgEv1M7D+uBi6XSj45kWGYit9TNroTOOOP
1qyDfbg48/doaxIrSRzV7M5xHkeXgnjaUC+cKUxlvDzLtEc7gDzPkOEjDmCtbmX1
sR8pb3AFuoKyGbjOITVbKml3zoWQw4JB1sdAjvQQYsSZ3C4xrS4r2AfUZ+dKoXB/
RuVpH0RGtf7EBExysHrbN6sh0Zi6MvBNoV2/Q8/Jq45ujvAX6CRXI+3jOBqQLqc4
y1I82nHP1ITB0Uaxt44PrGa0CVRWNCcYkRDvS3FWyUQ4XN+IyTbKirFQRBdg+orK
VG9UkGZgzBKUrWwRqX5NiqsTb6KZlKb6PQ7cvDloGCFAzE0jGvg=
=0Cyh
-----END PGP SIGNATURE-----`)

func TestMirrors_List(t *testing.T) {
	mirrors, err := mirrorsMessage.List()
	if err != nil {
		t.Fatal(err)
	}

	expectedMirrors := []string{"http://darkfailllnkf4vf.onion", "https://dark.fail"}

	if len(mirrors) != len(expectedMirrors) {
		t.Fatalf("unexpected number of mirrors: %v", mirrors)
	}

	for i := range mirrors {
		if mirrors[i] != expectedMirrors[i] {
			t.Fatalf("invalid mirrors: %v", mirrors)
		}
	}
}

func TestMirrors_VerifySignature(t *testing.T) {
	key, err := ioutil.ReadFile("testdata/pgp.txt")
	if err != nil {
		t.Fatal(err)
	}
	keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		t.Fatal(err)
	}
	_, err = mirrorsMessage.VerifySignature(keyRing)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMirrors_VerifySignatureWrongKey(t *testing.T) {
	key, err := ioutil.ReadFile("testdata/pgp2.txt")
	if err != nil {
		t.Fatal(err)
	}
	keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		t.Fatal(err)
	}
	if _, err = mirrorsMessage.VerifySignature(keyRing); err != errors.ErrUnknownIssuer {
		t.Fatal(err)
	}
}
