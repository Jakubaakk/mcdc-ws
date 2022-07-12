package users

import (
	"errors"
	"net/http"
)

func CreateMcdcError(err error) McdcError {
	if errors.Is(err, MissingMinecraftNicknameAndDiscordIdError) {
		return McdcError{"MISSING_MINECRAFT_NICKNAME_AND_DISCORD_ID", err.Error(), http.StatusBadRequest}
	}
	if errors.Is(err, IdAlreadyPresentMessageError) {
		return McdcError{"ID_ALREADY_PRESENT", err.Error(), http.StatusBadRequest}
	}
	if errors.Is(err, CannotCreateVerifiedUserError) {
		return McdcError{"CANNOT_CREATE_VERIFIED_USER", err.Error(), http.StatusBadRequest}
	}
	if errors.Is(err, MismatchedIdError) {
		return McdcError{"MISMATCHED_ID", err.Error(), http.StatusBadRequest}
	}
	if errors.Is(err, UserNotFoundError) {
		return McdcError{"USER_NOT_FOUND", err.Error(), http.StatusNotFound}
	}
	if errors.Is(err, MinecraftNicknameAlreadyExists) {
		return McdcError{"MINECRAFT_NICKNAME_ALREADY_EXISTS", err.Error(), http.StatusConflict}
	}
	if errors.Is(err, DiscordIdAlreadyExists) {
		return McdcError{"DISCORD_ID_ALREADY_EXISTS", err.Error(), http.StatusConflict}
	}
	if errors.Is(err, MissingApiKey) {
		return McdcError{"MISSING_API_KEY", err.Error(), http.StatusForbidden}
	}
	if errors.Is(err, InvalidApiKey) {
		return McdcError{"INVALID_API_KEY", err.Error(), http.StatusForbidden}
	}
	return McdcError{"INTERNAL_ERROR", err.Error(), http.StatusInternalServerError}
}

var MissingMinecraftNicknameAndDiscordIdError = errors.New("discordId or minecraftNickname has to be present")
var IdAlreadyPresentMessageError = errors.New("id cannot be present when creating user")
var InvalidJsonError = errors.New("provided json was invalid")
var CannotCreateVerifiedUserError = errors.New("user has to be created as unverified")
var MismatchedIdError = errors.New("id in URL has to match id in body")
var UserNotFoundError = errors.New("user was not found")
var MinecraftNicknameAlreadyExists = errors.New("user with provided minecraftNickname already exists")
var DiscordIdAlreadyExists = errors.New("user with provided discordId already exists")
var MissingApiKey = errors.New("missing api-key header")
var InvalidApiKey = errors.New("provided api-key is not valid")
