package csim

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

func writeString(w io.Writer, command string) (int, error) {
	return w.Write([]byte(command))
}
func FormAtcsim(command []byte) string {
	cmd := strings.ToUpper(hex.EncodeToString(command))
	return fmt.Sprintf("AT+CSIM=%d,\"%s\"\r\n", len(cmd), cmd)
}
func ExpectATResp(expectee io.Reader, expected string) (r string, err error) {
	r = ""
	res := ""

	buff := make([]byte, 530)

	for {
		n, err := expectee.Read(buff)
		if err != nil {
			break
		}

		res += string(buff[:n])

		if n == 0 || strings.Contains(res, "OK\r\n") || strings.Contains(res, "ERROR\r\n") {
			break
		}
	}
	respIndex := strings.Index(res, expected)
	if respIndex >= 0 {
		r := res[respIndex+len(expected):]
		return r, nil
	}
	return "", errors.New("Expected string not found in: " + res)
}
func ParseCsimResp(csimResp string) (response string) {
	response = ""
	cm := strings.Index(csimResp, ",")
	if cm < 0 {
		return
	}
	qu := strings.Index(csimResp[cm+2:], "\"")
	response = csimResp[cm+2 : cm+2+qu]
	return
}
func Csim(transport io.ReadWriteCloser, command []byte) (response []byte, err error) {
	response = []byte{}
	if transport == nil {
		return response, errors.New("com port not initialized")
	}

	c := FormAtcsim(command)

	_, err = writeString(transport, c)
	if err != nil {
		return
	}

	r, err := ExpectATResp(transport, "+CSIM: ")

	response, err = hex.DecodeString(ParseCsimResp(r))

	return

}
