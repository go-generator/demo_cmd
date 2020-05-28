package model

type BookCategory struct {
	CategoryId string `json:"categoryId" gorm:"column:categoryid;primary_key" bson:"_id" dynamodbav:"categoryId" firestore:"categoryId"`

	CategoryName string `json:"categoryName,omitempty" gorm:"column:categoryname" bson:"categoryName,omitempty" dynamodbav:"categoryName,omitempty" firestore:"categoryName,omitempty"`

	Book *[]Book `json:"book,omitempty" gorm:"column:book" bson:"book,omitempty" dynamodbav:"book,omitempty" firestore:"book,omitempty"`

	CompanyId string `json:"companyId" gorm:"column:companyid;primary_key" bson:"_id" dynamodbav:"companyId" firestore:"companyId"`
}
