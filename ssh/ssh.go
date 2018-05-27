package ssh

import (
	"github.com/banzaicloud/pipeline/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var log *logrus.Entry

// Simple init for logging
func init() {
	logger = config.Logger()
	log = logger.WithFields(logrus.Fields{"action": "ssh"})
}

// Get SSH Key from Bank Vaults
func KeyGet() (string, error) {
	return "", nil
}

// Add SSH Key to Bank Vaults
func KeyAdd(key string) (string, error) {
	return key, nil
}

// Delete SSH Key from Bank Vaults
func KeyDelete(key string) (string, error) {
	return key, nil
}

//privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//if err != nil {
//log.Panicf("PrivateKey generator failed reason: %s", err)
//}
//
//privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
//keyBuff := new(bytes.Buffer)
//if err := pem.Encode(keyBuff, privateKeyPEM); err != nil {
//log.Panicf("PrivateKeyFile content failed reason: %s", err)
//}
//pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
//if err != nil {
//log.Panicf("PrivateKeyFile content failed reason: %s", err)
//}
//pubKey := fmt.Sprintf("%s %s \n", strings.TrimSuffix(string(ssh.MarshalAuthorizedKey(pub)), "\n"), keySignature)
//
//content := map[string][]byte{privateKeyName: []byte(keyBuff.Bytes()), pubKeyName: []byte(pubKey)}
//
