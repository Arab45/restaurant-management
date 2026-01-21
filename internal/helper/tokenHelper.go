package helper

import(
	"RESTAURANT-MANAGEMENT/internal/database"
	"os"
	"log"
	"time"


	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

)

type SignedDetails struct{
	Email string
	First_name string
	Last_name string
	Uid string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GenerateAllToken( email string, firstName string, lastName string, uid string)(signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails
	Email: email,
	First_name: firstName,
	Last_name: lastName,
	Uid: uid,
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24).Unix())
	}

	refreshClaim := &SignedDetails{
		StandardClaims: {
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168).Unix())
		}
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim).SignedString([]byte(SECRET_KEY))
	if err != nil{
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func updateAllTokens(signedToken string, signedRefreshToken string, userId){
	var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		upsert: &upsert
	}

	_, err := userCollection.UpdateOne(
		ctx, 
		filter,
		bson.D{
			{"$set", updateObj}
		}
	)

	if err != nil {
		log.Panic(err)
		return
	}

}

func ValidateToken( signedToken string)( claims *signedToken, msg string){
	token, err := jwt.PersWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token)(interface{}, error){
			return []byte(SECRET_KEY), nil
		}
	)

	claim, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}


	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg = fmt.Sprintf("token is required")
		msg = err.Error()
		return
	}

	return claims, msg

	//the token is invalid
	//the token is expired
}