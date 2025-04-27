package entity

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Status   string `json:"status"`
    Location string `json:"location"`
}
