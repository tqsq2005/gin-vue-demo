package model

func initMigrate()  {
	//自动建表
	db.AutoMigrate(&User{})
}
