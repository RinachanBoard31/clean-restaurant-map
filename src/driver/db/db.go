package db

/* interfaceと型は同義。仮にgatewayがDBの型を知ったとしても、どんなDBから来たかわかるわけではないのでおk */
type DbStore struct {
	Id   int
	Name string
}

type DbStoreDriver struct{}

func NewStoreDriver() *DbStoreDriver {
	return &DbStoreDriver{}
}

func (dbs *DbStoreDriver) GetStores() ([]*DbStore, error) {
	// 本来はDBからデータを取得する処理
	return []*DbStore{
		{Id: 1, Name: "サイゼ"},
		{Id: 2, Name: "ジョナサン"},
		{Id: 3, Name: "OK調布店"},
	}, nil
}
