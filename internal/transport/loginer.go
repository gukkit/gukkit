package transport

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"gukkit/internal/packet/login"
	"gukkit/net/data/types"
	"sync"

	"github.com/google/uuid"
)

var (
	LoginerPool = &sync.Pool{
		New: func() interface{} {
			return &Loginer{}
		},
	}

	LoginStart         = LoginState(0)
	EncryptionRequest  = LoginState(1)
	EncryptionResponse = LoginState(2)
	LoginSuccess       = LoginState(3)
)

type LoginState int

type Loginer struct {
	session *Session

	Username    types.String
	VerifyToken types.ByteArray
}

func (loginer *Loginer) Success() (err error) {
	err = loginer.session.NextState(Playing)
	return
}

func (loginer *Loginer) Start(pk *login.LoginStartPacket) (err error) {
	pri, pub, err := generateKey()

	if err != nil {
		return err
	}

	loginer.session.PrivateKey = pri
	loginer.session.PublicKey = pub

	h := md5.New()
	request := &login.EncryptionRequestPacket{
		ServerID:    types.String(""),
		PublicKey:   []byte(pub),
		VerifyToken: h.Sum(nil),
	}

	err = loginer.session.SendPacket(request)
	return
}

func (loginer *Loginer) EncryptionResponse(pk *login.EncryptionResponsePacket) (err error) {
	uuid := uuid.New()

	success := &login.LoginSuccessPacket{
		UUID:     types.UUID(uuid),
		Username: types.String(loginer.session.Username),
	}

	if err = loginer.session.SendPacket(success); err != nil {
		loginer.session.UUID = uuid.String()
	}
	return
}

func generateKey() (pri string, pub string, err error) {
	var (
		privateKey *rsa.PrivateKey
		buf        = BufferPool.Get()

		tempPriKey string
	)

	if privateKey, err = rsa.GenerateKey(rand.Reader, 1024); err != nil {
		return
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{

		Type: "RSA PRIVATE KEY",

		Bytes: derStream,
	}

	if err = pem.Encode(buf, block); err != nil {
		return
	}

	tempPriKey = buf.String()
	buf.Reset()

	// 生成公钥文件
	derPkix, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	if err = pem.Encode(buf, block); err != nil {
		return
	}

	pri = tempPriKey
	pub = buf.String()
	return
}
