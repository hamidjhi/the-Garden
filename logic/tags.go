package logic

import (
	"chemex/db"
	"chemex/model"
	"log"
)

func ShowTags(date model.Date, tagId string, page *model.Paginate) (resp *model.PaginateTagsResponse, err error) {

	resp, err = db.ShowTags(date, tagId, page)
	if err != nil {
		log.Println("we have an err to getting response from tags table")
	}

	return resp, nil
}

func AddTag(tags *model.Tags) (err error) {
	err = db.AddTag(tags)
	if err != nil {
		log.Println("someThing wrong to add tag, check it!!")
	}
	return nil
}
func UpdateTag(tag *model.Tags, id string) (err error) {
	err = db.UpdateTag(tag, id)
	if err != nil {
		log.Println("cannot update tags....")
	}
	return nil
}

func DeleteTag(id string) (err error) {
	err = db.DeleteTag(id)
	if err != nil {
		log.Println("cannot delete tag ....")
	}
	return nil
}
