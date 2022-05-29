package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func ShowTrees(date model.Date, treeId string, gardenId string, paginate *model.Paginate) (*model.PaginateTreeResponse, error) {
	var all model.PaginateTreeResponse
	var err error

	if treeId != "" {

		v, err := strconv.Atoi(treeId)
		if err != nil {
			log.Println("cannot convert")
		}
		id := uint(v)

		response := MySQL.Model(&model.Tree{}).
			Where(&model.Tree{Model: gorm.Model{ID: id}}).
			Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
		if response.Error != nil {
			return nil, response.Error
		}
		p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
		p.SetPage(paginate.Page)

		if err := p.Results(&all.Resp); err != nil {
			return nil, err
		}

		all.Paging.Next, _ = p.NextPage() // 8
		all.Paging.Perv, _ = p.PrevPage() // 6
		all.Paging.Current = p.Page()     // 7
		totalPage := p.Nums()             // 7
		all.Paging.Total = totalPage
		return &all, nil
	}


	if gardenId != "" {

		v, err := strconv.Atoi(gardenId)
		if err != nil {
			log.Println("cannot convert")
		}
		id := uint(v)

		response := MySQL.Model(&model.Tree{}).
			Where(&model.Tree{Model: gorm.Model{ID: id}}).
			Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
		if response.Error != nil {
			return nil, response.Error
		}
		p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
		p.SetPage(paginate.Page)

		if err := p.Results(&all.Resp); err != nil {
			return nil, err
		}

		all.Paging.Next, _ = p.NextPage() // 8
		all.Paging.Perv, _ = p.PrevPage() // 6
		all.Paging.Current = p.Page()     // 7
		totalPage := p.Nums()             // 7
		all.Paging.Total = totalPage
		return &all, nil
	}

	response := MySQL.Table("trees").
		Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate).Order("id desc")
	if response == nil {
		return nil, err
	}
	p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
	p.SetPage(paginate.Page)

	if err = p.Results(&all.Resp); err != nil {
		return nil, err
	}

	all.Paging.Next, _ = p.NextPage() // 8
	all.Paging.Perv, _ = p.PrevPage() // 6
	all.Paging.Current = p.Page()     // 7
	totalPage := p.Nums()             // 7
	all.Paging.Total = totalPage
	return &all, nil

}

func ShowTreesByQr(qr string, paginate *model.Paginate) (*model.PaginateTreeResponse, error) {
	var c model.PaginateTreeResponse
	var err error

//	res :=MySQL.Table("trees").Select("trees.id, comments.text").Joins("LEFT JOIN comments on comments.tree_id = trees.id").Scan(&c)
//res := MySQL.Table("trees").Preload(model.Comment{TreeId: 1}.Text).Scan(&c)
//     res := MySQL.Model(&model.Tree{}).Where(model.Tree{Qr: qr}).Select("trees.id, trees.full_name, trees.lat, comments.text").Joins("left join comments on comments.tree_id = trees.id").Scan(&c)
//	    res :=MySQL.Model(&model.Tree{}).Where(model.Tree{Qr: qr}).
//	    	Select("trees.id, trees.full_name, trees.lat, comments.text, gardens.name").
//	    	Joins("left join comments on comments.tree_id = trees.id").
//	    	Joins("left join gardens on gardens.tree_id = trees.id").
//	    	Scan(&c)
//
	    res := MySQL.Model(&model.Tree{}).Where(model.Tree{Qr: qr}).Preload("Comments")


	if res == nil {

		return nil, err
	}
	p:= paginator.New(adapter.NewGORMAdapter(res),paginate.PerPage)
	    p.SetPage(paginate.Page)

	//for _, v := range c{
		if err = p.Results(&c.Resp) ;err != nil {
			return nil, err
		}
		c.Paging.Next, _ = p.NextPage() // 8
		c.Paging.Perv, _ = p.PrevPage() // 6
		c.Paging.Current = p.Page()     // 7
		totalPage := p.Nums()             // 7
		c.Paging.Total = totalPage
		//c= append(c,v)
		//}
		return &c, nil
}

func CreateTree(tree *model.Tree, userId string)  (err error) {
	ids, err := strconv.Atoi(userId)
	id := uint(ids)

	response := MySQL.Model(model.Tree{}).Where(&model.User{Model: gorm.Model{ID: id}}).Create(&tree)
	if response.Error != nil {
		return err
	}
	return nil
}

func UpdateTree(tree *model.Tree, id uint) (err error) {
	response := MySQL.Model(model.Tree{}).Where(&model.Tree{Model: gorm.Model{ID: id}}).Updates(&tree)
	if response.Error != nil {
		return err
	}
	return nil
}

func DeleteTree(id uint) ( err error) {
	response := MySQL.Table("trees").Where("id = ?", id).Delete(&id)
	if response.Error != nil {
		return err
	}
	return  nil
}



