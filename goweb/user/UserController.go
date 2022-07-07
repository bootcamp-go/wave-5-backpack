package user

import "github.com/gin-gonic/gin"

func validateToken(c *gin.Context) bool {
	if token := c.GetHeader("token"); token != "123" {
		c.JSON(401, gin.H{
			"error": "Token Invalido",
		})
		return false
	}
	return true
}
func CreateUser(c *gin.Context) {
	var userReq CreateUserRequest
	var user = User{}
	if validateToken(c) {
		if err := c.ShouldBindJSON(&userReq); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		userReq.Id = 3
		user.createUser(userReq)
		c.JSON(200, userReq)
	}
}
