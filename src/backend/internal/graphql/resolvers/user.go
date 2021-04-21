package resolver

import (
	"errors"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/iftech-a/lookum/src/backend/internal/store"
	"github.com/sirupsen/logrus"
)

func User(s store.UserRepo) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		idStr, ok := p.Args["id"]
		if !ok {
			return nil, errors.New("user id parameter missing")
		}

		id, err := strconv.Atoi(idStr.(string))
		if err != nil {
			return nil, errors.New("malformed param")
		}

		return s.GetUser(id)
	}
}

func UserCreate(s store.UserRepo) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {

		user := model.NewUser()
		user.FirstName = p.Args["firstName"].(string)
		user.LastName = p.Args["lastName"].(string)
		user.MiddleName = p.Args["middleName"].(string)
		user.Email = p.Args["email"].(string)
		user.Intro = p.Args["intro"].(string)
		user.Mobile = p.Args["mobile"].(string)
		user.Admin = p.Args["isAdmin"].(bool)
		user.Vendor = p.Args["isVendor"].(bool)
		user.Profile = p.Args["profile"].(string)
		user.Password = p.Args["password"].(string)

		dbUser, err := s.GetUserByEmail(user.Email)
		if err != nil {
			logrus.Errorf("GetUserByEmail failed: %v", err.Error())
			return nil, errors.New("internal error")
		}

		if dbUser != nil {
			logrus.Errorf("user alread exists")
			return nil, errors.New("user_already_exists")
		}
		password, err := user.GeneratePasswordHash()
		if err != nil {
			logrus.Errorf("GeneratePasswordHash: %v", err.Error())
			return nil, errors.New("internal error")
		}

		user.Password = string(password)

		user.ID, err = s.Create(user)
		if err != nil {
			logrus.Errorf("user create failed: %v", err.Error())
			return nil, errors.New("internal error")
		}

		return user, nil
	}
}
