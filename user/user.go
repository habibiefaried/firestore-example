package user

import (
  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
  "golang.org/x/net/context"
  "cloud.google.com/go/firestore"
  "errors"
  //"google.golang.org/api/iterator"
)

type UserModel struct { // username as the path
	Fullname 	string  `firestore:"fullname"`
	Email 		string  `firestore:"email"`
	Password 	string  `firestore:"password"`
}

type UserLib struct {
	Ctx      context.Context
	Users 	 *firestore.CollectionRef
}

func InitUser(sdkLocation string, database string) (*UserLib, error) {
	var ul UserLib
	ul.Ctx = context.Background()
	
	sa := option.WithCredentialsFile("/home/habibie/sdk.json") // This should be parameterized
	app, err := firebase.NewApp(ul.Ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ul.Ctx)
	if err != nil {
		return nil, err
	}

	ul.Users = client.Collection(database) // This should be parameterized
	
	return &ul, nil
}

func (ul *UserLib) InsertOrUpdateData(username string, um UserModel) (error) {
	u := ul.Users.Doc("user/accounts/"+username)
	_, err := u.Set(ul.Ctx, um)
	return err
}

func (ul *UserLib) GetData(username string) (*UserModel, error) {
	u, err := ul.Users.Doc("user/accounts/"+username).Get(ul.Ctx)
	if (u.Exists()) {
		if (err != nil){
			return nil, err
		}

		var um UserModel
		d := u.Data()
		um.Fullname, _ = d["fullname"].(string)
		um.Email, _ = d["email"].(string)
		um.Password, _ = d["password"].(string)
		return &um, nil
	} else {
		return nil, errors.New("Not found")
	}
}