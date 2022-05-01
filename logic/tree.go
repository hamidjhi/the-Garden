package logic

import (
	"chemex/db"
	"chemex/model"
	"gorm.io/gorm"
	"log"
)

func ShowTrees(date model.Date, treeId string, page *model.Paginate, )(resp *model.PaginateTreeResponse,err error) {

	resp, err = db.ShowTrees(date, treeId, page)
	if err != nil {
		log.Println("we have an err to getting response from tree table")
	}

	return resp, nil
}

//func ShowByQr(qr string)(resp *model.Tree,err  error)  {
//
//	resp := db.ShowByQr(qr)
//
//
//}







func CreateTree(tree *model.Tree, qr string) error  {
 var err error
	var NewTree = model.Tree{
		Model:  gorm.Model{},
		Name:   tree.Name,
		TreeId: tree.TreeId,
		Qr:     qr,
		Lat:    tree.Lat,
		Long:   tree.Long,
		Garden: tree.Garden,
	}
	 err = db.CreateTree(&NewTree)
	if err != nil {
		log.Println("cannot add anything to tree table ")
	}
	return nil
}

func UpdateTree(tree *model.Tree, str string)( err error)  {
	 err = db.UpdateTree(tree, str)
	if err != nil {
		log.Println("cannot update tree table")
	}
	return  nil
}

func DeleteTree(treeId string)(resp *model.Tree, err error) {

	resp, err = db.DeleteTree(treeId)
	if err != nil {
		log.Println("cannot delete from tree table")
	}
	return resp, nil
}

