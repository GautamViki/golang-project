package initializer

import "jwk-token/model"

func SyncDatabase() {
	DB.AutoMigrate(&model.User{})
}
