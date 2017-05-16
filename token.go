package philote

import(
  "github.com/dgrijalva/jwt-go"
)

func NewToken(secret string, read, write []string) (string, error) {
  auth := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "read": read,
      "write": write,
  })

  token, err := auth.SignedString([]byte(secret)); if err != nil {
    return "", err
  }

  return token, nil
}
