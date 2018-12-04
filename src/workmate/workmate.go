package workmate

import (
	"crypto/rand"
	"fmt"
	"log"
)

// UUID used by entities
func UUID() string {
	// 16 bytes are needed for a valid UUID
	b, err := randBytes(16)
	if err != nil {
		log.Println("Error: Cannot generate random bytes - ", err)
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// SimpleID used by entities
func SimpleID() string {
	b, err := randBytes(2)
	if err != nil {
		log.Println("Error: Cannot generate random bytes - ", err)
		return ""
	}
	return fmt.Sprintf("%d%d", b[0], b[1])
}

func randBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
