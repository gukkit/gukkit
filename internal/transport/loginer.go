package transport

import (
	"bytes"
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
	var (
		privateKey *rsa.PrivateKey
		publicKey  interface{}

		buf = BufferPool.Get().(*bytes.Buffer)
	)
	buf.Reset()

	loginer.Username = pk.Name

	defer BufferPool.Put(buf)

	if privateKey, err = rsa.GenerateKey(rand.Reader, 1024); err != nil {
		return
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	if err = pem.Encode(buf, block); err != nil {
		return
	}

	defer func(privateKey string) {
		if err == nil {
			loginer.session.PrivateKey = privateKey
		}
	}(buf.String())

	derpStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derpStream,
	}

	buf.Reset()
	if err = pem.Encode(buf, block); err != nil {
		return
	}

	h := md5.New()
	h.Write(buf.Bytes())

	request := &login.EncryptionRequestPacket{
		ServerID:    types.String(""),
		PublicKey:   buf.Bytes(),
		VerifyToken: h.Sum(nil),
	}

	err = loginer.session.SendPacket(request)
	return
}

func (loginer *Loginer) EncryptionResponse(pk *login.EncryptionResponsePacket) (err error) {

	success := &login.LoginSuccessPacket{
		UUID:     types.UUID(uuid.New()),
		Username: "",
	}

	err = loginer.session.SendPacket(success)
	return
}
