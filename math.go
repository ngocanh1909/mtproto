package mtproto

import (
	"bytes"
	"crypto/aes"
	"crypto/rsa"
	sha1lib "crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

const (
	telegramPublicKey_N  = "b3454124f30f3fac04816ee8ceeeb0050c18fd13de4b20721df916ac2f5d4cfe933ea4a65d1f1442a18e55aa396fd5fe9c1470786cdfdf04545e15fde86417900bea0eed95942447fbb5069ee5c91728e7ad5ec23d211af3223ce34c247fee9be9e3522f234aeeed389e3b8ebe2f6f5b6a1404efeb6435dc1efccd6835b161fa12fbf25e624983b45d543ccf7f27b5fb15554e41f5df40cbc7b72fbf06ae0945447a49e3d8dd6f08c1bc66f5c4c1136e6c1e1049efc4adce594890ca2f6323ec8bee7e46b68fb32c7e63d453f9726ddce5521abb2c15ae0f0586be26b6331d35c85f95064cea803a5c38ba2dc1ce0b331f90e862cc336ab28bd61b74d9b94be7"
	telegramPublicKey_E  = 65537
	telegramPublicKey_FP = 2066079364791309323
)

var telegramPublicKey rsa.PublicKey

func init() {
	telegramPublicKey.N, _ = new(big.Int).SetString(telegramPublicKey_N, 16)
	telegramPublicKey.E = telegramPublicKey_E
}

func sha1(data []byte) []byte {
	r := sha1lib.Sum(data)
	return r[:]
}

func doRSAencrypt(em []byte) []byte {
	z := make([]byte, 255)
	copy(z, em)

	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(z), big.NewInt(int64(telegramPublicKey.E)), telegramPublicKey.N)

	res := make([]byte, 256)
	copy(res, c.Bytes())

	return res
}

func splitPQ(pq *big.Int) (p1, p2 *big.Int) {
	value_0 := big.NewInt(0)
	value_1 := big.NewInt(1)
	value_15 := big.NewInt(15)
	value_17 := big.NewInt(17)
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 64, 1)

	what := big.NewInt(0).Set(pq)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := big.NewInt(0)
	i := 0
	for !(g.Cmp(value_1) == 1 && g.Cmp(what) == -1) {
		q := big.NewInt(0).Rand(rnd, rndmax)
		q = q.And(q, value_15)
		q = q.Add(q, value_17)
		q = q.Mod(q, what)

		x := big.NewInt(0).Rand(rnd, rndmax)
		whatnext := big.NewInt(0).Sub(what, value_1)
		x = x.Mod(x, whatnext)
		x = x.Add(x, value_1)

		y := big.NewInt(0).Set(x)
		lim := 1 << (uint(i) + 18)
		j := 1
		flag := true

		for j < lim && flag {
			a := big.NewInt(0).Set(x)
			b := big.NewInt(0).Set(x)
			c := big.NewInt(0).Set(q)

			for b.Cmp(value_0) == 1 {
				b2 := big.NewInt(0)
				if b2.And(b, value_1).Cmp(value_0) == 1 {
					c.Add(c, a)
					if c.Cmp(what) >= 0 {
						c.Sub(c, what)
					}
				}
				a.Add(a, a)
				if a.Cmp(what) >= 0 {
					a.Sub(a, what)
				}
				b.Rsh(b, 1)
			}
			x.Set(c)

			z := big.NewInt(0)
			if x.Cmp(y) == -1 {
				z.Add(what, x)
				z.Sub(z, y)
			} else {
				z.Sub(x, y)
			}
			g.GCD(nil, nil, z, what)

			if (j & (j - 1)) == 0 {
				y.Set(x)
			}
			j = j + 1

			if g.Cmp(value_1) != 0 {
				flag = false
			}
		}
		i = i + 1
	}

	p1 = big.NewInt(0).Set(g)
	p2 = big.NewInt(0).Div(what, g)

	if p1.Cmp(p2) == 1 {
		p1, p2 = p2, p1
	}

	return
}

func makeGAB(g int32, g_a, dh_prime *big.Int) (b, g_b, g_ab *big.Int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 2048, 1)
	b = big.NewInt(0).Rand(rnd, rndmax)
	g_b = big.NewInt(0).Exp(big.NewInt(int64(g)), b, dh_prime)
	g_ab = big.NewInt(0).Exp(g_a, b, dh_prime)

	return
}

func generateAES(msgKey, authKey []byte, decode bool) (aesKey[]byte, aesIV[]byte) {
	var x = 0
	if decode {
		x = 8
	}
	tA := make([]byte, 0, 52)
	tA = append(tA, msgKey[:16]...)
	tA = append(tA, authKey[x:x+36]...)
	sha256A := sha256.Sum256(tA)

	tB := make([]byte, 0, 52)
	tB = append(tB, authKey[40+x:40+x+36]...)
	tB = append(tB, msgKey[:16]...)
	sha256B := sha256.Sum256(tB)

	aesKey = make([]byte, 0, 32)
	aesKey = append(aesKey, sha256A[:8]...)
	aesKey = append(aesKey, sha256B[8:8+16]...)
	aesKey = append(aesKey, sha256A[24:24+8]...)

	aesIV = make([]byte, 0, 32)
	aesIV = append(aesIV, sha256B[:8]...)
	aesIV = append(aesIV, sha256A[8:8+16]...)
	aesIV = append(aesIV, sha256B[24:24+8]...)

	return
}
func Sha256Digest(data []byte) []byte {
	r := sha256.Sum256(data)
	return r[:]
}
func Decrypt(authKey []byte, dataBuf *DecodeBuf) (messageID int64, seqNo int32, data interface{}, err error) {
	var dataLen = int32(dataBuf.Size)

	msgKeyFromServer := dataBuf.Bytes(16)
	dataEncrypted := dataBuf.Bytes(dataBuf.Size - 24)
	// aesKey, aesIV
	aesKey, aesIV := generateAES(msgKeyFromServer, authKey, true)
	// x, err := doDecryptAES256IGE(data, aesKey, aesIV)

	x, err:= doAES256IGEdecrypt(dataEncrypted,aesKey, aesIV)

	if err != nil {
		return 0, 0, nil, err
	}

	dBuf := NewDecodeBuf(x)

	tmpBufArr := make([]byte, dBuf.Size)
	copy(tmpBufArr, dBuf.Buf)

	_ = dBuf.Long() // salt
	_ = dBuf.Long() // session_id
	messageID = dBuf.Long()
	seqNo = dBuf.Int()
	messageLen := dBuf.Int()
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("decrypted data error: Wrong message length %d", messageLen)
		fmt.Println(err)
		return 0, 0, nil, err
	}

	data = dBuf.Object()

	// verify msgKeyFromServer
	messageKey := make([]byte, 32)

		tmpData := make([]byte, 0, 32+len(tmpBufArr))
		tmpData = append(tmpData, authKey[88+8:88+8+32]...)
		tmpData = append(tmpData, tmpBufArr...)
		copy(messageKey, Sha256Digest(tmpData))


	if !bytes.Equal(messageKey[8:8+16], msgKeyFromServer[:16]) {
		err = fmt.Errorf("decrypted data error: (data: %v, aesKey: %s, aseIV: %s, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			data,
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			hex.EncodeToString(authKey[88:88+32]),
			hex.EncodeToString(messageKey[8:8+16]),
			hex.EncodeToString(msgKeyFromServer[:16]))
		fmt.Println(err)
		return 0, 0, nil, err
	}

	return messageID,seqNo,data,nil
}
func doAES256IGEencrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to encrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

func doAES256IGEdecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to decrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil

}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}
