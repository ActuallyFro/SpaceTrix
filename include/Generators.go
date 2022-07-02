package include

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func InitRandSeedTime() {
	rand.Seed(time.Now().UnixNano()) //Random start, to then create a re-usable seed
}

func RandUUID() uuid.UUID {
	return uuid.New()
}

func GenerateInt64() int64 {

	randIntStr := rand.Int63()

	//cast randIntStr as randInt64 (int64)
	randInt64 := int64(randIntStr)

	// randIntStr := fmt.Sprintf("%x", rand.Int63())
	fmt.Println("[DEBUG] [GenerateInt64] generated:", randIntStr)

	return randInt64

}

func UUIDInt64Lower(passedUUID uuid.UUID) int64 {
	return int64(binary.BigEndian.Uint64(passedUUID[0:8]))
}

//Hex print of UUID upper/lower
// fmt.Printf("%x %x\n", u1, u2)

// u2 := UUIDInt64Upper(newUUID)
// fmt.Println("[DEBUG] [UUID] split: ", u1, u2)

func UUIDInt64Upper(passedUUID uuid.UUID) int64 {
	return int64(binary.BigEndian.Uint64(passedUUID[8:16]))
}
