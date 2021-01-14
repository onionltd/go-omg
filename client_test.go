package goomg_test

import (
	"bytes"
	"context"
	goomg "github.com/onionltd/go-omg"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_GetCanaryMessage(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		canary, err := ioutil.ReadFile("testdata/canary.txt")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := res.Write(canary); err != nil {
			t.Fatal(err)
		}
	}))
	defer testServer.Close()
	key, err := ioutil.ReadFile("testdata/pgp.txt")
	if err != nil {
		t.Fatal(err)
	}
	keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		t.Fatal(err)
	}
	c := goomg.NewClient(testServer.Client())
	canary, err := c.GetCanaryMessage(context.Background(), testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	_, err = canary.VerifySignature(keyRing)
	if err != nil {
		t.Fatal(err)
	}
	date, err := time.Parse(time.RFC3339, "2019-11-11T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	if err := canary.Validate(date); err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetMirrorsMessage(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		mirrors, err := ioutil.ReadFile("testdata/mirrors.txt")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := res.Write(mirrors); err != nil {
			t.Fatal(err)
		}
	}))
	defer testServer.Close()
	key, err := ioutil.ReadFile("testdata/pgp.txt")
	if err != nil {
		t.Fatal(err)
	}
	keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		t.Fatal(err)
	}
	c := goomg.NewClient(testServer.Client())
	mirrors, err := c.GetMirrorsMessage(context.Background(), testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	_, err = mirrors.VerifySignature(keyRing)
	if err != nil {
		t.Fatal(err)
	}
	urls, err := mirrors.List()
	if err != nil {
		t.Fatal(err)
	}
	expectedUrls := []string{"http://darkfailllnkf4vf.onion", "https://dark.fail"}
	for i := range urls {
		if urls[i] != expectedUrls[i] {
			t.Fatalf("invalid mirrors: %v", urls)
		}
	}
}
