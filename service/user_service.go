package service

import (
	"hulio-user-service/constants/bizerror"
	"hulio-user-service/dao"
	"hulio-user-service/dao/model"
	"hulio-user-service/handler/request"
	"hulio-user-service/utils"
)

func Register(request *request.RegisterRequest) error {

	//check if username exist
	if user, err := dao.GetUserByUsername(request.Username); err != nil {
		return err
	} else {
		if user != nil {
			return bizerror.UsernameExisted
		}
	}

	//hash password
	hashedPwd, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Username: request.Username,
		Password: hashedPwd,
		Nickname: request.Nickname,
		Avatar:   request.Avatar,
		Email:    request.Email,
		Phone:    request.Phone,
		Extinfo:  request.Extinfo,
	}
	if err := dao.CreateUser(&user); err != nil {
		return err
	}
	return nil
}

func Login(request *request.LoginRequest) (*UserInfo, error) {
	username := request.Username
	password := request.Password
	userInfo, err := dao.GetUserByUsername(username)
	if err != nil {
		return nil, err
	} else if userInfo == nil {
		return nil, bizerror.UserNotExist
	}

	if !utils.CheckPassword(userInfo.Password, password) {
		return nil, bizerror.PasswordError
	}

	// username and password match, generate token
	token, err := utils.GenerateToken(userInfo.Username, userInfo.ID)
	if err != nil {
		return nil, err
	}

	userInfo.Password = "" //clear password
	loginRes := UserInfo{
		Token: token,
		User:  *userInfo,
	}

	return &loginRes, nil
}

func GetUserInfo(request *request.LoginRequest, claims *utils.UserClaims) (*UserInfo, error) {
	userId := claims.UserId
	userInfo, err := dao.GetUserByID(uint(userId))
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, bizerror.UserNotExist
	}
	userInfo.Password = ""
	return &UserInfo{
		Token: "",
		User:  *userInfo,
	}, nil
}

type UserInfo struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

// DeleteUser deletes the authenticated user identified by claims
func DeleteUser(claims *utils.UserClaims) error {
	userId := claims.UserId
	userInfo, err := dao.GetUserByID(uint(userId))
	if err != nil {
		return err
	}
	if userInfo == nil || userInfo.ID == 0 {
		return bizerror.UserNotExist
	}
	return dao.DeleteUser(uint(userId))
}

// UpdateUser updates authenticated user's info with non-empty fields
func UpdateUser(req *request.UpdateUserRequest, claims *utils.UserClaims) (*UserInfo, error) {
	userId := claims.UserId
	user, err := dao.GetUserByID(uint(userId))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, bizerror.UserNotExist
	}

	// Apply partial updates
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Extinfo != "" {
		user.Extinfo = req.Extinfo
	}

	if err := dao.UpdateUser(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return &UserInfo{Token: "", User: *user}, nil
}
