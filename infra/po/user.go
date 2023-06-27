package po

type User struct {
	Id   int64  `gorm:"column:id;type:bigint(21);auto_increment;primary_key;not null;comment:'用户id'"`
	Name string `gorm:"column:name;type:varchar(64);not null;comment:'用户名字'"`
	Sex  int    `gorm:"column:sex;type:tinyint(1);not null;comment:'用户性别'"`
	Age  int    `gorm:"column:age;type:int(3);not null;comment:'年龄'"`
	Home string `gorm:"column:home;type:varchar(255);not null;comment:'住址'"`
}
