package model

type Response struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"msg" xml:"msg"`
}

type (
	User struct {
		ID        int    `json:"id" gorm:"column:id"`
		Birthday  string `json:"Birthday" gorm:"column:birthday"`
		Name      string `json:"name" gorm:"column:name"`
		Gender    string `json:"gender" gorm:"column:gender"`
		Phone     string `json:"phone" gorm:"column:phone"`
		Email     string `json:"email" gorm:"column:email"`
		Leave     int    `json:"Leave" gorm:"column:leave"`
		CreatedAt string `json:"createdAt" gorm:"column:createdAt"`
		UpdatedAt string `json:"updatedAt" gorm:"column:updatedAt"`
	}

	Users struct {
		Data []User `json:"data"`
	}
)

func (User) TableName() string {
	return "user"
}
