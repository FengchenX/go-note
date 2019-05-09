package service

import (
	"github.com/kataras/iris"
	"mime/multipart"
	"service"
	"strings"
)

const maxSize = 1 << 30 // 5MB

type Router struct {
	service.Router
	Svr *FileSvc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Party(path string) iris.Party {
	return r.Svr.Party(path)
}

func (r *Router) Handle() {
	root := r.Party("/files")

	upload := root.Party("/upload")
	upload.Post("/upload", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files ("multiple" property on the form input).
		//

		// The second, optional, argument
		// can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		ctx.UploadFormFiles("./uploads", beforeSave)
	})
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
