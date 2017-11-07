package controllers

import (
	"crawl_movices/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"fmt"
)

type CrawlMovieController struct {
	beego.Controller
}

func (this *CrawlMovieController) CrawlMovie() {
	//sUrl := "https://movie.douban.com/subject/26688480/?from=subject-page"
	sUrl := "https://movie.douban.com/feed/subject/3604148/reviews]<br/>"
	rsp := httplib.Get(sUrl)
	sMovieHtml, err := rsp.String()
	if err != nil {
		panic(err)
	}

	var movieInfo models.MovieInfo

	movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
	movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
	movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
	movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
	movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
	movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
	movieInfo.Movie_span = models.GetMovieRunningTime(sMovieHtml)

	//创建好对象后，添加到数据库
	//id, _ := models.AddMovie(&movieInfo)
	//this.Ctx.WriteString(fmt.Sprintf("%v", id))

	//this.Ctx.WriteString("<h1>电影名称:" + models.GetMovieName(sMovieHtml) + "</h1><br/>")
	//this.Ctx.WriteString("<h2>电影导演:" + models.GetMovieDirector(sMovieHtml) + "</h2><br/>")
	//this.Ctx.WriteString("<h2>电影主演:" + models.GetMovieMainCharacters(sMovieHtml) + "</h2><br/>")
	//this.Ctx.WriteString("<h2>豆瓣评分:" + models.GetMovieGrade(sMovieHtml) + "</h2><br/>")
	//this.Ctx.WriteString("<h2>电影类型:" + models.GetMovieGenre(sMovieHtml) + "</h2><br/>")
	//this.Ctx.WriteString("<h2>上映日期:" + models.GetMovieOnTime(sMovieHtml) + "</h2><br/>")
	//this.Ctx.WriteString("<h2>电影片长:" + models.GetMovieRunningTime(sMovieHtml) + "</h2><br/>")

	//连接到Redis
	models.ConnectRedis("127.0.0.1", "6379", 1, "3071611103")
	urls := models.GetMovieUrls(sMovieHtml)
	for _, url := range urls {
		models.PushInQueue(url)
	}
	this.Ctx.WriteString(fmt.Sprintf("%v<br/>", urls))
}
