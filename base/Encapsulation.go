package base

/* 封装
如果一个变量或方法对调用方不可见，就是封装行为
go：首字目大小写进行封装
使用：方法名、结构体变量
优势：调用方只需专注于需要的、隐藏细节、防止被恶心操作
*/

func FengZhuang() {
	var p Family
	p.name = "lizhi"
}
