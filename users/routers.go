package users

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
)

func (self *UsersRouter) UsersRegister(router *gin.RouterGroup) {
	router.GET("/", self.getUsers)
	router.POST("/", self.createUser)
	router.GET("/:id", self.getUserById)
	router.PUT("/:id", self.updateUser)
	router.GET("/byMinecraftNickname/:minecraftNickname", self.getUserByMinecraftNickname)
	router.PUT("/byMinecraftNickname/:minecraftNickname", self.updateUserByNickname)
	router.PUT("/byMinecraftNickname/:minecraftNickname/status", self.updateUserStatusByMinecraftNickname)
	router.GET("/byDiscordId/:discordId", self.getUserByDiscordId)
	router.PUT("/byDiscordId/:discordId", self.updateUserByDiscordId)
	router.PUT("/byDiscordId/:discordId/status", self.updateUserStatusByDiscordId)
}

type UsersRouter struct {
	service *Service
}

func Router(db *gorm.DB) UsersRouter {
	return UsersRouter{&Service{db}}
}

func (self *UsersRouter) getUsers(ctx *gin.Context) {
	users, err := self.service.GetUsers()
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, users)
}

func (self *UsersRouter) createUser(ctx *gin.Context) {
	user, err := ValidateUserForCreation(ctx)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	createdUser, err := self.service.CreateUser(user)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, createdUser)
}

func (self *UsersRouter) updateUser(ctx *gin.Context) {
	user, err := ValidateUserForUpdateById(ctx)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	updatedUser, err := self.service.UpdateUser(user)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, updatedUser)
}

func (self *UsersRouter) getUserById(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	user, err := self.service.GetUserById(id)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}

func (self *UsersRouter) getUserByMinecraftNickname(ctx *gin.Context) {
	nickname := ctx.Param("minecraftNickname")
	user, err := self.service.GetUserByMinecraftNickname(nickname)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}

func (self *UsersRouter) updateUserByNickname(ctx *gin.Context) {
	nickname := ctx.Param("minecraftNickname")
	user, err := ValidateUserForUpdate(ctx)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	updatedUser, err := self.service.UpdateUserByMinecraftNickname(nickname, user)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedUser)
}

func (self *UsersRouter) getUserByDiscordId(ctx *gin.Context) {
	discordId := ctx.Param("discordId")
	user, err := self.service.GetUserByDiscordId(discordId)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)

}

func (self *UsersRouter) updateUserByDiscordId(ctx *gin.Context) {
	discordId := ctx.Param("discordId")
	user, err := ValidateUserForUpdate(ctx)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	updatedUser, err := self.service.UpdateUserByDiscordId(discordId, user)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedUser)

}

func (self *UsersRouter) updateUserStatusByMinecraftNickname(ctx *gin.Context) {
	nickname := ctx.Param("minecraftNickname")
	var status Status
	err := ctx.BindJSON(&status)
	if err != nil {
		handleMcdcError(ctx, InvalidJsonError)
		return
	}

	updatedUser, err := self.service.UpdateStatusByMinecraftNickname(nickname, status)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedUser)
}

func (self *UsersRouter) updateUserStatusByDiscordId(ctx *gin.Context) {
	discordId := ctx.Param("discordId")
	var status Status
	err := ctx.BindJSON(&status)
	if err != nil {
		handleMcdcError(ctx, InvalidJsonError)
		return
	}

	updatedUser, err := self.service.UpdateStatusByDiscordId(discordId, status)
	if err != nil {
		handleMcdcError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedUser)
}

func handleMcdcError(ctx *gin.Context, err error) {
	mcdcError := CreateMcdcError(err)
	ctx.IndentedJSON(mcdcError.HttpStatus, mcdcError)
}
