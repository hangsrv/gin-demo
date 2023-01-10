package model

// 业主信息
type Owner struct {
	ID    uint64 `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name  string `gorm:"type:varchar(10);NOT NULL" json:"name" form:"name"`
	Work  string `gorm:"type:varchar(50);NOT NULL" json:"work" form:"work"`
	Phone string `gorm:"type:varchar(11);NOT NULL" json:"phone" form:"phone"`
}

func (m *Owner) TableName() string {
	return "owner"
}

// 房屋信息
type Room struct {
	ID        uint64 `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	OwnerName string `gorm:"type:varchar(10);NOT NULL" json:"owner_name" form:"owner_name"`
	Area      uint64 `gorm:"type:int(10);NOT NULL" json:"area" form:"area"`
	Number    string `gorm:"type:varchar(10);;NOT NULL" json:"number" form:"number"`
}

func (m *Room) TableName() string {
	return "room"
}

// 房屋费用
type RoomInfo struct {
	ID          uint64  `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Number      string  `gorm:"type:varchar(10);NOT NULL" json:"number" form:"number"`
	Water       float64 `gorm:"type:float(10,2);NOT NULL" json:"water" form:"water"`
	Electricity float64 `gorm:"type:float(10,2);NOT NULL" json:"electricity" form:"electricity"`
	Year        uint64  `gorm:"type:int(11);NOT NULL" json:"year" form:"year"`
	Month       uint64  `gorm:"type:int(11);NOT NULL" json:"month" form:"month"`
	Fee         float64 `gorm:"type:int(11)" json:"fee"`
}

func (m *RoomInfo) TableName() string {
	return "room_info"
}

// 部门信息
type Department struct {
	ID          uint64 `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name        string `gorm:"type:varchar(10);NOT NULL" json:"name" form:"name"`
	Phone       string `gorm:"type:varchar(11);NOT NULL" json:"phone" form:"phone"`
	ManagerName string `gorm:"type:varchar(11);NOT NULL" json:"manager_name" form:"manager_name"`
}

func (m *Department) TableName() string {
	return "department"
}

// 员工信息
type Employee struct {
	ID             uint64 `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	DepartmentName string `gorm:"type:varchar(10);NOT NULL" json:"department_name" form:"department_name"`
	Username       string `gorm:"type:varchar(10);NOT NULL" json:"username" form:"username"`
	Password       string `gorm:"type:varchar(10);NOT NULL" json:"password" form:"password"`
	Sex            string `gorm:"type:varchar(10);NOT NULL" json:"sex" form:"sex"`
	Phone          string `gorm:"type:varchar(11);NOT NULL" json:"phone" form:"phone"`
	IsManger       uint64 `json:"is_manger" sql:"-"`
}

func (m *Employee) TableName() string {
	return "employee"
}
