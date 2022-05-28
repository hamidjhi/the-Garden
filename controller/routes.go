package controller

import (
	"github.com/labstack/echo"
)

func RegisterRoutes(c *echo.Echo) {
	v := c.Group("/v")

	// garden crud
	v.GET("/showGarden", showGardens)
	v.POST("/createGarden", createGarden)
	v.PUT("/updateGarden", updateGarden)
	v.DELETE("/deleteGarden", deleteGarden)

	// tree crud
	v.GET("/showTrees", showTrees)
	v.POST("/createTree", createTree)
	v.PUT("/updateTree", updateTree)
	v.DELETE("/deleteTree", deleteTree)

	// comment crud
	v.GET("/getComment", showComments)
	v.POST("/postComment", createComment)
	v.PUT("/updateComment", updateComment)
	v.DELETE("/deleteComment", deleteComment)

	// user qr scanner
	v.GET("/getTreeByQr", showTreesByQr)
	v.GET("/getGardenByNumber", showGardenByNumber)


	// tags crud

	v.GET("/getTags", showTags)
	v.POST("/postTag", addTag)
	v.PUT("/updateTag", updateTag)
	v.DELETE("/deleteTag", deleteTag)



	//User crud

	v.GET("/getUser", getUser)
	v.POST("/createUser", createUser)
	v.PUT("/updateUser", updateUser)

}


