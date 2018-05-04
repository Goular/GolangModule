package model

import "encoding/json"

type Profile struct {
	Url string
	Id  string

	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xinzuo     string
	House      string
	Car        string
}

// 将陌生的字符串转为对象的json字符串
func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
