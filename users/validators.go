package users

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func ValidateUserForCreation(ctx *gin.Context) (*UserJson, error) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if user.DiscordId == "" && user.MinecraftNickname == "" {
		return nil, MissingMinecraftNicknameAndDiscordIdError
	}
	if user.Id != uuid.Nil {
		return nil, IdAlreadyPresentMessageError
	}
	if user.Verified == true {
		return nil, CannotCreateVerifiedUserError
	}
	user.MinecraftNickname = strings.ToLower(user.MinecraftNickname)
	return user, nil
}

func ValidateUserForUpdateById(ctx *gin.Context) (*UserJson, error) {
	user, err := ValidateUserForUpdate(ctx)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	if userId != user.Id {
		return nil, MismatchedIdError
	}

	return user, nil
}

func ValidateUserForUpdate(ctx *gin.Context) (*UserJson, error) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if user.DiscordId == "" && user.MinecraftNickname == "" {
		return nil, MissingMinecraftNicknameAndDiscordIdError
	}

	user.MinecraftNickname = strings.ToLower(user.MinecraftNickname)
	return user, nil
}

func getUserFromContext(ctx *gin.Context) (*UserJson, error) {
	var user UserJson
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return nil, InvalidJsonError
	}

	return &user, nil
}
