package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/FreightTrackr/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// ReadPrivateKeyFromFile reads an RSA private key from a file
func ReadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	var privateKey *rsa.PrivateKey
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(private.pem): %v", err)
	}
	privateBlock, _ := pem.Decode(privateKeyBytes)
	privateKey, err = x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		log.Fatalf("x509.ParsePKCS1PrivateKey: %v", err)
	}

	return privateKey, nil
}

func ReadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(public.pem): %v", err)
	}
	publicBlock, _ := pem.Decode(publicKeyBytes)
	pub, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		log.Fatalf("x509.ParsePKIXPublicKey: %v", err)
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("not ok: %v", err)
	}
	return publicKey, nil
}

func ReadPrivateKeyFromEnv(private string) (*rsa.PrivateKey, error) {
	privateKeyPEM := os.Getenv(private)
	if privateKeyPEM == "" {
		log.Fatalf("PRIVATE_KEY environment variable not set")
	}

	// Replace escaped newlines with actual newlines
	privateKeyPEM = strings.ReplaceAll(privateKeyPEM, `\n`, "\n")

	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKCS1PrivateKey: %v", err)
	}

	return privateKey, nil
}

func ReadPublicKeyFromEnv(oublic string) (*rsa.PublicKey, error) {
	publicKeyPEM := os.Getenv(oublic)
	if publicKeyPEM == "" {
		log.Fatalf("PUBLIC_KEY environment variable not set")
	}
	publicKeyPEM = strings.ReplaceAll(publicKeyPEM, `\n`, "\n")
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKIXPublicKey: %v", err)
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not ok: %v", err)
	}
	return publicKey, nil
}

func GenerateRSAPem(privateFilename string, publicFilename string, bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return fmt.Errorf("failed to generate RSA private key: %v", err)
	}
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	}
	privFile, err := os.Create(privateFilename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", privateFilename, err)
	}
	defer privFile.Close()
	err = pem.Encode(privFile, privPem)
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %v", err)
	}
	publicKey := &privateKey.PublicKey
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %v", err)
	}
	pubPem := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	pubFile, err := os.Create(publicFilename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", publicFilename, err)
	}
	defer pubFile.Close()
	err = pem.Encode(pubFile, pubPem)
	if err != nil {
		return fmt.Errorf("failed to write public key to file: %v", err)
	}
	return nil
}

func GenerateSecretKeyEnv(privateKeyPath string) (string, string, error) {
	privateKeyPEM, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read private key file: %v", err)
	}
	cleanPrivateKey := CleanPEMString(string(privateKeyPEM))
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return "", "", fmt.Errorf("failed to decode private key PEM block")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse private key: %v", err)
	}
	publicKey := privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal public key: %v", err)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})
	cleanPublicKey := CleanPEMString(string(publicKeyPEM))

	return cleanPrivateKey, cleanPublicKey, nil
}

func CleanPEMString(pem string) string {
	// pem = strings.ReplaceAll(pem, "-----BEGIN RSA PRIVATE KEY-----", "")
	// pem = strings.ReplaceAll(pem, "-----END RSA PRIVATE KEY-----", "")
	// pem = strings.ReplaceAll(pem, "-----BEGIN PUBLIC KEY-----", "")
	// pem = strings.ReplaceAll(pem, "-----END PUBLIC KEY-----", "")
	pem = strings.ReplaceAll(pem, "\n", `\n`)
	// pem = strings.ReplaceAll(pem, "\r", `\r`)
	return pem
}

func SignedJWT(mongoenv *mongo.Database, collname string, user models.Users) (string, error) {
	datauser := FindUser(mongoenv, collname, user)

	claims := jwt.MapClaims{
		"username":      user.Username,
		"nama":          datauser.Nama,
		"no_telp":       datauser.No_Telp,
		"email":         datauser.Email,
		"role":          datauser.Role,
		"no_pend":       datauser.No_Pend,
		"kode_pengguna": datauser.Kode_Pelanggan,
		"exp":           time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privateKey, err := ReadPrivateKeyFromEnv("PRIVATE_KEY")
	if err != nil {
		return "", fmt.Errorf("error loading private key: %v", err)
	}

	t, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing string: %v", err)
	}
	return t, nil
}

func FiberDecodeJWT(c *fiber.Ctx) models.Users {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var session models.Users

	session.Username = claims["username"].(string)
	session.Nama = claims["nama"].(string)
	session.No_Telp = claims["no_telp"].(string)
	session.Email = claims["email"].(string)
	session.Role = claims["role"].(string)
	session.No_Pend = claims["no_pend"].(string)
	session.Kode_Pelanggan = claims["kode_pengguna"].(string)

	return session
}

func StdDecodeJWT(r *http.Request) (models.Users, error) {
	var session models.Users
	tokenString := r.Header.Get("Authorization")
	parts := strings.Split(tokenString, " ")
	tokenString = parts[1]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ReadPublicKeyFromEnv("PUBLIC_KEY")
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return session, fmt.Errorf("invalid token")
	}

	session.Username = claims["username"].(string)
	session.Nama = claims["nama"].(string)
	session.No_Telp = claims["no_telp"].(string)
	session.Email = claims["email"].(string)
	session.Role = claims["role"].(string)
	session.No_Pend = claims["no_pend"].(string)
	session.Kode_Pelanggan = claims["kode_pengguna"].(string)

	return session, nil
}

func FiberIsAdmin(c *fiber.Ctx) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(string)
	return role == "admin"
}

func FiberIsKantor(c *fiber.Ctx) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(string)
	return role == "kantor"
}

func FiberIsPelanggan(c *fiber.Ctx) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(string)
	return role == "pelanggan"
}

func StdIsAdmin(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 {
		return false
	}
	tokenString = parts[1]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ReadPublicKeyFromEnv("PUBLIC_KEY")
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}
	role, ok := claims["role"].(string)
	return ok && role == "admin"
}

func StdIsKantor(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 {
		return false
	}
	tokenString = parts[1]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ReadPublicKeyFromEnv("PUBLIC_KEY")
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}
	role, ok := claims["role"].(string)
	return ok && role == "kantor"
}

func StdIsPelanggan(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 {
		return false
	}
	tokenString = parts[1]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ReadPublicKeyFromEnv("PUBLIC_KEY")
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}
	role, ok := claims["role"].(string)
	return ok && role == "pelanggan"
}
