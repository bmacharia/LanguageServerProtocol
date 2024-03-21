package rpc_test

import (
	"lspeductional/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

// this test to make sure that we can encode a message
func TestEncode(t *testing.T) {
	// expected value of the content encoded message
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	// what does the message lool like
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	// does expected = actual if not then fail the test
	if expected != actual {
		t.Errorf("Expected %s,  Actual %s", expected, actual)

	}

}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Errorf("Expected 16,  Got %d", contentLength)
	}
	if method != "hi" {
		t.Errorf("Expected: 'hi',  Got %s", method)
	}
}
