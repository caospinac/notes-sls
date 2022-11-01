package helper

import (
	"time"

	"github.com/google/uuid"
	"github.com/speps/go-hashids/v2"
)

func NewUniqueID() string {
	hashData := hashids.NewData()
	hashData.Salt = uuid.NewString()

	hashID, _ := hashids.NewWithData(hashData)
	timestamp := time.Now().UnixMicro()
	hashString, _ := hashID.Encode([]int{int(timestamp)})

	return hashString
}
