package controller

import (
	"chemex/auth"
	"github.com/labstack/echo"
)

func RegisterRoutes(c *echo.Echo) {
	v := c.Group("/v")

	// garden crud
	v.GET("/showGarden", showGardens, auth.Allow(showGardens))
	v.POST("/createGarden", createGarden, auth.Allow(createGarden))
	v.PUT("/updateGarden", updateGarden, auth.Allow(updateGarden))
	v.DELETE("/deleteGarden", deleteGarden, auth.Allow(deleteGarden))

	// tree crud
	v.GET("/showTrees", showTrees, auth.Allow(showTrees))
	v.POST("/createTree", createTree, auth.Allow(createTree))
	v.PUT("/updateTree", updateTree, auth.Allow(updateTree))
	v.DELETE("/deleteTree", deleteTree, auth.Allow(deleteTree))

	// comment crud
	v.GET("/getComment", showComments, auth.Allow(showComments))
	v.POST("/postComment", createComment, auth.Allow(createComment))
	v.PUT("/updateComment", updateComment, auth.Allow(updateComment))
	v.DELETE("/deleteComment", deleteComment, auth.Allow(deleteComment))

	// user qr scanner
	v.GET("/getTreeByQr", showTreesByQr, auth.Allow(showTreesByQr))
	v.GET("/getGardenByNumber", showGardenByNumber, auth.Allow(showGardenByNumber))


	// tags crud

	v.GET("/getTags", showTags, auth.Allow(showTags))
	v.POST("/postTag", addTag, auth.Allow(addTag))
	v.PUT("/updateTag", updateTag, auth.Allow(updateTag))
	v.DELETE("/deleteTag", deleteTag, auth.Allow(deleteTag))



	//User crud
	v.GET("/createUser", createUser, auth.Allow(createUser))
	v.GET("/updateUser", updateUser, auth.Allow(updateUser))

}


