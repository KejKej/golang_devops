package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

func OAuth2Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")

		log.Println("RECEIVED ACCESS TOKEN: " + token)

		var toValidate = map[string]string{
			"aud": os.Getenv("OKTA_AUDIENCE"),
			"cid": os.Getenv("OKTA_CLIENT_ID"),
		}
		log.Println(toValidate["aud"])
		log.Println(toValidate["cid"])
		log.Println(os.Getenv("OKTA_ISSUER_URI"))

		if strings.HasPrefix(token, "Bearer ") {

			token = strings.TrimPrefix(token, "Bearer ")

			verifierSetup := jwtverifier.JwtVerifier{
				Issuer:           os.Getenv("OKTA_ISSUER_URI"),
				ClaimsToValidate: toValidate,
			}
			verifier := verifierSetup.New()
			_, err := verifier.VerifyAccessToken(token)

			if err != nil {
				log.Printf(err.Error())
				context.JSON(403, gin.H{"error": err.Error()})
				context.Abort()
				return
			}

		} else {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
		}

		context.Next()
	}
}
