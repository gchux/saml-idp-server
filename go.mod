module saml-idp

go 1.18

replace github.com/crewjam/saml => ./saml

require (
	github.com/crewjam/saml v0.4.14-0.20230420111643-34930b26d33b
	github.com/zenazn/goji v1.0.1
	golang.org/x/crypto v0.0.0-20220128200615-198e4374d7ed
)

require (
	github.com/beevik/etree v1.1.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/russellhaering/goxmldsig v1.3.0 // indirect
)
