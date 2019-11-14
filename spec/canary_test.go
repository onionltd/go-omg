package spec

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/errors"
	"io/ioutil"
	"testing"
	"time"
)

var canaryMessage = Canary(`-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

I am the admin of DarkDotFail.
I am in control of my PGP key.
I will update this canary within 14 days.
Today is 2019-10-29. 

Latest bitcoin block hash: 
00000000000000000006a5d6871d8407abcba2a5ff2546dcbc179ddae0950331
-----BEGIN PGP SIGNATURE-----

iQIzBAEBCAAdFiEEbf2uZtQ/we7OuH584uRp3H2MPaIFAl24uQEACgkQ4uRp3H2M
PaLrXw/+Mm5oXm2Ttyr44/nUeUDihuY4vh4ooxl4zo7wBuhka0j2VEYnBO8jZXig
oHszZ+oyZ31fen7a6GgkgXJXrZHBNT6+kxw8hWCmroJzxwLLmByGbm0ezuOFwa9l
ZDhaLutbOjrL1xHNNVJVSWoivlbzToqOhePdqOTmr5bXiMLpalBFvM+BLG1eM764
Au9GOThToI9HDDFOGO7SeoZwaYgvCMc5JhKY9LwWO1/BMjdl9tZ9aUF4sqf23h5v
2caWu+rr7pJSDmpeeMI8zHMhMc+dj+7CyZ3bMFfqqD5WacSqlMDsD1ROY/hU8nyh
JaBogc/Cw/1YlZ6rpfL/3bZa+UR5vhwLlQAcvomfZZRfZoi4FIlKvN/JXhtgZHE+
v2sl85WTURngTbep5XRTqJBoQzNjqnqRZql5bqe4rWsP+wuWMkG0prtZye/dphXs
W2O+KdVz5pLOXI5thwaYQEu77uqvDKR/+40U4LcipM0XfTyTZt89c++FRFioVP+l
uxwvwKiNz2Nlb2C7CtMwNhbCACgEMKvQacMLNQg8tYwyuKIyJONZzpMe2hq2xoRp
daWKD2tQ4YLPJDbAMQdi1L7UA3Wdk41RhYWpFUc5Cc5inOp+o7ZNWY2OfVRidkwm
Kx8e7I7fSVYEpPNMWfRzi6ucQ0iOt8I6GNVqjwOYIa6f2ee36E8=
=uMGr
-----END PGP SIGNATURE-----`)

func TestCanary_Validate(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2019-11-11T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	if err := canaryMessage.Validate(date); err != nil {
		t.Fatal(err)
	}
}

func TestCanary_ValidateExpired(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2019-11-13T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	if err := canaryMessage.Validate(date); err != ErrExpired {
		t.Fatal(err)
	}
}

func TestCanary_VerifySignature(t *testing.T) {
	key, err := ioutil.ReadFile("testdata/pgp.txt")
	if err != nil {
		t.Fatal(err)
	}
	keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := canaryMessage.VerifySignature(keyRing); err != nil {
		t.Fatal(err)
	}
}

func TestCanary_VerifySignatureWrongKey(t *testing.T) {
	key, err := ioutil.ReadFile("testdata/pgp2.txt")
	if err != nil {
		t.Fatal(err)
	}
	keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		t.Fatal(err)
	}
	if _, err = canaryMessage.VerifySignature(keyRing); err != errors.ErrUnknownIssuer {
		t.Fatal(err)
	}
}
