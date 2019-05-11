package service

import (
	"agfun/file/dbcentral/etcd"
	"agfun/file/dbcentral/pg"
	"agfun/video/dto"
	entity2 "agfun/video/entity"
	"bytes"
	"conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kataras/iris"
	"io"
	"mime/multipart"
	"os"
	"service"
	"strings"
	"util"
)

type FileSvc struct {
	service.Svc
	SysDB   *pg.SysDB
	AuthDB  *pg.AuthDB
	Dynamic *etcd.Client
}

func NewFileSvc() *FileSvc {
	svc := &FileSvc{
		Svc: *service.DefaultSvc(),
		//SysDB: &pg.SysDB{SysDB: *pg2.DefaultSysDB()},
		//AuthDB:  &pg.AuthDB{AuthDB: *pg2.DefaultAuthDB()},
		//Dynamic: &etcd.Client{Client: *etcd2.DefaultCli()},
	}
	return svc
}

func (s *FileSvc) AddVideo(ctx iris.Context) {

	//
	// UploadFormFiles
	// uploads any number of incoming files ("multiple" property on the form input).
	//

	// The second, optional, argument
	// can be used to change a file's name based on the request,
	// at this example we will showcase how to use it
	// by prefixing the uploaded file with the current user's ip.
	//ctx.UploadFormFiles("C:/Users/TITAN/Desktop/workspace/go-note/src/agfun/file/assets", beforeSave)

	// Get the file from the request.
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename

	typ := ctx.FormValue("type")

	// 写入文件系统
	dir := ""
	if typ == "电影" {
		dir = "./file/assets/videos/movies/"
	} else if typ == "电视剧" {
		dir = "./file/assets/videos/tvs/"
	} else if typ == "动漫" {
		dir = "./file/assets/videos/cartoons/"
	} else {
		dir = "./file/assets/videos/short-videos/"
	}

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(dir+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	//向video服务发送添加视频
	after := strings.SplitAfter(dir+fname, "assets")
	url := ""
	if len(after) == 2 {
		url = fmt.Sprintf("File:%s", after[1])
	}
	meta := dto.Video{
		Video: entity2.Video{
			ID:       util.NewUUID(),
			Name:     ctx.FormValue("name"),
			Url:      url,
			Describe: ctx.FormValue("describe"),
			Thumb:    ctx.FormValue("thumb"),
			Creator:  "",
			CreateAt: util.TimeNowStd(),
		},
	}
	client := &http.Client{}
	buf, e := json.Marshal(meta)
	if e != nil {
		util.Fail(ctx, e)
		return
	}
	videoReader := bytes.NewBuffer(buf)
	req, _ := http.NewRequest("POST", conf.AgfunInst().VideoHost+"/videos", videoReader)
	req.Header.Set("Content-type", "application/json")
	res, e := client.Do(req)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	if e != nil {
		util.Fail(ctx, e)
		return
	}
	util.Success(ctx, &meta)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}
