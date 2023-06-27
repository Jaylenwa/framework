package entity

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
	Home string `json:"home"`
}
