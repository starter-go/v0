package iauthx

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type UserPasswordSumComputer struct {
	Username string
	Password string
	Salt     []byte
}

func (inst *UserPasswordSumComputer) Sum() []byte {

	buffer := new(bytes.Buffer)

	buffer.WriteString(inst.Username)
	buffer.WriteString(":")
	buffer.WriteString(inst.Password)
	buffer.WriteString(":")
	buffer.Write(inst.Salt)

	data := buffer.Bytes()
	sum := sha256.Sum256(data)
	return sum[:]
}

func (inst *UserPasswordSumComputer) CheckSum(sum1 []byte) error {
	sum2 := inst.Sum()
	if bytes.Equal(sum1, sum2) {
		return nil
	}
	return fmt.Errorf("bad auth")
}
