package logic

import (
	"chemex/db"
	"chemex/model"
	"log"

)

func ShowGardens(date model.Date, GardenId string, page *model.Paginate, )(resp *model.PaginateGardenResponse,err error) {

	resp, err = db.ShowGardens(date, GardenId, page)
	if err != nil {
		log.Println("we have an err to getting response from garden table")
	}

	return resp, nil
}

func CreateGarden(garden *model.Garden)(resp *model.Garden, err error)  {
	 resp, err = db.CreateGarden(garden)
	if err != nil {
		log.Println("cannot add anything to garden table ")
	}
	return resp, nil
}

func UpdateGarden(garden *model.Garden, id string)(resp *model.Garden, err error)  {
	resp, err = db.UpdateGarden(garden, id)
	if err != nil {
		log.Println("cannot update garden db table")
	}
	return resp, nil
}

func DeleteGarden(gardenId string)(resp *model.Garden, err error) {

	resp, err = db.DeleteGarden(gardenId)
	if err != nil {
		log.Println("cannot ")
	}
	return resp, nil
}




