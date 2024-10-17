package kgscrypto

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"

	"golang.org/x/crypto/bcrypt"
)

type KgsCrypto struct{}

// AESKey is a struct that holds the key and IV for AES encryption.
type AESKey struct {
	Key string
	IV  string
}

// New creates a new KgsCrypto service.
func New() KgsCrypto {
	return KgsCrypto{}
}

// GenerateRandomSecret generates a random secret of the given length.
func (k *KgsCrypto) GenerateRandomSecret(ctx context.Context, length int) ([]byte, *kgserr.KgsError) {
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to generate secret", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return []byte{}, kgsErr
	}

	return bytes, nil
}

// EncryptAES encrypts the given data using AES-CFB encryption.
func (k *KgsCrypto) EncryptAES(ctx context.Context, key AESKey, data string) ([]byte, *kgserr.KgsError) {
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	plaintext := []byte(data)

	block, err := aes.NewCipher([]byte(key.Key))
	if err != nil {
		kgsotel.Error(ctx, err.Error())
		return []byte{}, kgserr.New(kgserr.InternalServerError, "failed to create cipher", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := []byte(key.IV)
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// DecryptAES256 decrypts the given data using AES-CFB decryption.
func (k *KgsCrypto) DecryptAES(ctx context.Context, key AESKey, data []byte) (string, *kgserr.KgsError) {
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	block, err := aes.NewCipher([]byte(key.Key))
	if err != nil {
		kgsotel.Error(ctx, err.Error())
		return "", kgserr.New(kgserr.InternalServerError, "failed to create cipher", err)
	}

	iv := []byte(key.IV)
	plaintext := make([]byte, len(data)-aes.BlockSize)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, data[aes.BlockSize:])

	return string(plaintext), nil
}

// EncryptAESCBC encrypts the given data using AES-CBC encryption.
func (k *KgsCrypto) EncryptAESCBC(ctx context.Context, key string, data string) ([]byte, *kgserr.KgsError) {
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	plaintext := []byte(data)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		kgsotel.Error(ctx, err.Error())
		return []byte{}, kgserr.New(kgserr.InternalServerError, "failed to create cipher", err)
	}

	// Use PKCS7 padding
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	for i := 0; i < padding; i++ {
		plaintext = append(plaintext, byte(padding))
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		kgsotel.Error(ctx, err.Error())
		return []byte{}, kgserr.New(kgserr.InternalServerError, "failed to generate IV", err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// DecryptAESCBC decrypts the given data using AES-CBC decryption.
func (k *KgsCrypto) DecryptAESCBC(ctx context.Context, key string, data []byte) (string, *kgserr.KgsError) {
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		kgsotel.Error(ctx, err.Error())
		return "", kgserr.New(kgserr.InternalServerError, "failed to create cipher", err)
	}

	// Check if the data is long enough to contain the IV
	if len(data) < aes.BlockSize {
		kgsErr := kgserr.New(kgserr.InvalidArgument, "ciphertext too short")
		kgsotel.Error(ctx, kgsErr.Error())
		return "", kgsErr
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad the plaintext
	padding := int(ciphertext[len(ciphertext)-1])
	plaintext := ciphertext[:len(ciphertext)-padding]

	return string(plaintext), nil
}

// HashMD5 hashes the given data using MD5.
func (k *KgsCrypto) HashMD5(ctx context.Context, data string) []byte {
	_, span := kgsotel.StartTrace(ctx)
	defer span.End()

	hasher := md5.New()

	hasher.Write([]byte(data))

	hashInbytes := hasher.Sum(nil)

	return hashInbytes
}

// HashSHA256 hashes the given data using SHA-256.
func (k *KgsCrypto) HashSHA256(ctx context.Context, data string) []byte {
	_, span := kgsotel.StartTrace(ctx)
	defer span.End()

	hasher := sha256.New()

	hasher.Write([]byte(data))

	hashInbytes := hasher.Sum(nil)

	return hashInbytes
}

// HashPassword hashes the given password using bcrypt.
func (k *KgsCrypto) HashPassword(ctx context.Context, data string) (string, *kgserr.KgsError) {
	_, span := kgsotel.StartTrace(ctx)
	defer span.End()

	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to hash bcrypt", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return "", kgsErr
	}

	return string(hash), nil
}

// CompareHashAndPassword compares the given hash and password using bcrypt.
// Returns true if the password matches the hash.
func (k *KgsCrypto) CompareHashAndPassword(ctx context.Context, hash, password string) bool {
	_, span := kgsotel.StartTrace(ctx)
	defer span.End()

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// EncodeHex encodes the given data to a hex string.
func (k *KgsCrypto) EncodeHex(ctx context.Context, data []byte) string {
	return hex.EncodeToString(data)
}

// DecodeHex decodes the given hex string to a byte array.
func (k *KgsCrypto) DecodeHex(ctx context.Context, data string) ([]byte, *kgserr.KgsError) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to decode hex", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return []byte{}, kgsErr
	}

	return bytes, nil
}

// EncodeBase64 encodes the given data to a base64 string.
func (k *KgsCrypto) EncodeBase64(ctx context.Context, data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// DecodeBase64 decodes the given base64 string to a byte array.
func (k *KgsCrypto) DecodeBase64(ctx context.Context, data string) ([]byte, *kgserr.KgsError) {
	bytes, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to decode base64", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return []byte{}, kgsErr
	}

	return bytes, nil
}
