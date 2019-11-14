# go-omg

*go-omg* is an HTTP client and a parser of files specified in Onion Mirror Guidelines.
See spec.txt.

## Features

* download and validate `/mirrors.txt`
* download and validate `/canary.txt`
* download and validate `/related.txt`

The package intentionally doesn't implement a method to download PGP keys from `/pgp.txt`.
This is to save people from dangerous assumption that verifying signed messages with a key downloaded
from the same host, protects them from phishing.

In order to verify signed data, you must trust the key first. To do so, it must be obtained from
a trusted source.
