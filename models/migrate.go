package models

import (
	"fmt"
	"mobile-rest-api/db"
)

//Migrate db
func Migrate(){
	//connect To DB
   db.Init()
   var db = db.GetDB()
   defer db.Close()
   //disconnect To DB
   db.AutoMigrate(
      &Pages{},
      &Banner{},
      &Article{},
      &Feedback{},
   )
   fmt.Println("\t Auto migration complated successfully !!!!\t")
}