package op

import (
	"time"

	"codeberg.org/alist/alist/v3/internal/db"
	"codeberg.org/alist/alist/v3/internal/errs"
	"codeberg.org/alist/alist/v3/internal/model"
	"codeberg.org/alist/alist/v3/pkg/singleflight"
	"codeberg.org/alist/alist/v3/pkg/utils"
	"codeberg.org/alist/go-cache"
)

var userCache = cache.NewMemCache(cache.WithShards[*model.User](2))
var userG singleflight.Group[*model.User]
var guestUser *model.User
var adminUser *model.User

func GetAdmin() (*model.User, error) {
	if adminUser == nil {
		user, err := db.GetUserByRole(model.ADMIN)
		if err != nil {
			return nil, err
		}
		adminUser = user
	}
	return adminUser, nil
}

func GetGuest() (*model.User, error) {
	if guestUser == nil {
		user, err := db.GetUserByRole(model.GUEST)
		if err != nil {
			return nil, err
		}
		guestUser = user
	}
	return guestUser, nil
}

func GetUserByRole(role int) (*model.User, error) {
	return db.GetUserByRole(role)
}

func GetUserByName(username string) (*model.User, error) {
	if username == "" {
		return nil, errs.EmptyUsername
	}
	if user, ok := userCache.Get(username); ok {
		return user, nil
	}
	user, err, _ := userG.Do(username, func() (*model.User, error) {
		_user, err := db.GetUserByName(username)
		if err != nil {
			return nil, err
		}
		userCache.Set(username, _user, cache.WithEx[*model.User](time.Hour))
		return _user, nil
	})
	return user, err
}

func GetUserById(id uint) (*model.User, error) {
	return db.GetUserById(id)
}

func GetUsers(pageIndex, pageSize int) (users []model.User, count int64, err error) {
	return db.GetUsers(pageIndex, pageSize)
}

func CreateUser(u *model.User) error {
	u.BasePath = utils.FixAndCleanPath(u.BasePath)
	return db.CreateUser(u)
}

func DeleteUserById(id uint) error {
	old, err := db.GetUserById(id)
	if err != nil {
		return err
	}
	if old.IsAdmin() || old.IsGuest() {
		return errs.DeleteAdminOrGuest
	}
	userCache.Del(old.Username)
	return db.DeleteUserById(id)
}

func UpdateUser(u *model.User) error {
	old, err := db.GetUserById(u.ID)
	if err != nil {
		return err
	}
	if u.IsAdmin() {
		adminUser = nil
	}
	if u.IsGuest() {
		guestUser = nil
	}
	userCache.Del(old.Username)
	u.BasePath = utils.FixAndCleanPath(u.BasePath)
	return db.UpdateUser(u)
}

func Cancel2FAByUser(u *model.User) error {
	u.OtpSecret = ""
	return UpdateUser(u)
}

func Cancel2FAById(id uint) error {
	user, err := db.GetUserById(id)
	if err != nil {
		return err
	}
	return Cancel2FAByUser(user)
}

func DelUserCache(username string) error {
	user, err := GetUserByName(username)
	if err != nil {
		return err
	}
	if user.IsAdmin() {
		adminUser = nil
	}
	if user.IsGuest() {
		guestUser = nil
	}
	userCache.Del(username)
	return nil
}
