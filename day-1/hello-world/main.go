// Trong Go,mỗi chương trình cần có một package "main" để có thể được biên dịch và chạy.
package main

// Được sử dụng để import package "fmt". Package "fmt" cung cấp các hàm để định dạng và in các chuỗi và giá trị.
import "fmt"

// Đây là hàm chính mà chương trình sẽ bắt đầu thực hiện
func main() {
	fmt.Println("Hello, I am Gopher!")
}
