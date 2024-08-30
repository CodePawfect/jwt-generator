package pkg

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateUnsignedJwt(companyTag string, purchaseOrganisation string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodNone,
		jwt.MapClaims{
			"userName":             "M37272",
			"idpId":                "M37272",
			"uuid":                 "ea65b61e-3bde-11ec-8d3d-0242ac130003",
			"email":                "test@smexnet.de",
			"iss":                  "smexnet-gateway",
			"sub":                  "M37272",
			"purchaseOrganisation": purchaseOrganisation,
			"companyTags":          []string{companyTag},
			"roles":                []string{role},
		})

	jwtString, err := token.SigningString()
	if err != nil {
		return "", err
	}

	return jwtString + ".", nil
}
