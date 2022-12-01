package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	jose "gopkg.in/square/go-jose.v2"
)

const (
	// Copy/pasted from the JSON field `payload` in the network traffic.
	payloadBase64 = "eyJiYWNrdXBUeXBlIjoiTVNBIiwiY3JlYXRpb25EYXRlIjoiV2VkIEp1biAwMSAyMTo1NDoyOSBQRFQgMjAyMiIsImRldmljZUlkZW50aWZpZXIiOiI4MDdhMTRlZi0zNGVmLTRlODgtYjAwZS04M2FmMGU1NzY3MWUiLCJkZXZpY2VOYW1lIjoiUGl4ZWwgM2EiLCJlbmNyeXB0ZWRCYWNrdXAiOiJleUpoYkdjaU9pSkJNVEk0UzFjaUxDSmxibU1pT2lKQk1USTRRMEpETFVoVE1qVTJJbjAuT3BzdjkzM2pMOEhQakZtWTZ6UFREUDEtamZvUVd0MEI5VW5XOWxuZV9ZWFcza0ExOFBBbnRRLnpsU1JscERvdU4yT3R5eTlhc2JBclEuOFVXOGtrSVZrRjQ5X0p1NUVmQ28zdXNuQzFqaEZscTA1Vm03UmlocE9tMjdXTERoXzlKZzczOXp0RXluRmJobExnaEpBbzNrTEJoVTU1UFg0TXlUcFNtZDhMN2tlMENWbGswSk4tbUp2dW1XU0RfRDFQVmVkY3E3SDRPcWI2NHdzcXg4UGlCZDU5aXU4WGlzSFBOUkFsY2I2UGsxV1pmUWNwaHhacFhzcE4zU0x5ZERFMWdtR2xieEFuUW9ucERzTk9xX0d5NWVKeU5DN2JIRXE2M25NajVvbXBvZEx2YjhFZjNlNFJyczBhbVpCaElVUlVxUmwtbW0yT1Y4SXI4ZVJ6bU8yZlVybEJ6M0UxbUtvSllQQWZqUE1WTjZ4UU5XZUdnU0FMZkhSTjN0WklaM0R3bFJ1eUMwU1ZrOEZFMkNIUUtqWWlnTGVDYWYxWEJEU1JFSGh5Nnp3YkZmakNvaFoyTld1cTNqaUpZQllLN2FLZ3c1TVhwVXd4bjJhT1c4d3Q3d3JNbXlaTzdCZXcyUEVOMWl5X2N6TnBXNkZNQjJIRVJhMUh6MWFJWm9MMjhPdEw0YldKZ0puSzVKNUlNRlhKVHBYR050bDgyUElrbVJGUkY0MGFuakhPd2R5WkVEQUFsNnFUcktRRnhUZVZLNVE0bEp1SmkycmZORUpnaElJY1hYRGNjUDQxNG1tYnV3TGw4NEdzaWNGZnMyYTNpcUp4SFJJRmFfS1ZiVXBNUVBGM3daMWVDTjdmX3dUZkdlUUI5TlRyb1RCc3lReUIwVk14SHBDcTBSZlpNQnBoZWtYUnRKQkRXU2VQNFNDcWlIVnB3VnRXOEpTWTRNZHRSb3djZFlhQm5wMUtucXpxWkJoSkxRSzZtd1B0WGlLYVZBQkFCV2hobjQtV2lUdGh3a1R5Nm0tYUVtZk9wLWNjekRpOU5GS1F3QWg5bEhCaWVwbFY4dmhTdkFPRnBYSVdOeFRFem5NNVRZeGt5NkEzbTV3cWozM29oYmJyb0YyVVNjRi1tZEUtNERiVWFQRVpJTXlkWWJvdFpja3BGS0Z2VHVzUkxzb0NnUkU2VFJEMy1FY0tlY3NhVnlMQllONnRKa3V2bUFYaE5mV0ZwX21DN0tvMzYtVV9la0ZxcHNDWktuaVlieVdOdkdDaE5kb1AwS1FhbmVDUFBvRjBucW9yaEh2dElMNzVmTzFSWWo1TmFCR3B6V3hBSWlSRmNWMEpSY0xSZ0ZNbnI1bTM5WmZnemN1OTBRbHphckM5UUlueDhLV0FvV1dfcGxqdk1HT1hua2w2dUNSMlcyOWYyT2V6VW9Vek9LMFpJLVdFQkxMLXAzQkJrbDdYeHdLRm9ZTFk3NWhqTzFDTzIwYml3VmdTSFZ4N3BZdFVCNTlQYml2cXBQemNwRC1NNEpkM0w0WDhfS3A0YVVKVjNWRDd3azlzZXZWMGxsaUg0amJqM0hqWHBEY2hTaFhZSG9ScUpEZnl6aVFQdXJSYnlab0NfZWZJamlobjlwRVpiUnQ2aVhHcnRmeWd5RGU1SXZpVEloV3NjdTZpSllhbnhHSV9vckgtbDg4Z3c4Q0J3TWN2XzV4Y21kVWpWdkRlVDIwU0xCbG1iUktPYkotVnBwUmlkekhuai1QT1lkbWU3V000eHFBUV9KQk5QTGktblMyVTVobVNxcnN4d0xESWcwZkJtSlVzeXkuNEtuNGtIYjBOM3lsUWFIZlZrQmV4QSIsIm1zYUtleUlkIjoiMjAyMi0wNi0wMlQwNDo1NDoxN1oiLCJ1cGRhdGVkRGF0ZSI6IldlZCBKdW4gMDEgMjE6NTQ6MzIgUERUIDIwMjIifQ\u003d\u003d"

	// Copy/pasted from the XML node <ps:Property name="KeyMaterial"> in the
	// network traffic.
	keyBase64 = "SbQv5hX0zBCLw70reWvBX3PvmzChD1QL6UnpGFfG/5w="
)

func main() {
	// Decode the Base64 payload.
	jsonPayload, err := base64.StdEncoding.DecodeString(payloadBase64)
	if err != nil {
		panic(err)
	}

	// Parse the JSON payload into a JSON Web Encryption object.
	var data struct {
		EncryptedBackup string
	}
	err = json.Unmarshal(jsonPayload, &data)
	if err != nil {
		panic(err)
	}

	jwe, err := jose.ParseEncrypted(data.EncryptedBackup)
	if err != nil {
		panic(err)
	}

	// Decode the Base64 key.
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		panic(err)
	}

	// Use the key to decrypt the ciphertext in the JWE.
	plaintext, err := jwe.Decrypt(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Plaintext TOTP backup = %s\n", string(plaintext))
}
