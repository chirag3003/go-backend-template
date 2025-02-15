package models

type Media struct {
  ID string `json:"id" bson:"_id"`
  Key string `json:"key" bson:"key"`
  Etag string `json:"etag" bson:"etag"`
  Size int64 `json:"size" bson:"size"`
  Mime string `json:"mime" bson:"mime"`
  Url string `json:"url" bson:"url"`
}
