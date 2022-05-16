package request

import (
	"net/http"
	"time"

	"resource-service/src/model"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	logger "resource-service/utils/logging"
	"resource-service/utils/wrapper/response"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"
)

var usrdto = []model.User{
	{ID: 1, FirstName: "Deter", LastName: "Fir", Email: "ety@k1.co", Password: "poetry@", Role: "Dev"},
	{ID: 2, FirstName: "Peter", LastName: "Ge", Email: "frt@k1.co", Password: "yutr@12", Role: "Art"},
}

var log *zerolog.Logger = logger.GetInstance()

var ACCESS_SECRET = "jduqwrtgfuoaklu"

type Claims struct {
	UID int `json:"UserId"`
	jwt.StandardClaims
}

var claims = &Claims{}

var dt model.User

//Check JSON according to given structure
func CheckJSON(c *gin.Context, data interface{}) bool {
	if err := c.BindJSON(&data); err != nil {
		log.Error().Msgf(constants.WRONG_REQUEST_BODY + err.Error())
		//Return Error Invalid Request Body
		response.Send400(c, constants.INVALID_REQUEST_BODY)
		return true
	}
	return false
}

//Validate Data According to Request Body
func CheckRequestBodyValidator(c *gin.Context, data interface{}) bool {
	validation := validator.New()
	if err := validation.Struct(data); err != nil {
		log.Error().Msgf(constants.WRONG_REQUEST_BODY + err.Error())
		//Return Error Required Request Body
		response.Send400(c, constants.REQUIRED_REQUEST_BODY)
		return true
	}
	return false
}

//Authenticate and Authorise the user request with the provided credentials
func CheckAuthorisedUser(c *gin.Context, data dto.RequestUserDTO) (string, error) {

	// Encode the password
	// hash := sha512.New()
	// hash.Write([]byte(data.Email))
	// hash.Write([]byte(data.Password))
	// pwd := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	var Id int

	for _, u := range usrdto {
		if u.Email == data.Email && u.Password == data.Password {

			expirationTime := time.Now().Add(90 * time.Second)

			Id = u.ID

			//access the token with the claims used for signing
			token, err := AccessToken(Id, dt.Email, expirationTime)

			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return " ", err
			}

			usrdt := GetResData(token)

			return usrdt, nil
		}
	}

	return "Null User Data", nil

}

// Access token from the user details
func AccessToken(userId int, userEmail string, expTime time.Time) (string, error) {

	claims := Claims{
		UID: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    userEmail,
			ExpiresAt: expTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(ACCESS_SECRET))

	return token, err
}

// Get the User Data to response user
func GetResData(token string) string {

	return token
}

func GetUId(jwtoken string) (int, error) {

	token, err := jwt.ParseWithClaims(jwtoken, claims, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("There was an error")
		// }
		return []byte(ACCESS_SECRET), nil
	})

	if err != nil {
		return -1, err
	}
	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		// c.JSON(http.StatusOK, gin.H{"message": "success"})
		return claims.UID, nil
	} else {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return -1, err
	}

	// token, _, err := new(jwt.Parser).ParseUnverified(jwtoken, jwt.StandardClaims{})
	// if err != nil {
	// 	return " ", err
	// }

	// if claims, ok := token.Claims.(jwt.StandardClaims); ok {
	// 	UId = fmt.Sprintf("%v", claims.Issuer)
	// }
}
