package logic

import (
	"chemex/db"
	"chemex/model"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func ShowTrees(date model.Date, treeId string, page *model.Paginate, )(resp *model.PaginateTreeResponse,err error) {

	resp, err = db.ShowTrees(date, treeId, page)
	if err != nil {
		log.Println("we have an err to getting response from tree table")
	}

	return resp, nil
}

func ShowTreesByQr(qr string)(resp []*model.Result, err error)  {
	 resp, err = db.ShowTreesByQr(qr)
	if err != nil {
		log.Println("we dont have response from db")
	}
	return resp, nil
}



func CreateTree(tree *model.Tree, qr string) error  {
 var err error
	var NewTree = model.Tree{
		Model:       gorm.Model{},
		FullName:    tree.FullName,
		Age:         tree.Age,
		DateOfBirth: tree.DateOfBirth,
		Type:        tree.Type,
		Lat:         tree.Lat,
		Long:        tree.Long,
		Qr:          qr,
		Length:      tree.Length,
		CommentId: tree.CommentId,
		GardenId: tree.GardenId,
	}
	 err = db.CreateTree(&NewTree)
	if err != nil {
		log.Println("cannot add anything to tree table ")
	}
	return nil
}

func UpdateTree(tree *model.Tree, str string)( err error)  {
	v , err:= strconv.Atoi(str)
	if err != nil {
		log.Println("cannot convert")
	}
	  id := uint(v)
	 err = db.UpdateTree(tree, id)
	if err != nil {
		log.Println("cannot update tree table")
	}
	return  nil
}

func DeleteTree(treeId string)(err error) {
	v , err:= strconv.ParseUint(treeId, 64,10)
	id := uint(v)
	 err = db.DeleteTree(id)
	if err != nil {
		log.Println("cannot delete from tree table")
	}
	return nil
}

