package main

//db, err := gorm.Open("mysql", "root:123456@/baidu?charset=utf8&parseTime=True")
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
	Name  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/baidu?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000, Name: "hello grom"})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	//db.Delete(&product)

	// Raw SQL
	rows, err := db.Raw("select id, code, name from products where id = ?", 14).Rows() // (*sql.Rows, error)
	if err != nil {
		panic(err)
	}
	var id uint
	var code string
	var name string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &code, &name)
		log.Println(id, code, name)
	}

}
