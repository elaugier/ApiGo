package apigodb

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	//drivers databases
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Job ...
type Job struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Message   string
	Result    string
	Status    string
}

//TableName ...
func (Job) TableName() string {
	return "jobs"
}

//Request ...
type Request struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IP        string
	Method    string
	URI       string
	Browser   string
	Os        string
	Status    int
	Username  string
	Device    string
	Duration  float64
}

//TableName ...
func (Request) TableName() string {
	return "requests"
}

//User ...
type User struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Username  string
	Apikey    string
}

//TableName ...
func (User) TableName() string {
	return "users"
}

//InitDb ...
func InitDb() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}
