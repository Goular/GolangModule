package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"regexp"
)

var (
	db orm.Ormer
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/crawl_movices?")
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
	_Create_time         string
}

//添加爬虫的电影片源信息
func AddMovie(movie_info *MovieInfo) (int64, error) {
	id, err := db.Insert(movie_info)
	return id, err
}

//获取电影的导演名字
func GetMovieDirector(movieHtml string) (string) {
	if movieHtml == "" {
		return ""
	} else {
		reg := regexp.MustCompile(`<a.*rel="v:directedBy">(.*)</a>`)
		result := reg.FindAllStringSubmatch(movieHtml, -1)
		return string(result[0][1])
	}
}

//获取电影的名字
func GetMovieName(movieHtml string) (string) {
	if movieHtml == "" {
		return ""
	} else {
		reg := regexp.MustCompile(`<span.*property="v:itemreviewed">(.*)</span>`)
		result := reg.FindAllStringSubmatch(movieHtml, -1)
		return string(result[0][1])
	}
}

//获取电影的主要演员列表
func GetMovieMainCharacters(movieHtml string) string {
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	mainCharacters := ""
	for _, v := range result {
		mainCharacters += v[1] + "/"
	}
	return mainCharacters
}
