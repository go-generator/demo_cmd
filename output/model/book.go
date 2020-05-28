package model

type Book struct {
	CatId string `json:"catId,omitempty" gorm:"column:catid" bson:"catId,omitempty" dynamodbav:"catId,omitempty" firestore:"catId,omitempty"`

	ComId string `json:"comId,omitempty" gorm:"column:comid" bson:"comId,omitempty" dynamodbav:"comId,omitempty" firestore:"comId,omitempty"`

	Id string `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id"`

	Name string `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
}
