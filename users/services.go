package users

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Service struct {
	db *gorm.DB
}

func (self *Service) GetUsers() ([]User, error) {
	var users []User
	result := self.db.Model(&User{}).Find(&users)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (self *Service) GetUserById(id uuid.UUID) (*User, error) {
	user, err := self.findUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (self *Service) CreateUser(user *User) (*User, error) {
	user.Id = uuid.NewV1()
	result := self.db.Model(&User{}).Create(user)
	err := result.Error
	if err != nil {
		return nil, handleDuplicateError(err)
	}

	createdUser, err := self.findUserById(user.Id)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (self *Service) UpdateUser(user *User) (*User, error) {
	_, err := self.findUserById(user.Id)
	if err != nil {
		return nil, err
	}

	updatedUser, err := self.justUpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (self *Service) GetUserByMinecraftNickname(nickname string) (*User, error) {
	var user User
	result := self.userModel().Where("minecraft_nickname = ?", nickname).First(&user)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundError
		}
		return nil, err
	}

	err = checkIfAnythingReturned(result)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (self *Service) UpdateUserByMinecraftNickname(minecraftNickname string, user *User) (*User, error) {
	existingUser, err := self.GetUserByMinecraftNickname(minecraftNickname)
	if err != nil {
		return nil, err
	}
	if existingUser.Id != user.Id {
		return nil, MismatchedIdError
	}

	updatedUser, err := self.justUpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (self *Service) UpdateStatusByMinecraftNickname(minecraftNickname string, status Status) (*User, error) {
	existingUser, err := self.GetUserByMinecraftNickname(minecraftNickname)
	if err != nil {
		return nil, err
	}
	existingUser.Verified = status.Verified

	updatedUser, err := self.justUpdateUser(existingUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (self *Service) GetUserByDiscordId(discordId string) (*User, error) {
	var user User
	result := self.userModel().Where("discord_id = ?", discordId).First(&user)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundError
		}
		return nil, err
	}

	err = checkIfAnythingReturned(result)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (self *Service) UpdateUserByDiscordId(discordId string, user *User) (*User, error) {
	existingUser, err := self.GetUserByDiscordId(discordId)
	if err != nil {
		return nil, err
	}
	if existingUser.Id != user.Id {
		return nil, MismatchedIdError
	}

	updatedUser, err := self.justUpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (self *Service) UpdateStatusByDiscordId(discordId string, status Status) (*User, error) {
	existingUser, err := self.GetUserByDiscordId(discordId)
	if err != nil {
		return nil, err
	}
	existingUser.Verified = status.Verified

	updatedUser, err := self.justUpdateUser(existingUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (self *Service) userModel() *gorm.DB {
	return self.db.Model(&User{})
}

func (self *Service) findUserById(userId uuid.UUID) (*User, error) {
	var user User
	result := self.userModel().Find(&user, userId)
	err := result.Error
	if err != nil {
		return nil, err
	}
	err = checkIfAnythingReturned(result)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func handleDuplicateError(err error) error {
	msg := err.Error()
	if strings.Contains(msg, "Error 1062") {
		if strings.Contains(msg, "minecraft_nickname") {
			return MinecraftNicknameAlreadyExists
		}
		if strings.Contains(msg, "discord_id") {
			return DiscordIdAlreadyExists
		}
	}
	return err
}

func checkIfAnythingReturned(tx *gorm.DB) error {
	if tx.RowsAffected == 0 {
		return UserNotFoundError
	}
	return nil
}

func (self *Service) justUpdateUser(user *User) (*User, error) {
	var updatedUser User
	result := self.userModel().Where("ID = ?", user.Id).Updates(user).First(&updatedUser)
	err := result.Error
	if err != nil {
		return nil, handleDuplicateError(err)
	}

	err = checkIfAnythingReturned(result)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
