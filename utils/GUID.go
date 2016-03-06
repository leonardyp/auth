package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"
	"sync/atomic"
	"time"
)

var objectIdCounter uint32 = 0
var machineId = ReadMachineId()

type ObjectId string

func GenGUID() string {
	return NewObjectId().String()
}
func ReadMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err := os.Hostname()
	if err != nil {
		_, err = io.ReadFull(rand.Reader, id)
		if err != nil {
			panic(err)
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

func NewObjectId() ObjectId {
	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]

	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return ObjectId(b[:])
}

func (id ObjectId) String() string {
	return hex.EncodeToString([]byte(id))
}
