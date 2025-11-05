package itokens

import (
	"testing"

	"github.com/starter-go/v0/rbac-web-app/app/classes/tokens"
)

func TestInnerJWTCodec(t *testing.T) {
	// Create a new codec with a test key

	keyStr := "test-secret-key"
	codec := &innerJWTCodec{
		key: []byte(keyStr),
	}

	// Test data
	attrs := map[string]string{
		"user_id": "123",
		"role":    "admin",
		"name":    "testuser",
	}

	// Test encoding
	jwt, err := codec.encode(attrs)
	if err != nil {
		t.Fatalf("Failed to encode JWT: %v", err)
	}

	if jwt == "" {
		t.Fatal("Encoded JWT is empty")
	}

	// Test decoding
	decodedAttrs, err := codec.decode(jwt)
	if err != nil {
		t.Fatalf("Failed to decode JWT: %v", err)
	}

	// Verify the decoded attributes match the original
	if len(decodedAttrs) != len(attrs) {
		t.Fatalf("Decoded attributes length mismatch: got %d, want %d", len(decodedAttrs), len(attrs))
	}

	for k, v := range attrs {
		if decodedAttrs[k] != v {
			t.Errorf("Attribute mismatch for key %s: got %s, want %s", k, decodedAttrs[k], v)
		}
	}
}

func TestInnerJWTCodecInvalidToken(t *testing.T) {
	// Create a new codec with a test key

	keyStr := "test-secret-key"
	codec := &innerJWTCodec{
		key: []byte(keyStr),
	}

	// Try to decode an invalid token
	_, err := codec.decode(tokens.JWT("invalid.token.string"))
	if err == nil {
		t.Fatal("Expected error when decoding invalid token, but got none")
	}
}

func TestInnerJWTCodecWrongKey(t *testing.T) {
	// Create two codecs with different keys

	keyStr1 := "secret-key-1"
	keyStr2 := "secret-key-2"

	codec1 := &innerJWTCodec{
		key: []byte(keyStr1),
	}
	codec2 := &innerJWTCodec{
		key: []byte(keyStr2),
	}

	// Test data
	attrs := map[string]string{
		"user_id": "123",
		"role":    "admin",
	}

	// Encode with first codec
	jwt, err := codec1.encode(attrs)
	if err != nil {
		t.Fatalf("Failed to encode JWT: %v", err)
	}

	// Try to decode with second codec (should fail)
	_, err = codec2.decode(tokens.JWT(jwt))
	if err == nil {
		t.Fatal("Expected error when decoding with wrong key, but got none")
	}
}
