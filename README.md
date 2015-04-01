# hmac-signature

HMAC Signature example

## Installation

### Using `go get`

	go get github.com/dcu/hmac-signature

### Using `git clone`

	git clone https://github.com/dcu/hmac-signature
	cd hmac-signature
	go build hmac-signature.go

## Usage

	hmac-signature -key="<a key>" -url="<url>" -nonce="<optional nonce>" -params="<param list>"

For example:

	$ hmac-signature -key="key" -url="example.com/path" -nonce="123" -params="c=3 b=2 a=1"
	URL: example.com/path
	Params: a=1&b=2&c=3
	Nonce: 123
	Data: 123example.com/patha=1&b=2&c=3
	
	Signature: R3DT7W9IGZjiRAip9LJTLbEzX981JZfZpd7/gmt8BNU=
	Nonce: 123

if the `nonce` is not provided, one is generated automatically:

	$ hmac-signature -key="key" -url="example.com/path" -params="c=3 b=2 a=1"
	URL: example.com/path
	Params: a=1&b=2&c=3
	Nonce: 1427897981316727363
	Data: 1427897981316727363example.com/patha=1&b=2&c=3
	
	Signature: GiegYK7oTBjSDSzo7/wZ2V0BHasXj313/4N4FV+DsLI=
	Nonce: 1427897981316727363

## Algorithm

You can follow the code and comments [here](https://github.com/dcu/hmac-signature/blob/8a75ba10016b0ebd763f826f90f67747dbcb07a1/hmac-signature.go#L44)



