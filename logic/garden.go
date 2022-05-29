package logic

import (
	"chemex/db"
	"chemex/model"
	"log"

)

func ShowGardens(date model.Date, GardenId string,userId,gardenName  string, page *model.Paginate, )(resp *model.PaginateGardenResponse,err error) {

	resp, err = db.ShowGardens(date, GardenId,userId, gardenName, page)
	if err != nil {
		log.Println("we have an err to getting response from garden table")
	}

	return resp, nil
}



func CreateGarden(garden *model.Garden)( err error)  {
	  err = db.CreateGarden(garden)
	if err != nil {
		log.Println("cannot add anything to garden table ")
	}
	return nil
}

func UpdateGarden(garden *model.Garden, id string)(err error)  {
	err = db.UpdateGarden(garden, id)
	if err != nil {
		log.Println("cannot update garden db table")
	}
	return nil
}

func DeleteGarden(gardenId string)(err error) {

	err = db.DeleteGarden(gardenId)
	if err != nil {
		log.Println("cannot delete row from garden database ")
	}
	return nil
}




