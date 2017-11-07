package controllers

import (
	"crawl_movices/models"
	"github.com/astaxie/beego"
)

type CrawlMovieController struct {
	beego.Controller
}

func (this *CrawlMovieController) CrawlMovie() {
	sMovieHtml := `<!DOCTYPE html>
<html lang="zh-cmn-Hans" class="ua-windows ua-ff50">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="renderer" content="webkit">
    <meta name="referrer" content="always">
    <title>
        七月与安生 (豆瓣)
</title>

    <meta name="baidu-site-verification" content="cZdR4xxR7RxmM4zE" />
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Expires" content="Sun, 6 Mar 2005 01:00:00 GMT">

    <link rel="apple-touch-icon" href="https://img3.doubanio.com/f/movie/d59b2715fdea4968a450ee5f6c95c7d7a2030065/pics/movie/apple-touch-icon.png">
    <link href="https://img3.doubanio.com/f/shire/82d7c82f084d5157d7b87a17c76af984dd23dbbf/css/douban.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/shire/ae3f5a3e3085968370b1fc63afcecb22d3284848/css/separation/_all.css" rel="stylesheet" type="text/css">
    <script type="text/javascript">var _head_start = new Date();</script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/movie/e9d9543ebc06f2964039a2e94898f84ce77fc070/js/movie/lib/jquery.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/696bc71b3f521eaba8262ccec4859c9849288c6f/js/douban.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/0efdc63b77f895eaf85281fb0e44d435c6239a3f/js/separation/_all.js"></script>

    <meta name="keywords" content="七月与安生,七月与安生,七月与安生影评,剧情介绍,电影图片,预告片,影讯,在线购票,论坛">
    <meta name="description" content="七月与安生电影简介和剧情介绍,七月与安生影评、图片、预告片、影讯、论坛、在线购票">
    <meta name="mobile-agent" content="format=html5; url=http://m.douban.com/movie/subject/25827935/"/>
    <link rel="alternate" href="android-app://com.douban.movie/doubanmovie/subject/25827935/" />
    <link rel="stylesheet" href="https://img3.doubanio.com/dae/cdnlib/libs/LikeButton/1.0.5/style.min.css">
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript">
        Do.add('dialog', {path: 'https://img3.doubanio.com/f/shire/3d185ca912c999ee7f464749201139ebf8eb6972/js/ui/dialog.js', type: 'js'});
        Do.add('dialog-css', {path: 'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css', type: 'css'});
        Do.add('handlebarsjs', {path: 'https://img3.doubanio.com/f/movie/ca54296583b9faa46ca0d1eec2c85d855a5f3ca1/js/movie/lib/handlebars.current.js', type: 'js'});
    </script>

  <script type='text/javascript'>
  var _vwo_code = (function() {
    var account_id = 249272,
      settings_tolerance = 2000,
      library_tolerance = 2500,
      use_existing_jquery = false,
      // DO NOT EDIT BELOW THIS LINE
      f=false,d=document;return{use_existing_jquery:function(){return use_existing_jquery;},library_tolerance:function(){return library_tolerance;},finish:function(){if(!f){f=true;var a=d.getElementById('_vis_opt_path_hides');if(a)a.parentNode.removeChild(a);}},finished:function(){return f;},load:function(a){var b=d.createElement('script');b.src=a;b.type='text/javascript';b.innerText;b.onerror=function(){_vwo_code.finish();};d.getElementsByTagName('head')[0].appendChild(b);},init:function(){settings_timer=setTimeout('_vwo_code.finish()',settings_tolerance);var a=d.createElement('style'),b='body{opacity:0 !important;filter:alpha(opacity=0) !important;background:none !important;}',h=d.getElementsByTagName('head')[0];a.setAttribute('id','_vis_opt_path_hides');a.setAttribute('type','text/css');if(a.styleSheet)a.styleSheet.cssText=b;else a.appendChild(d.createTextNode(b));h.appendChild(a);this.load('//dev.visualwebsiteoptimizer.com/j.php?a='+account_id+'&u='+encodeURIComponent(d.URL)+'&r='+Math.random());return settings_timer;}};}());

  +function () {
    var bindEvent = function (el, type, handler) {
        var $ = window.jQuery || window.Zepto || window.$
       if ($ && $.fn && $.fn.on) {
           $(el).on(type, handler)
       } else if($ && $.fn && $.fn.bind) {
           $(el).bind(type, handler)
       } else if (el.addEventListener){
         el.addEventListener(type, handler, false);
       } else if (el.attachEvent){
         el.attachEvent("on" + type, handler);
       } else {
         el["on" + type] = handler;
       }
     }

    var _origin_load = _vwo_code.load
    _vwo_code.load = function () {
      var args = [].slice.call(arguments)
      bindEvent(window, 'load', function () {
        _origin_load.apply(_vwo_code, args)
      })
    }
  }()

  _vwo_settings_timer = _vwo_code.init();
  </script>


    <style type="text/css">

</style>
    <style type="text/css">img { max-width: 100%; }</style>
    <script type="text/javascript"></script>
    <link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/76cf2ded37a390ea.css">

    <link rel="shortcut icon" href="https://img3.doubanio.com/favicon.ico" type="image/x-icon">
</head>

<body>

    <script type="text/javascript">var _body_start = new Date();</script>










<div id="db-global-nav" class="global-nav">
  <div class="bd">






<div class="top-nav-info">

      <a href="https://www.douban.com/accounts/login?source=movie" class="nav-login" rel="nofollow">登录</a>
      <a href="https://www.douban.com/accounts/register?source=movie" class="nav-register" rel="nofollow">注册</a>
</div>


    <div class="top-nav-doubanapp">

        <a href="https://www.douban.com/doubanapp/app?channel=top-nav" class="lnk-doubanapp">下载豆瓣客户端</a>
        <div id="top-nav-appintro" class="more-items">
            <p class="appintro-title">豆瓣</p>
            <p class="slogan">我们的精神角落</p>
            <p class="qrcode">
                扫码直接下载
            </p>
            <div class="download">
                <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&amp;direct_dl=1&amp;download=iOS">iPhone</a>
                <span>·</span>
                <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&amp;direct_dl=1&amp;download=Android" class="download-android">Android</a>
            </div>
        </div>
        <div id="doubanapp-tip">
            <a href="https://www.douban.com/doubanapp/app?channel=qipao" class="tip-link">豆瓣 4.0 全新发布</a>
            <a href="javascript: void 0;" class="tip-close">×</a>
        </div>
    </div>

    <div class="global-nav-items">
      <ul>


            <li class="">
            <a href="https://www.douban.com/" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-main&#34;,&#34;uid&#34;:&#34;0&#34;}">
              豆瓣
            </a>
            </li>


            <li class="">
            <a href="https://book.douban.com/" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-book&#34;,&#34;uid&#34;:&#34;0&#34;}">
              读书
            </a>
            </li>


            <li class="on">
            <a href="https://movie.douban.com/" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-movie&#34;,&#34;uid&#34;:&#34;0&#34;}">
              电影
            </a>
            </li>


            <li class="">
            <a href="https://music.douban.com/" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-music&#34;,&#34;uid&#34;:&#34;0&#34;}">
              音乐
            </a>
            </li>


            <li class="">
            <a href="https://www.douban.com/location/" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-location&#34;,&#34;uid&#34;:&#34;0&#34;}">
              同城
            </a>
            </li>


            <li class="">
            <a href="https://www.douban.com/group/" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-group&#34;,&#34;uid&#34;:&#34;0&#34;}">
              小组
            </a>
            </li>


            <li class="">
            <a href="https://read.douban.com/?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-read&#34;,&#34;uid&#34;:&#34;0&#34;}">
              阅读
            </a>
            </li>


            <li class="">
            <a href="https://douban.fm/?from_=shire_top_nav" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-fm&#34;,&#34;uid&#34;:&#34;0&#34;}">
              FM
            </a>
            </li>


            <li class="">
            <a href="https://dongxi.douban.com/?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-commodity&#34;,&#34;uid&#34;:&#34;0&#34;}">
              东西
            </a>
            </li>


            <li class="">
            <a href="https://market.douban.com/?utm_campaign=douban_top_nav&amp;utm_source=douban&amp;utm_medium=pc_web" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-market&#34;,&#34;uid&#34;:&#34;0&#34;}">
              市集
            </a>
            </li>

            <li>
              <a href="#more" class="bn-more"><span>更多</span></a>
              <div class="more-items">
                <table cellpadding="0" cellspacing="0">

                    <tr><td><a href="https://moment.douban.com" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-moment&#34;,&#34;uid&#34;:&#34;0&#34;}">一刻</a></td></tr>

                    <tr><td><a href="https://ypy.douban.com" target="_blank" data-moreurl-dict="{&#34;from&#34;:&#34;top-nav-click-ypy&#34;,&#34;uid&#34;:&#34;0&#34;}">豆瓣摄影</a></td></tr>
                </table>
              </div>
            </li>
      </ul>
    </div>
  </div>
</div>












<div id="db-nav-movie" class="nav">
  <div class="nav-wrap">
  <div class="nav-primary">
    <div class="nav-logo">
      <a href="https://movie.douban.com">豆瓣电影</a>
    </div>
    <div class="nav-search">
      <form action="https://movie.douban.com/subject_search" method="get">
        <fieldset>
          <legend>搜索：</legend>
          <label for="inp-query">电影、影人、影院、电视剧</label>
          <div class="inp"><input id="inp-query" name="search_text" size="22" maxlength="60" value=""></div>
          <div class="inp-btn"><input type="submit" value="搜索"></div>
          <input type="hidden" name="cat" value="1002" />
        </fieldset>
      </form>
    </div>

    <!-- header ad begin -->






    <div id="dale_movie_top_nav_subject" class="ban-header"></div>




    <!-- header ad end -->

  </div>
  </div>
  <div class="nav-secondary">

      <div class="nav-items nav-anon">
          <ul>

                     <li><a href="https://movie.douban.com/nowplaying/">影讯&amp;购票</a></li>

                     <li><a href="https://movie.douban.com/explore">选电影</a></li>

                    <li class="site-nav-bt">
                        <a href="https://movie.douban.com/tv/">电视剧</a>
                    </li>

                     <li><a href="https://movie.douban.com/chart">排行榜</a></li>

                     <li><a href="https://movie.douban.com/tag/">分类</a></li>

                     <li><a href="https://movie.douban.com/review/best/">影评</a></li>

                     <li><a href="https://movie.douban.com/annual2015/">2015年度榜单</a></li>

                     <li><a href="https://movie.douban.com/standbyme/?source=navigation">豆瓣观影报告</a></li>

          </ul>
      </div>
  </div>
</div>



  <script id="suggResult" type="text/x-jquery-tmpl">
      <li data-link="{{= url}}">
          <a href="{{= url}}" onclick="moreurl(this, {from:'movie_search_sugg'})">
              <img src="{{= img}}" width="40" />
              <p>
                  <em>{{= title}}</em>
                  {{if year}}
                      <span>{{= year}}</span>
                  {{/if}}
                  {{if sub_title}}
                      <br /><span>{{= sub_title}}</span>
                  {{/if}}
                  {{if address}}
                      <br /><span>{{= address}}</span>
                  {{/if}}
                  {{if episode}}
                      {{if episode=="unknow"}}
                          <br /><span>集数未知</span>
                      {{else}}
                          <br /><span>共{{= episode}}集</span>
                      {{/if}}
                  {{/if}}
              </p>
          </a>
      </li>
  </script>







    <div id="wrapper">



    <div id="content">


    <div id="dale_movie_subject_top_icon"></div>
    <h1>
        <span property="v:itemreviewed">七月与安生</span>
            <span class="year">(2016)</span>
    </h1>

        <div class="grid-16-8 clearfix">



            <div class="article">









        <div class="indent clearfix">
            <div class="subjectwrap clearfix" xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Movie">
                <div class="subject clearfix">




<div id="mainpic" class="">
    <a class="nbgnbg" href="https://movie.douban.com/subject/25827935/photos?type=R" title="点击看更多海报">
        <img src="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p2378140502.jpg" title="点击看更多海报" alt="七月与安生" rel="v:image" />
   </a>
</div>




<div id="info">
        <span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/1274534/" rel="v:directedBy">曾国祥</a></span></span><br/>
        <span ><span class='pl'>编剧</span>: <span class='attrs'><a href="/celebrity/1359881/">林咏琛</a> / <a href="/celebrity/1359882/">李媛</a> / <a href="/celebrity/1330070/">许伊萌</a> / <a href="/celebrity/1322986/">吴楠</a> / <a href="/celebrity/1355461/">安妮宝贝</a></span></span><br/>
        <span class="actor"><span class='pl'>主演</span>: <span class='attrs'><a href="/celebrity/1274224/" rel="v:starring">周冬雨</a> / <a href="/celebrity/1275243/" rel="v:starring">马思纯</a> / <a href="/celebrity/1349387/" rel="v:starring">李程彬</a> / <a href="/celebrity/1328349/" rel="v:starring">李萍</a> / <a href="/celebrity/1365506/" rel="v:starring">蔡纲</a> / <a href="/celebrity/1316715/" rel="v:starring">蒙亭宜</a> / <a href="/celebrity/1365507/" rel="v:starring">沙全泽</a> / <a href="/celebrity/1330226/" rel="v:starring">姚欣言</a> / <a href="/celebrity/1365508/" rel="v:starring">李昊芳</a> / <a href="/celebrity/1365509/" rel="v:starring">蒋亭轩</a></span></span><br/>
        <span class="pl">类型:</span> <span property="v:genre">剧情</span> / <span property="v:genre">爱情</span><br/>

        <span class="pl">制片国家/地区:</span> 中国大陆 / 香港<br/>
        <span class="pl">语言:</span> 汉语普通话<br/>
        <span class="pl">上映日期:</span> <span property="v:initialReleaseDate" content="2016-09-14(中国大陆)">2016-09-14(中国大陆)</span> / <span property="v:initialReleaseDate" content="2016-10-27(香港)">2016-10-27(香港)</span><br/>
        <span class="pl">片长:</span> <span property="v:runtime" content="109">109分钟</span><br/>
        <span class="pl">又名:</span> SoulMate<br/>
        <span class="pl">IMDb链接:</span> <a href="http://www.imdb.com/title/tt6054290" target="_blank" rel="nofollow">tt6054290</a><br>

</div>




                </div>



<div id="interest_sectl">
    <div class="rating_wrap clearbox" rel="v:rating">
        <div class="rating_logo">豆瓣评分</div>
        <div class="rating_self clearfix" typeof="v:Rating">
            <strong class="ll rating_num" property="v:average">7.6</strong>
            <span property="v:best" content="10.0"></span>
            <div class="rating_right ">
                <div class="ll bigstar40"></div>
                <div class="rating_sum">
                    <a href="collections" class="rating_people"><span property="v:votes">162418</span>人评价</a>
                </div>
            </div>
        </div>


                <span class="stars5 starstop" title="力荐">
                    5星
                </span>
                <div class="power" style="width:24px"></div>
                <span class="rating_per">18.8%</span>
                <br />

                <span class="stars4 starstop" title="推荐">
                    4星
                </span>
                <div class="power" style="width:64px"></div>
                <span class="rating_per">48.7%</span>
                <br />

                <span class="stars3 starstop" title="还行">
                    3星
                </span>
                <div class="power" style="width:36px"></div>
                <span class="rating_per">27.4%</span>
                <br />

                <span class="stars2 starstop" title="较差">
                    2星
                </span>
                <div class="power" style="width:5px"></div>
                <span class="rating_per">3.9%</span>
                <br />

                <span class="stars1 starstop" title="很差">
                    1星
                </span>
                <div class="power" style="width:1px"></div>
                <span class="rating_per">1.3%</span>
                <br />

    </div>
        <div class="rating_betterthan">
            好于 <a href="/typerank?type_name=爱情&type=13&interval_id=75:65&action=">71% 爱情片</a><br/>
            好于 <a href="/typerank?type_name=剧情&type=11&interval_id=60:50&action=">56% 剧情片</a><br/>
        </div>
</div>



            </div>





<div id="interest_sect_level" class="clearfix">

            <a href="https://www.douban.com/reason=collectwish&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-25827935-wish">
                <span>想看</span>
            </a>
            <a href="https://www.douban.com/reason=collectcollect&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-25827935-collect">
                <span>看过</span>
            </a>
        <div class="ll j a_stars">


    评价:
    <span id="rating"> <span id="stars" data-solid="https://img3.doubanio.com/f/shire/5a2327c04c0c231bced131ddf3f4467eb80c1c86/pics/rating_icons/star_onmouseover.png" data-hollow="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" data-solid-2x="https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png" data-hollow-2x="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png">

            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-1">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star1" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-2">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star2" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-3">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star3" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-4">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star4" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-5">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star5" width="16" height="16"/></a>
    </span><span id="rateword" class="pl"></span>
    <input id="n_rating" type="hidden" value=""  />
    </span>

        </div>


</div>





















<div class="gtleft">
    <ul class="ul_subject_menu bicelink color_gray pt6 clearfix">



                <li>
    <img src="https://img3.doubanio.com/f/shire/cc03d0fcf32b7ce3af7b160a0b85e5e66b47cc42/pics/short-comment.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_cmnt_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写短评</a>
 </li>
                    <li>

    <img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_rv_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写影评</a>
 </li>
                    <li>
    <img src="https://img3.doubanio.com/f/shire/61cc48ba7c40e0272d46bb93fe0dc514f3b71ec5/pics/add-doulist.gif" />&nbsp;
    <a href="/subject/25827935/questions/ask?from=subject_top">提问题</a>
 </li>
                <li>




 </li>
                <li>




    <span class="rec" id="电影-25827935">
    <a href= "#"
        data-type="电影"
        data-url="https://movie.douban.com/subject/25827935/"
        data-desc="电影《七月与安生》 (来自豆瓣) "
        data-title="电影《七月与安生》 (来自豆瓣) "
        data-pic="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p2378140502.jpg"
        class="bn-sharing ">
        分享到
    </a> &nbsp;&nbsp;
    </span>

    <script>
        if (!window.DoubanShareMenuList) {
            window.DoubanShareMenuList = [];
        }
        var __cache_url = __cache_url || {};

        (function(u){
            if(__cache_url[u]) return;
            __cache_url[u] = true;
            window.DoubanShareIcons = 'https://img3.doubanio.com/f/shire/d15ffd71f3f10a7210448fec5a68eaec66e7f7d0/pics/ic_shares.png';

            var initShareButton = function() {
                $.ajax({url:u,dataType:'script',cache:true});
            };

            if (typeof Do == 'function' && 'ready' in Do) {
                Do(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/3d185ca912c999ee7f464749201139ebf8eb6972/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js',
                    initShareButton
                );
            } else if(typeof Douban == 'object' && 'loader' in Douban) {
                Douban.loader.batch(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/3d185ca912c999ee7f464749201139ebf8eb6972/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js'
                ).done(initShareButton);
            }

        })('https://img3.doubanio.com/f/movie/2ec292f1a260672c94a9417e81b4f1667ed42583/js/movie/lib/sharebutton.js');
    </script>


  </li>


    </ul>

    <script type="text/javascript">
        $(function(){
            $(".ul_subject_menu li.rec .bn-sharing").bind("click", function(){
                $.get("/blank?sbj_page_click=bn_sharing");
            });
        });
    </script>
</div>










<div class="rec-sec">
<span class="rec">
    <script id="movie-share" type="text/x-html-snippet">

    <form class="movie-share" action="/j/share" method="POST">
        <div class="clearfix form-bd">
            <div class="input-area">
                <textarea name="text" class="share-text" cols="72" data-mention-api="https://api.douban.com/shuo/in/complete?alt=xd&amp;callback=?"></textarea>
                <input type="hidden" name="target-id" value="25827935">
                <input type="hidden" name="target-type" value="0">
                <input type="hidden" name="title" value="七月与安生 (2016)">
                <input type="hidden" name="desc" value="导演 曾国祥 主演 周冬雨 / 马思纯 / 中国大陆 / 香港 / 7.6分(162418评价)">
                <input type="hidden" name="redir" value=""/>
                <div class="mentioned-highlighter"></div>
            </div>

            <div class="info-area">
                    <img class="media" src="https://img3.doubanio.com/view/movie_poster_cover/ipst/public/p2378140502.jpg" />
                <strong>七月与安生 (2016)</strong>
                <p>导演 曾国祥 主演 周冬雨 / 马思纯 / 中国大陆 / 香港 / 7.6分(162418评价)</p>
                <p class="error server-error">&nbsp;</p>
            </div>
        </div>
        <div class="form-ft">
            <div class="form-ft-inner">




                <span class="avail-num-indicator">140</span>
                <span class="bn-flat">
                    <input type="submit" value="推荐" />
                </span>
            </div>
        </div>
    </form>

    <div id="suggest-mention-tmpl" style="display:none;">
        <ul>
            {{#users}}
            <li id="{{uid}}">
              <img src="{{avatar}}">{{{username}}}&nbsp;<span>({{{uid}}})</span>
            </li>
            {{/users}}
        </ul>
    </div>


    </script>


        <a href="/accounts/register?reason=recommend"  class="j a_show_login lnk-sharing" share-id="25827935" data-mode="plain" data-name="七月与安生 (2016)" data-type="movie" data-desc="导演 曾国祥 主演 周冬雨 / 马思纯 / 中国大陆 / 香港 / 7.6分(162418评价)" data-href="https://movie.douban.com/subject/25827935/" data-image="https://img3.doubanio.com/view/movie_poster_cover/ipst/public/p2378140502.jpg" data-properties="{}" data-redir="" data-text="" data-apikey="" data-curl="" data-count="10" data-object_kind="1002" data-object_id="25827935" data-target_type="rec" data-target_action="0" data-action_props="{&#34;subject_url&#34;:&#34;https:\/\/movie.douban.com\/subject\/25827935\/&#34;,&#34;subject_title&#34;:&#34;七月与安生 (2016)&#34;}">推荐</a>
</span>


</div>






            <script type="text/javascript">
                $(function() {
                    $('.collect_btn', '#interest_sect_level').each(function() {
                        Douban.init_collect_btn(this);
                    });
                    $('html').delegate(".indent .rec-sec .lnk-sharing", "click", function() {
                        moreurl(this, {
                            from : 'mv_sbj_db_share'
                        });
                    });
                });
            </script>
        </div>



    <div id="collect_form_25827935"></div>



<div class="related-info" style="margin-bottom:-10px;">
    <a name="intro"></a>



            <h2>七月与安生的剧情简介&nbsp;&nbsp;·&nbsp;&nbsp;·&nbsp;&nbsp;·&nbsp;&nbsp;·&nbsp;&nbsp;·&nbsp;&nbsp;·</h2>
            <div class="indent" id="link-report">

                        <span property="v:summary" class="">
                                　　七月（马思纯 饰）和安生（周冬雨 饰）的第一次相识在十三岁，她们一个是特立独行飞扬跋扈的“野孩子”，一个是单纯温婉循规蹈矩的“乖乖女”，从那一年开始，七月和安生几乎形影不离，她是她的光，她是她的影子，直到某一天，一位名为苏家明（李程彬 饰）的少年出现在了七月的身边，七月恋爱了。
                                    <br />
                                　　安生决定前往北京讨生活，临别之前，七月意外的发现苏家明贴身带着的玉佩，竟然出现在了安生的衣领里。安生走了，七月和苏家明的恋情持续着，他们考入了同一所大学，约定一毕业就结婚。可是，事情并没有像七月所想象的那样发展，而她和苏家明之间的关系，亦因为安生的归来而产生了新的变数。
                        </span>
                        <span class="pl"><a href="https://movie.douban.com/help/movie#t0-qs">&copy;豆瓣</a></span>
            </div>
</div>













    <div id="related-pic" class="related-pic">



    <h2>
        <i class="">七月与安生的预告片和图片</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/25827935/trailer">预告片12</a>&nbsp;|&nbsp; <a href="https://movie.douban.com/subject/25827935/all_photos">图片618</a>&nbsp;|&nbsp; <a href="https://movie.douban.com/subject/25827935/mupload">添加图片</a>
            ) </span>
    </h2>


        <ul class="related-pic-bd narrow">
                <li>
                    <a class="related-pic-video" href="https://movie.douban.com/trailer/203039/#content">
                        <span></span>
                        <img src="https://img1.doubanio.com/img/trailer/medium/2378146297.jpg?" width="178" height="100" alt="预告片" />
                    </a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2368151442/"><img src="https://img3.doubanio.com/view/photo/albumicon/public/p2368151442.jpg" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2375428209/"><img src="https://img1.doubanio.com/view/photo/albumicon/public/p2375428209.jpg" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2372679263/"><img src="https://img3.doubanio.com/view/photo/albumicon/public/p2372679263.jpg" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2378508116/"><img src="https://img5.doubanio.com/view/photo/albumicon/public/p2378508116.jpg" alt="图片" /></a>
                </li>
        </ul>
    </div>









<style type="text/css">
.award li { display: inline; margin-right: 5px }
.awards { margin-bottom: 20px }
.awards h2 { background: none; color: #000; font-size: 14px; padding-bottom: 5px; margin-bottom: 8px; border-bottom: 1px dashed #dddddd }
.awards .year { color: #666666; margin-left: -5px }
.mod { margin-bottom: 25px }
.mod .hd { margin-bottom: 10px }
.mod .hd h2 {margin:24px 0 3px 0}
</style>


<div class="mod">
    <div class="hd">

    <h2>
        <i class="">七月与安生的获奖情况</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/25827935/awards/">全部</a>
            ) </span>
    </h2>

    </div>

        <ul class="award">
            <li>
                <a href="https://movie.douban.com/awards/goldenhorse/53/">第53届台北金马影展</a>
            </li>
            <li>金马奖 最佳导演(提名)</li>
            <li><a href='https://movie.douban.com/celebrity/1274534/' target='_blank'>曾国祥</a></li>
        </ul>

        <ul class="award">
            <li>
                <a href="https://movie.douban.com/awards/goldenhorse/53/">第53届台北金马影展</a>
            </li>
            <li>金马奖 最佳女主角</li>
            <li><a href='https://movie.douban.com/celebrity/1274224/' target='_blank'>周冬雨</a>&nbsp;/&nbsp;<a href='https://movie.douban.com/celebrity/1275243/' target='_blank'>马思纯</a></li>
        </ul>

        <ul class="award">
            <li>
                <a href="https://movie.douban.com/awards/goldenhorse/53/">第53届台北金马影展</a>
            </li>
            <li>金马奖 最佳改编剧本(提名)</li>
            <li><a href='https://movie.douban.com/celebrity/1359881/' target='_blank'>林咏琛</a>&nbsp;/&nbsp;<a href='https://movie.douban.com/celebrity/1359882/' target='_blank'>李媛</a>&nbsp;/&nbsp;<a href='https://movie.douban.com/celebrity/1322986/' target='_blank'>吴楠</a>&nbsp;/&nbsp;<a href='https://movie.douban.com/celebrity/1330070/' target='_blank'>许伊萌</a></li>
        </ul>
</div>










    <div id="recommendations" class="">


    <h2>
        <i class="">喜欢这部电影的人也喜欢</i>
              · · · · · ·
    </h2>



    <div class="recommendations-bd">
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/3319755/?from=subject-page" >
                    <img src="https://img5.doubanio.com/view/movie_poster_cover/lpst/public/p663036666.jpg" alt="怦然心动" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/3319755/?from=subject-page" class="" >怦然心动</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26366465/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p2285115802.jpg" alt="我的少女时代" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26366465/?from=subject-page" class="" >我的少女时代</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/1308575/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p584822570.jpg" alt="蓝色大门" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/1308575/?from=subject-page" class="" >蓝色大门</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/4917726/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/movie_poster_cover/lpst/public/p1374786017.jpg" alt="阳光姐妹淘" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/4917726/?from=subject-page" class="" >阳光姐妹淘</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/4920528/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p1062824805.jpg" alt="那些年，我们一起追的女孩" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/4920528/?from=subject-page" class="" >那些年，我们一起追的女孩</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/4739952/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p1505312273.jpg" alt="初恋这件小事" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/4739952/?from=subject-page" class="" >初恋这件小事</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26683290/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/movie_poster_cover/lpst/public/p2395733377.jpg" alt="你的名字。" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26683290/?from=subject-page" class="" >你的名字。</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26259677/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/movie_poster_cover/lpst/public/p2329080834.jpg" alt="垫底辣妹" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26259677/?from=subject-page" class="" >垫底辣妹</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/11529526/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/movie_poster_cover/lpst/public/p1959304567.jpg" alt="中国合伙人" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/11529526/?from=subject-page" class="" >中国合伙人</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/1292402/?from=subject-page" >
                    <img src="https://img5.doubanio.com/view/movie_poster_cover/lpst/public/p792400696.jpg" alt="西西里的美丽传说" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/1292402/?from=subject-page" class="" >西西里的美丽传说</a>
            </dd>
        </dl>
    </div>

    </div>





<script type="text/x-handlebar-tmpl" id="comment-tmpl">
    <div class="dummy-fold">
        {{#each comments}}
        <div class="comment-item" data-cid="id">
            <div class="comment">
                <h3>
                    <span class="comment-vote">
                            <span class="votes pr5">{{votes}}</span>
                        <input value="{{id}}" type="hidden"/>
                        <a href="javascript:;" class="j {{#if ../if_logined}}a_vote_comment{{else}}a_show_login{{/if}}">有用</a>
                    </span>
                    <span class="comment-info">
                        <a href="{{user.path}}" class="">{{user.name}}</a>
                        {{#if rating}}
                        <span class="allstar{{rating}}0 rating" title="{{rating_word}}"></span>
                        {{/if}}
                        <span>
                            {{time}}
                        </span>
                        <p> {{content}} </p>
                    </span>
            </div>
        </div>
        {{/each}}
    </div>
</script>














    <div id="comments-section">
        <div class="mod-hd">

        <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">
            <span>我要写短评</span>
        </a>



    <h2>
        <i class="">七月与安生的短评</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/25827935/comments?status=P">全部 77223 条</a>
            ) </span>
    </h2>

        </div>
        <div class="mod-bd">

    <div class="tab-hd">
        <a id="hot-comments-tab" href="comments" data-id="hot" class="on">热门</a>&nbsp;/&nbsp;
        <a id="new-comments-tab" href="comments?sort=time" data-id="new">最新</a>&nbsp;/&nbsp;
        <a id="following-comments-tab" href="follows_comments" data-id="following"  class="j a_show_login">好友</a>
    </div>

    <div class="tab-bd">
        <div id="hot-comments" class="tab">


        <div class="comment-item" data-cid="1084594509">


    <div class="comment">
        <h3>

            <span class="comment-vote">
                <span class="votes pr5">5242</span>
                <input value="1084594509" type="hidden"/>
                <a href="javascript:;" class="j a_show_login">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/51061132/" class="">indicent time</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time ">
                    2016-09-14
                </span>
            </span>
        </h3>
        <p class=""> 在大陆青春片里看到七月与安生，就如同在屎堆里掏出了一块银子，对，还不是金子，但是终于不再是夏有乔木，不再是微微一笑，不再是原来你还在这里，不再是左耳，不再是……更不再是tnnd同桌的你
        </p>
    </div>

        </div>
        <div class="comment-item" data-cid="1085370135">


    <div class="comment">
        <h3>

            <span class="comment-vote">
                <span class="votes pr5">777</span>
                <input value="1085370135" type="hidden"/>
                <a href="javascript:;" class="j a_show_login">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/renjiananhuo/" class="">邓安庆</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time ">
                    2016-09-16
                </span>
            </span>
        </h3>
        <p class=""> 1：剧本扎实，结构完整，女性向，那种相爱相杀互生互灭的一体感，表达得很细致；2：周冬雨演得不错啊，马思纯跟高圆圆我有点分不清了，浴室对戏很打动人；3：女性向电影，都感觉男人真是个可有可无的物种，这种蠢物不要也罢！
        </p>
    </div>

        </div>
        <div class="comment-item" data-cid="1085176805">


    <div class="comment">
        <h3>

            <span class="comment-vote">
                <span class="votes pr5">368</span>
                <input value="1085176805" type="hidden"/>
                <a href="javascript:;" class="j a_show_login">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/antonia/" class="">安东妮</a>
                    <span>看过</span>
                    <span class="allstar10 rating" title="很差"></span>
                <span class="comment-time ">
                    2016-09-15
                </span>
            </span>
        </h3>
        <p class=""> 希望新一辈的情感教育不再是安宝这么残障
        </p>
    </div>

        </div>
        <div class="comment-item" data-cid="1090317031">


    <div class="comment">
        <h3>

            <span class="comment-vote">
                <span class="votes pr5">379</span>
                <input value="1090317031" type="hidden"/>
                <a href="javascript:;" class="j a_show_login">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/waggawagga/" class="">H2O.</a>
                    <span>看过</span>
                    <span class="allstar20 rating" title="较差"></span>
                <span class="comment-time ">
                    2016-09-28
                </span>
            </span>
        </h3>
        <p class=""> 年龄大了，这种友谊想着都头疼。
        </p>
    </div>

        </div>




                &gt; <a href="comments?sort=new_score&status=P" >更多短评77223条</a>
        </div>
        <div id="new-comments" class="tab">
            <div id="normal">
            </div>
            <div class="fold-hd hide">
                <a class="qa" href="/help/opinion#t2-q0" target="_blank">为什么被折叠？</a>
                <a class="btn-unfold" href="#">有一些短评被折叠了</a>
                <div class="qa-tip">
                    评论被折叠，是因为发布这条评论的帐号行为异常。评论仍可以被展开阅读，对发布人的账号不造成其他影响。如果认为有问题，可以<a href="https://help.douban.com/help/ask?category=movie">联系</a>豆瓣电影。
                </div>
            </div>
            <div class="fold-bd">
            </div>
            <span id="total-num"></span>
        </div>
        <div id="following-comments" class="tab">




        <div class="comment-item">
            你关注的人还没写过短评
        </div>

        </div>
    </div>




        </div>
    </div>












<div id="askmatrix">
    <div class="mod-hd">
        <h2>
            关于《七月与安生》的问题
            · · · · · ·
            <span class="pl">
                (<a href='https://movie.douban.com/subject/25827935/questions/?from=subject'>
                    全部116个
                </a>)
            </span>
        </h2>




    <a class='j a_show_login comment_btn'
        href='https://movie.douban.com/subject/25827935/questions/ask/?from=subject'>我来提问</a>

    </div>

    <div class="mod-bd">
        <ul class="">
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/25827935/questions/727598/?from=subject" class="">
                            片尾为什么要感谢岩井俊二？
                    </a>
                </span>
                <span class="meta">
                    28人回答
                </span>
            </li>
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/25827935/questions/728068/?from=subject" class="">
                            只有我觉得安生深深地爱着的是七月，同性之爱的那种？
                    </a>
                </span>
                <span class="meta">
                    21人回答
                </span>
            </li>
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/25827935/questions/728096/?from=subject" class="">
                            家明是爱安生多一些还是爱七月多一些？
                    </a>
                </span>
                <span class="meta">
                    43人回答
                </span>
            </li>
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/25827935/questions/727703/?from=subject" class="">
                            开头家明追逐安生的过程中掉了个钱包有什么寓意？
                    </a>
                </span>
                <span class="meta">
                    25人回答
                </span>
            </li>
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/25827935/questions/727693/?from=subject" class="">
                            安生的男友老赵去哪了？
                    </a>
                </span>
                <span class="meta">
                    23人回答
                </span>
            </li>
        </ul>

        <p>&gt;
            <a href='https://movie.douban.com/subject/25827935/questions/?from=subject'>
                全部116个问题
            </a>
        </p>

    </div>
</div>






<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/33be3e9be41abbce.css">

<section class="reviews mod movie-content">
    <header>
        <a href="new_review" rel="nofollow" class="comment_btn">
            <span>我要写影评</span>
        </a>
        <h2>
            七月与安生的影评 · · · · · ·
            <span class="pl">(<a href="reviews">全部 3340 条</a>)</span>
        </h2>
    </header>





<div class="review-list">







  <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review">
    <div class="main review-item" id="8076701">


  <header class="main-hd">
    <h3 class="title">
      <a href="https://movie.douban.com/review/8076701/" class="title-link">《Vista看天下》副主编孟静：男朋友和闺蜜出轨，...</a>
      <div class="toggle_review right">        <a href="javascript:;" id="toggle-8076701"class="indicator unfold j" title="展开全文">        </a>
      </div>
    </h3>


<div class="header-more">
    <a class="author-avatar" href="https://www.douban.com/people/44724062/"><img width="48" height="48" src="https://img3.doubanio.com/icon/u44724062-2.jpg"></a>
    <a href="https://www.douban.com/people/44724062/" class="author">
      <span property="v:reviewer">夏！</span>
    </a>    <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>
  <span property="v:dtreviewed" content="2016-09-08"
      class="main-meta">2016-09-08 12:21:14</span>
</div>

  </header>

        <div class="main-bd">
          <div id="review_8076701_short" class="review-short">
            <div class="short-content">
                  去看电影《七月与安生》，在那之前，片方的发布会找了几个嘉宾辩论，主题是：你愿意当七月还是安生？遇到闺蜜抢男友，撕还是让？  这几个嘉宾里也有我，我选了七月，被抢男友的那个。在我看来，不存在撕或者让，如果我是七月，我尽量不让男友见到闺蜜，不让他们有互留微信和电...

                    <a class="pl" href="https://movie.douban.com/review/8076701/#comments">(500回应)</a>
                  <div class="more-info pl clearfix">
                    <span class="left">4552有用 / 661没用</span>
                  </div>
            </div>
          </div>
          <div id="review_8076701_full" class="hidden">
            <div id="review_8076701_full_content" class="full-content"></div>
          </div>
        </div>
    </div>
  </div>






  <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review">
    <div class="main review-item" id="8075045">


  <header class="main-hd">
    <h3 class="title">
      <a href="https://movie.douban.com/review/8075045/" class="title-link">给你一个七月，换你一个安生</a>
      <div class="toggle_review right">        <a href="javascript:;" id="toggle-8075045"class="indicator unfold j" title="展开全文">        </a>
      </div>
    </h3>


<div class="header-more">
    <a class="author-avatar" href="https://www.douban.com/people/nangongtianlun/"><img width="48" height="48" src="https://img3.doubanio.com/icon/u56094809-2.jpg"></a>
    <a href="https://www.douban.com/people/nangongtianlun/" class="author">
      <span property="v:reviewer">南宫天伦</span>
    </a>    <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>
  <span property="v:dtreviewed" content="2016-09-06"
      class="main-meta">2016-09-06 22:54:58</span>
</div>

  </header>

        <div class="main-bd">
          <div id="review_8075045_short" class="review-short">
            <div class="short-content">
                  <p class="main-title-tip">
                    这篇影评可能有剧透
                  </p>
                  <div class="toggle_review">
                    &gt;&nbsp;                    <a href="javascript:;" id="toggle-8075045-copy" class="unfold j" title="展开全文">                      没关系，可以显示全文                    </a>
                  </div>
            </div>
          </div>
          <div id="review_8075045_full" class="hidden">
            <div id="review_8075045_full_content" class="full-content"></div>
          </div>
        </div>
    </div>
  </div>






  <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review">
    <div class="main review-item" id="8085143">


  <header class="main-hd">
    <h3 class="title">
      <a href="https://movie.douban.com/review/8085143/" class="title-link">喜大普奔！双女主终于不是撕【哔】机器人了</a>
      <div class="toggle_review right">        <a href="javascript:;" id="toggle-8085143"class="indicator unfold j" title="展开全文">        </a>
      </div>
    </h3>


<div class="header-more">
    <a class="author-avatar" href="https://www.douban.com/people/60599665/"><img width="48" height="48" src="https://img3.doubanio.com/icon/u60599665-25.jpg"></a>
    <a href="https://www.douban.com/people/60599665/" class="author">
      <span property="v:reviewer">肖恩恩恩恩肖</span>
    </a>    <span property="v:rating" class="allstar40 main-title-rating" title="推荐"></span>
  <span property="v:dtreviewed" content="2016-09-14"
      class="main-meta">2016-09-14 21:43:16</span>
</div>

  </header>

        <div class="main-bd">
          <div id="review_8085143_short" class="review-short">
            <div class="short-content">
                  鉴于今年国产电影低到离奇的影片质量，在走进电影院前，我是不惮以最坏的恶意来揣测《七月与安生》的。这部电影实在太有被黑的潜力了。  让我们划一下关键词：  安妮宝贝，小清新，国产青春片，双女主，闺蜜。  全是重灾区。  《失恋三十三天》《小时代》《致青春》三部电影全...

                    <a class="pl" href="https://movie.douban.com/review/8085143/#comments">(183回应)</a>
                  <div class="more-info pl clearfix">
                    <span class="left">2316有用 / 326没用</span>
                  </div>
            </div>
          </div>
          <div id="review_8085143_full" class="hidden">
            <div id="review_8085143_full_content" class="full-content"></div>
          </div>
        </div>
    </div>
  </div>






  <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review">
    <div class="main review-item" id="8083968">


  <header class="main-hd">
    <h3 class="title">
      <a href="https://movie.douban.com/review/8083968/" class="title-link">这不是一部传统意义上的青春片，而是一部好看的...</a>
      <div class="toggle_review right">        <a href="javascript:;" id="toggle-8083968"class="indicator unfold j" title="展开全文">        </a>
      </div>
    </h3>


<div class="header-more">
    <a class="author-avatar" href="https://www.douban.com/people/63773386/"><img width="48" height="48" src="https://img1.doubanio.com/icon/u63773386-7.jpg"></a>
    <a href="https://www.douban.com/people/63773386/" class="author">
      <span property="v:reviewer">风中芭蕾</span>
    </a>    <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>
  <span property="v:dtreviewed" content="2016-09-13"
      class="main-meta">2016-09-13 23:05:30</span>
</div>

  </header>

        <div class="main-bd">
          <div id="review_8083968_short" class="review-short">
            <div class="short-content">
                  <p class="main-title-tip">
                    这篇影评可能有剧透
                  </p>
                  <div class="toggle_review">
                    &gt;&nbsp;                    <a href="javascript:;" id="toggle-8083968-copy" class="unfold j" title="展开全文">                      没关系，可以显示全文                    </a>
                  </div>
            </div>
          </div>
          <div id="review_8083968_full" class="hidden">
            <div id="review_8083968_full_content" class="full-content"></div>
          </div>
        </div>
    </div>
  </div>






  <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review">
    <div class="main review-item" id="8087571">


  <header class="main-hd">
    <h3 class="title">
      <a href="https://movie.douban.com/review/8087571/" class="title-link">专访 | 这是一部细节控导演拍给少女的诚意电影</a>
      <div class="toggle_review right">        <a href="javascript:;" id="toggle-8087571"class="indicator unfold j" title="展开全文">        </a>
      </div>
    </h3>


<div class="header-more">
    <a class="author-avatar" href="https://www.douban.com/people/lephemera/"><img width="48" height="48" src="https://img3.doubanio.com/icon/u2744437-144.jpg"></a>
    <a href="https://www.douban.com/people/lephemera/" class="author">
      <span property="v:reviewer">蜉蝣</span>
    </a>    <span property="v:rating" class="allstar30 main-title-rating" title="还行"></span>
  <span property="v:dtreviewed" content="2016-09-16"
      class="main-meta">2016-09-16 13:30:48</span>
</div>

  </header>

        <div class="main-bd">
          <div id="review_8087571_short" class="review-short">
            <div class="short-content">
                  去看《七月与安生》的时候是稍微带着点害羞的心情的，初中时候就很喜欢这部短篇，偷偷看完安妮宝贝所有的书，总有把握不好情绪被别人视为矫情的时刻。十几年后再面对这样一部少女电影，生怕它会痛杀早被更新过的神经。  于是完全没想到，在主角说出“我恨过你，但我也只有你”...

                    <a class="pl" href="https://movie.douban.com/review/8087571/#comments">(44回应)</a>
                  <div class="more-info pl clearfix">
                    <span class="left">378有用 / 42没用</span>
                  </div>
            </div>
          </div>
          <div id="review_8087571_full" class="hidden">
            <div id="review_8087571_full_content" class="full-content"></div>
          </div>
        </div>
    </div>
  </div>





</div>










            <p class="pl" align="right">
                &gt;
                <a href="reviews">
                    更多影评3340篇
                </a>
            </p>

</section>

<script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/7d836b094e0dc536.js"></script>

    <br/>

            <h2 class="discussion_link">
                <a href="/subject/25827935/discussion/">
                    &gt; 去这部影片的讨论区（全部111条）
                </a>
            </h2>
    <script type="text/javascript">
        $(function(){if($.browser.msie && $.browser.version == 6.0){
            var $info = $('#info'),
                maxWidth = parseInt($info.css('max-width'));
            if($info.width() > maxWidth) {
                $info.width(maxWidth);
            }
        }});
    </script>


            </div>
            <div class="aside">




















<script id="episode-tmpl" type="text/x-jsrender">
<div id="tv-play-source" class="play-source">
    <div class="cross">
        <span style="color:#494949; font-size:16px">{{:cn}}</span>
        <span style="cursor:pointer">✕</span>
    </div>
    <div class="episode-list">
        {{for playlist}}
            <a href="{{:play_link}}&episode={{:ep}}" target="_blank">{{:ep}}集</a>
        {{/for}}
     <div>
 </div>
</script>

<div class="gray_ad">

    <h2>
        在哪儿看这部电影
            &nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;
    </h2>

    <a href="/subject/25827935/report_subject_error?pname=在线观看" target="_blank" class="report">报错</a>

    <ul class="bs">
                <li>
                    <a class="playBtn" data-cn="腾讯视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fv.qq.com%2Fx%2Fcover%2Fq8742fxhp6pj7zv.html%3Fptag%3Ddouban.movie&amp;subtype=1&amp;type=online-video" target="_blank">
                        腾讯视频
                    </a>
                    <span class="buylink-price"><span>
                        免费
                    </span></span>
                </li>
                <li>
                    <a class="playBtn" data-cn="爱奇艺视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fwww.iqiyi.com%2Fv_19rr9nifyg.html%3Fvfm%3Dm_331_dbdy%26fv%3D4904d94982104144a1548dd9040df241&amp;subtype=9&amp;type=online-video&amp;link2key=65c1993f40" target="_blank">
                        爱奇艺视频
                    </a>
                    <span class="buylink-price"><span>
                        付费 5 元
                    </span></span>
                </li>
                <li>
                    <a class="playBtn" data-cn="芒果 TV" href="https://www.douban.com/link2/?url=http%3A%2F%2Fwww.mgtv.com%2Fh%2F297788.html%3Fcxid%3D94l62sj44&amp;subtype=7&amp;type=online-video&amp;link2key=65c1993f40" target="_blank">
                        芒果 TV
                    </a>
                    <span class="buylink-price"><span>
                        免费
                    </span></span>
                </li>
                <li>
                    <a class="playBtn" data-cn="优酷视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fcps.youku.com%2Fredirect.html%3Fid%3D0000a213%26url%3Dhttp%3A%2F%2Fv.youku.com%2Fv_show%2Fid_XMTc3NTgxOTIxNg%3D%3D.html&amp;subtype=3&amp;type=online-video" target="_blank">
                        优酷视频
                    </a>
                    <span class="buylink-price"><span>
                        付费 20 元/月
                    </span></span>
                </li>
                <li>
                    <a class="playBtn" data-cn="乐视视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fwww.le.com%2Fptv%2Fvplay%2F26880874.html%3Fch%3Ddouban_mf&amp;subtype=6&amp;type=online-video&amp;link2key=65c1993f40" target="_blank">
                        乐视视频
                    </a>
                    <span class="buylink-price"><span>
                        免费
                    </span></span>
                </li>


    </ul>
</div>


    <!-- douban ad begin -->
    <div id="dale_movie_subject_top_right"></div>
    <div id="dale_movie_subject_top_middle"></div>
    <!-- douban ad end -->



<style>
    .qrcode-app {
        margin: 20px 0;
    }
    .qrcode-app a.download {
        display: block;
        width: 300px;
        height: 80px;
        background: url(/pics/qrcode_app4@2x.png) no-repeat;
        background-size: contain;
        background-image: -webkit-image-set(url(/pics/qrcode_app4.png) 1x, url(/pics/qrcode_app4@2x.png) 2x);
        background-image: -moz-image-set(url(/pics/qrcode_app4.png) 1x, url(/pics/qrcode_app4@2x.png) 2x);
        background-image: -ms-image-set(url(/pics/qrcode_app4.png) 1x, url(/pics/qrcode_app4@2x.png) 2x);
        background-image: -o-image-set(url(/pics/qrcode_app4.png) 1x, url(/pics/qrcode_app4@2x.png) 2x);
    }
</style>
<div class="qrcode-app">
    <a class="download" href="https://www.douban.com/doubanapp/app?channel=Db_Homepage_Bar&amp;direct_dl=1" target="_blank">
    </a>
</div>





<style type="text/css">
    .m4 {margin-bottom:8px; padding-bottom:8px;}
    .movieOnline {background:#FFF6ED; padding:10px; margin-bottom:20px;}
    .movieOnline h2 {margin:0 0 5px;}
    .movieOnline .sitename {line-height:2em; width:160px;}
    .movieOnline td,.movieOnline td a:link,.movieOnline td a:visited{color:#666;}
    .movieOnline td a:hover {color:#fff;}
    .link-bt:link,
    .link-bt:visited,
    .link-bt:hover,
    .link-bt:active {margin:5px 0 0; padding:2px 8px; background:#a8c598; color:#fff; -moz-border-radius: 3px; -webkit-border-radius: 3px; border-radius: 3px; display:inline-block;}
</style>












    <div class="tags">


    <h2>
        <i class="">豆瓣成员常用的标签</i>
              · · · · · ·
    </h2>

        <div class="tags-body">
                <a href="/tag/青春" class="">青春</a>
                <a href="/tag/友情" class="">友情</a>
                <a href="/tag/成长" class="">成长</a>
                <a href="/tag/文艺" class="">文艺</a>
                <a href="/tag/爱情" class="">爱情</a>
                <a href="/tag/女权" class="">女权</a>
                <a href="/tag/小说改编" class="">小说改编</a>
                <a href="/tag/2016" class="">2016</a>
        </div>
    </div>











<div id="subject-doulist">


    <h2>
        <i class="">以下豆列推荐</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/25827935/doulists">全部</a>
            ) </span>
    </h2>



    <ul>
            <li>
                <a href="https://www.douban.com/doulist/30299/" target="_blank">豆瓣电影【口碑榜】2016-12-18更新</a>
                <span>(影志)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/257125/" target="_blank">我爱的文艺片（3）</a>
                <span>(微笑不语)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/44231886/" target="_blank">豆瓣电影口碑榜（10月更新）</a>
                <span>(豆瓣电影)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1295618/" target="_blank">【中国内地电影票房总排行】</a>
                <span>(荔枝超人)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/34652/" target="_blank">爱·感动</a>
                <span>({Claire}在想你)</span>
            </li>
    </ul>

</div>










<div id="subject-others-interests">


    <h2>
        <i class="">谁在看这部电影</i>
              · · · · · ·
    </h2>


    <ul class="">

            <li class="">
                <a href="https://www.douban.com/people/monika_yuh/" class="others-interest-avatar">
                    <img src="https://img5.doubanio.com/icon/u1608475-6.jpg" class="pil" alt="Monika™ 晏">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/monika_yuh/" class="">Monika™ 晏</a>
                    <div class="">
                        3分钟前
                        看过
                        <span class="allstar30" title="还行"></span>
                    </div>
                </div>
            </li>

            <li class="">
                <a href="https://www.douban.com/people/44220159/" class="others-interest-avatar">
                    <img src="https://img1.doubanio.com/icon/u44220159-68.jpg" class="pil" alt="理智的疯子">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/44220159/" class="">理智的疯子</a>
                    <div class="">
                        4分钟前
                        看过
                        <span class="allstar40" title="推荐"></span>
                    </div>
                </div>
            </li>
    </ul>


    <div class="subject-others-interests-ft">

            <a href="https://movie.douban.com/subject/25827935/collections">176386人看过</a>
                &nbsp;/&nbsp;
            <a href="https://movie.douban.com/subject/25827935/wishes">19400人想看</a>
    </div>

</div>






<!-- douban ad begin -->
<div id="dale_movie_subject_middle_right"></div>
<script type="text/javascript">
    (function (global) {
        if(!document.getElementsByClassName) {
            document.getElementsByClassName = function(className) {
                return this.querySelectorAll("." + className);
            };
            Element.prototype.getElementsByClassName = document.getElementsByClassName;

        }
        var articles = global.document.getElementsByClassName('article'),
            asides = global.document.getElementsByClassName('aside');

        if (articles.length > 0 && asides.length > 0 && articles[0].offsetHeight >= asides[0].offsetHeight) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_middle_right');
        }
    })(this);
</script>
<!-- douban ad end -->



    <br/>


<p class="pl">订阅七月与安生的评论: <br/><span class="feed">
    <a href="https://movie.douban.com/feed/subject/25827935/reviews"> feed: rss 2.0</a></span></p>


            </div>
            <div class="extra">


<!-- douban ad begin -->
<div id="dale_movie_subject_bottom_super_banner"></div>
<script type="text/javascript">
    (function (global) {
        var body = global.document.body,
            html = global.document.documentElement;

        var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
        if (height >= 2000) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_bottom_super_banner');
        }
    })(this);
</script>
<!-- douban ad end -->


            </div>
        </div>
    </div>


    <div id="footer">
            <div class="footer-extra"></div>

<span id="icp" class="fleft gray-link">
    &copy; 2005－2016 douban.com, all rights reserved 北京豆网科技有限公司
</span>

<a href="https://www.douban.com/hnypt/variformcyst.py" style="display: none;"></a>

<span class="fright">
    <a href="https://www.douban.com/about">关于豆瓣</a>
    · <a href="https://www.douban.com/jobs">在豆瓣工作</a>
    · <a href="https://www.douban.com/about?topic=contactus">联系我们</a>
    · <a href="https://www.douban.com/about?policy=disclaimer">免责声明</a>

    · <a href="https://help.douban.com/?app=movie" target="_blank">帮助中心</a>
    · <a href="https://www.douban.com/doubanapp/">移动应用</a>
    · <a href="https://www.douban.com/partner/">豆瓣广告</a>
</span>

    </div>

    </div>
    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/17ba8988d5a50cc2.js"></script>


    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css" />
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/movie/7345d4b259505f1a074773ce6b54f0253df2bb88/css/movie/mod/reg_login_pop.css" />
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/3d185ca912c999ee7f464749201139ebf8eb6972/js/ui/dialog.js"></script>
    <script type="text/javascript">
        var HTTPS_DB='https://www.douban.com';
var account_pop={open:function(e,b){if(b){referrer="?referrer="+encodeURIComponent(b)}else{referrer="?referrer="+window.location.href}var f="",a="",i=382;if(e==="reg"){f="用户注册";a="https://accounts.douban.com/popup/login?source=movie#popup_register";i=480}else{if(e==="login"){f="用户登录";a="https://accounts.douban.com/popup/login?source=movie"}}var h=document.location.protocol+"//"+document.location.hostname;var d=dui.Dialog({width:478,title:f,height:i,cls:"account_pop",isHideTitle:true,modal:true,content:"<iframe scrolling='no' frameborder='0' width='478' height='"+i+"' src='"+a+"' name='"+h+"'></iframe>"},true),c=d.node;c.undelegate();c.delegate(".dui-dialog-close","click",function(){var j=$("body");j.find("#login_msk").hide();j.find(".account_pop").remove()});if($(window).width()<478){var g="";if(e==="reg"){g=HTTPS_DB+"/accounts/register"+referrer}else{if(e==="login"){g=HTTPS_DB+"/accounts/login"+referrer}}window.location.href=g}else{d.open()}$(window).bind("message",function(j){if(j.originalEvent.origin==="https://accounts.douban.com"){c.find("iframe").css("height",j.originalEvent.data);c.height(j.originalEvent.data);d.update()}})}};if(Douban&&Douban.init_show_login){Douban.init_show_login=function(b){var a=$(b);a.click(function(){var c=a.data("ref")||"";account_pop.open("login",c);return false})}}Do(function(){$("body").delegate(".pop_register","click",function(b){b.preventDefault();var a=$(this).data("ref")||"";account_pop.open("reg",a);return false});$("body").delegate(".pop_login","click",function(b){b.preventDefault();var a=$(this).data("ref")||"";account_pop.open("login",a);return false})});
    </script>










<script type="text/javascript">
    (function (global) {
        var newNode = global.document.createElement('script'),
            existingNode = global.document.getElementsByTagName('script')[0],
            adSource = '//erebor.douban.com/',
            userId = '',
            browserId = 'gFP9qSgGTfA',
            criteria = '7:马思纯|7:蔡纲|7:姚欣言|7:中国大陆|7:感动|7:剧情|7:周冬雨|7:友情|7:蒋亭轩|7:沙全泽|7:曾国祥|7:2016|7:成长|7:女性|7:小说改编|7:爱情|7:蒙亭宜|7:文艺|7:李萍|7:青春|7:李程彬|7:李昊芳|3:/subject/25827935/',
            preview = '',
            debug = false,
            adSlots = ['dale_movie_subject_top_icon', 'dale_movie_subject_top_right', 'dale_movie_subject_top_middle'];

        global.DoubanAdRequest = {src: adSource, uid: userId, bid: browserId, crtr: criteria, prv: preview, debug: debug};
        global.DoubanAdSlots = (global.DoubanAdSlots || []).concat(adSlots);

        newNode.setAttribute('type', 'text/javascript');
        newNode.setAttribute('src', 'https://img3.doubanio.com/f/shire/99a1c905fb5a6d46d8e742d60199e3a99414a580/js/ad.js');
        newNode.setAttribute('async', true);
        existingNode.parentNode.insertBefore(newNode, existingNode);
    })(this);
</script>






















<script type="text/javascript">
var _paq = _paq || [];
_paq.push(['trackPageView']);
_paq.push(['enableLinkTracking']);
(function() {
    var p=(('https:' == document.location.protocol) ? 'https' : 'http'), u=p+'://fundin.douban.com/';
    _paq.push(['setTrackerUrl', u+'piwik']);
    _paq.push(['setSiteId', '100001']);
    var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
    g.type='text/javascript';
    g.defer=true;
    g.async=true;
    g.src=p+'://img3.doubanio.com/dae/fundin/piwik.js';
    s.parentNode.insertBefore(g,s);
})();
</script>

<script type="text/javascript">
var setMethodWithNs = function(namespace) {
  var ns = namespace ? namespace + '.' : ''
    , fn = function(string) {
        if(!ns) {return string}
        return ns + string
      }
  return fn
}

var gaWithNamespace = function(fn, namespace) {
  var method = setMethodWithNs(namespace)
  fn.call(this, method)
}

var _gaq = _gaq || []
  , accounts = [
      { id: 'UA-7019765-1', namespace: 'douban' }
    , { id: 'UA-7019765-19', namespace: '' }
    ]
  , gaInit = function(account) {
      gaWithNamespace(function(method) {
        gaInitFn.call(this, method, account)
      }, account.namespace)
    }
  , gaInitFn = function(method, account) {
      _gaq.push([method('_setAccount'), account.id]);
      _gaq.push([method('_setSampleRate'), '5']);


  _gaq.push([method('_addOrganic'), 'google', 'q'])
  _gaq.push([method('_addOrganic'), 'baidu', 'wd'])
  _gaq.push([method('_addOrganic'), 'soso', 'w'])
  _gaq.push([method('_addOrganic'), 'youdao', 'q'])
  _gaq.push([method('_addOrganic'), 'so.360.cn', 'q'])
  _gaq.push([method('_addOrganic'), 'sogou', 'query'])
  if (account.namespace) {
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣'])
    _gaq.push([method('_addIgnoredOrganic'), 'douban'])
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣网'])
    _gaq.push([method('_addIgnoredOrganic'), 'www.douban.com'])
  }

      if (account.namespace === 'douban') {
        _gaq.push([method('_setDomainName'), '.douban.com'])
      }

        _gaq.push([method('_setCustomVar'), 1, 'responsive_view_mode', 'desktop', 3])

        _gaq.push([method('_setCustomVar'), 2, 'login_status', '0', 2]);

      _gaq.push([method('_trackPageview')])
    }

for(var i = 0, l = accounts.length; i < l; i++) {
  var account = accounts[i]
  gaInit(account)
}


;(function() {
    var ga = document.createElement('script');
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    ga.setAttribute('async', 'true');
    document.documentElement.firstChild.appendChild(ga);
})()
</script>








    <!-- sindar17b-docker-->

  <script>_SPLITTEST=''</script>
</body>

</html>

`

	this.Ctx.WriteString("<h1>电影名称:" + models.GetMovieName(sMovieHtml) + "</h1><br/>")
	this.Ctx.WriteString("<h2>电影导演:" + models.GetMovieDirector(sMovieHtml) + "</h2><br/>")
	this.Ctx.WriteString("<h2>电影主演:" + models.GetMovieMainCharacters(sMovieHtml) + "</h2><br/>")
	this.Ctx.WriteString("<h2>豆瓣评分:" + models.GetMovieGrade(sMovieHtml) + "</h2><br/>")
	this.Ctx.WriteString("<h2>电影类型:" + models.GetMovieGenre(sMovieHtml) + "</h2><br/>")
	this.Ctx.WriteString("<h2>上映日期:" + models.GetMovieOnTime(sMovieHtml) + "</h2><br/>")
	this.Ctx.WriteString("<h2>电影片长:" + models.GetMovieRunningTime(sMovieHtml) + "</h2><br/>")

}
