package db

type DbStore struct { // interfaceと型は同義。仮にgatewayがDBの型を知ったとしても、どんなDBから来たかわかるわけではないのでおk
	Id   int
	Name string
}

// type DbDriver interface{
// 	GetStores() ([]*DbStore, error)
// }

// ここにGetStoresの実装を書く
