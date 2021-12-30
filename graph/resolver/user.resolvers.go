package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"main/database"
	"main/graph/model"
	"main/middleware"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", input.Username).First(&user).Error
	if err == nil {
		return nil, errors.New("username already exists")
	}

	err = database.DB.Where("email = ?", input.Email).First(&user).Error
	if err == nil {
		return nil, errors.New("email already exists")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user = model.User{
		Username:  input.Username,
		Password:  string(password),
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	database.DB.Create(&user)

	var token string
	// token, err = middleware.CreateTokenGraph(user.UserID)
	// if err != nil {
	// 	return nil, errors.New("error creating access token")
	// }

	// TODO: Add token to response
	print(token)
	return &user, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id *string) (*model.User, error) {
	gc, err := middleware.GinContextFromContext(ctx)
	userId, exist := gc.Get("user_id")
	if !exist {
		return nil, err
	}

	var user model.User
	err = database.DB.Where("user_id = ?", userId).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (r *queryResolver) Login(ctx context.Context, username string, password string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	var token string
	// token, err = middleware.CreateTokenGraph(user.UserID)
	// if err != nil {
	// 	return nil, errors.New("error creating access token")
	// }

	// TODO: Add token to response
	print(token)
	return &user, nil
}

func (r *queryResolver) Logout(ctx context.Context) (*bool, error) {
	res := true
	return &res, nil
}
