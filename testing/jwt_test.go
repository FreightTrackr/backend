package testing

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/FreightTrackr/backend/middleware"
	"github.com/FreightTrackr/backend/utils"
)

func TestGenerateRSAPem(t *testing.T) {
	err := utils.GenerateRSAPem("../keys/private.pem", "../keys/public.pem", 1024)
	if err != nil {
		log.Fatalf("Error generating private key file: %v", err)
	}
	fmt.Println("Private key successfully generated and saved to ./keys/private.pem")
}
func TestGenerateSecretKeyEnv(t *testing.T) {
	privateKeyPEM, publicKeyPEM, err := utils.GenerateSecretKeyEnv("../keys/secret.pem")
	if err != nil {
		log.Fatalf("Error converting private key to strings: %v", err)
	}
	fmt.Println("Private Key as String:")
	fmt.Println(privateKeyPEM)

	fmt.Println("\nPublic Key as String:")
	fmt.Println(publicKeyPEM)

	envFileContent := fmt.Sprintf("PRIVATE_KEY=%s\nPUBLIC_KEY=%s\n", privateKeyPEM, publicKeyPEM)

	err = os.WriteFile("../.env.secretkey", []byte(envFileContent), 0600)
	if err != nil {
		log.Fatalf("Error writing to .env.secretkey: %v", err)
	}

	fmt.Println(".env.secretkey file has been created with PRIVATE_KEY and PUBLIC_KEY.")
}
func TestReadPrivateKeyFromEnv(t *testing.T) {
	// The private key in PEM format
	privateKeyPEM := ``

	// Parse the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyPEM)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Print the private key (or use it as needed)
	fmt.Println("Successfully decoded and parsed the private key.")
	fmt.Println(privateKeyBytes)
}

func TestReadPrivateKeyFromEnvaaa(t *testing.T) {
	middleware.LoadEnv("../.env")
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")
	PUBLIC_KEY := os.Getenv("PUBLIC_KEY")
	fmt.Println("PRIVATE_KEY")
	fmt.Println(PRIVATE_KEY)
	fmt.Println("PUBLIC_KEY")
	fmt.Println(PUBLIC_KEY)
}
