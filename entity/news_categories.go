package entity

//go:generate reform

// NewsCategories represents a row in news_categories table.
//
//reform:news_categories
type NewsCategories struct {
	Id         int64 `reform:"id,pk"`
	NewsId     int64 `reform:"newsid"`
	CategoryId int64 `reform:"categoryid"`
}
