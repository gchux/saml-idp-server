package main

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"flag"

	"net/url"

	"github.com/zenazn/goji"
	"golang.org/x/crypto/bcrypt"

	"github.com/crewjam/saml/logger"
	"github.com/crewjam/saml/samlidp"
)

var key = func() crypto.PrivateKey {
	b, _ := pem.Decode([]byte(`-----BEGIN RSA PRIVATE KEY-----
	....
-----END RSA PRIVATE KEY-----`))
	k, _ := x509.ParsePKCS1PrivateKey(b.Bytes)
	return k
}()

var cert = func() *x509.Certificate {
	b, _ := pem.Decode([]byte(`-----BEGIN CERTIFICATE-----
	.....
-----END CERTIFICATE-----`))
	c, _ := x509.ParseCertificate(b.Bytes)
	return c
}()

func main() {
	logr := logger.DefaultLogger
 	// baseURLstr := flag.String("idp", "", "The URL to the IDP")
	flag.Parse()

	baseURL, err := url.Parse("https://saml-idp-dot-chux-testing.wn.r.appspot.com")
	if err != nil {
		logr.Fatalf("cannot parse base URL: %v", err)
	}

	idpServer, err := samlidp.New(samlidp.Options{
		URL:         *baseURL,
		Key:         key,
		Logger:      logr,
		Certificate: cert,
		Store:       &samlidp.MemoryStore{},
	})
	if err != nil {
		logr.Fatalf("%s", err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	err = idpServer.Store.Put("/users/test", samlidp.User{Name: "test",
		HashedPassword: hashedPassword,
		Groups:         []string{"Users"},
		Email:          "test@saml.app",
		CommonName:     "test",
		Surname:        "test",
		GivenName:      "test",
	})
	if err != nil {
		logr.Fatalf("%s", err)
	}

	goji.Handle("/*", idpServer)
	goji.Serve()
}
