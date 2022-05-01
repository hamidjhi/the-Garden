package controller

import "github.com/labstack/echo"

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
	v.PUT("/UpdateComment", updateComment)
	v.DELETE("/deleteComment", deleteComment)

	// user qr scanner
	v.GET("/getTreeByQr", showTreesByQr)


}


