package main

//type Product struct {
//	Name      string
//	ProductID int64
//	Number    int
//	Price     float64
//	IsOnSale  bool
//}
//
//// json编码
//func main() {
//	p := &Product{}
//	p.Name = "Xiao mi 6"
//	p.IsOnSale = true
//	p.Number = 10000
//	p.Price = 2499.00
//	p.ProductID = 1
//	data, _ := json.Marshal(p)
//	fmt.Println(string(data))
//}

//type Product struct {
//	Name      string  `json:"name"`
//	ProductID int64   `json:"-"` // 表示不进行序列化
//	Number    int     `json:"number"`
//	Price     float64 `json:"price"`
//	IsOnSale  bool    `json:"is_on_sale,string"`
//}
//
//func main() {
//	p := &Product{}
//	p.Name = "Xiao mi 6"
//	p.IsOnSale = true
//	p.Number = 10000
//	p.Price = 2499.00
//	p.ProductID = 1
//	data, _ := json.Marshal(p)
//	fmt.Println(string(data))
//}

// 标签使用
//type Product struct {
//	Name      string  `json:"name"`
//	ProductID int64   `json:"product_id,omitempty"`
//	Number    int     `json:"number"`
//	Price     float64 `json:"price,string"`
//	IsOnSale  bool    `json:"is_on_sale,omitempty"`
//}
//func main() {
//	p := &Product{}
//	p.Name = "Xiao mi 6"
//	p.IsOnSale = true
//	p.Number = 10000
//	p.Price = 2499.00
//	p.ProductID = 0
//
//	data, _ := json.Marshal(p)
//	fmt.Println(string(data))
//}

// 解压
//type Product struct {
//	Name      string  `json:"name"`
//	ProductID int64   `json:"product_id,string"`
//	Number    int     `json:"number,string"`
//	Price     float64 `json:"price,string"`
//	IsOnSale  bool    `json:"is_on_sale,string"`
//}
//
//func main() {
//	var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
//	p := &Product{}
//	err := json.Unmarshal([]byte(data), p)
//	fmt.Println(err)
//	fmt.Println(p)
//}

