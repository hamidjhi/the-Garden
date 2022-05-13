package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func ShowTrees(date model.Date, treeId string, paginate *model.Paginate) (*model.PaginateTreeResponse, error) {
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
	response := MySQL.Table("trees").
		Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
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

func ShowTreesByQr(qr string) (*model.Tree, error) {
	var c *model.Tree
	var err error
	//res := MySQL.Model(model.Tree{}).Where(model.Tree{Qr: qr}).Joins("").Joins("").Find(&c)
	res := MySQL.Model(&model.Tree{}).Where(model.Tree{Qr: qr}).Select("qr, comment_id").Joins("left join comment on comment_id = tree_id").Scan(c)
	if res.Error != nil {
		return nil, err
	}
	return c, nil

}

func CreateTree(tree *model.Tree) error {
	response := MySQL.Model(model.Tree{}).Create(&tree)
	if response.Error != nil {
		return response.Error
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

func DeleteTree(id uint) (resp *model.Tree, err error) {
	response := MySQL.Table("trees").Where("id = ?", id).Delete(&id)
	if response.Error != nil {
		return nil, response.Error
	}
	return resp, nil
}

type result struct {
	Name  string
	Email string
}
