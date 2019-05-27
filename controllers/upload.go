package controllers

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"
	"wt-app-api/common"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Saveimg() {
	fmt.Println(111)
}

func (u *UploadController) Test() {
	file, header, err := u.GetFile("myfile")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	ext := path.Ext(header.Filename)
	var allowExtMap map[string]bool = map[string]bool{
		// ".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := allowExtMap[ext]; !ok {
		u.Data["json"] = common.Success("后缀名不符合上传要求")
		// u.Ctx.WriteString("后缀名不符合上传要求")
		u.ServeJSON()
	}
	u.Ctx.WriteString("后缀名不符合上传要求111111")
	fmt.Println(123456)
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006-01/02/")
	err1 := os.MkdirAll(uploadDir, 777)
	if err1 != nil {
		u.Ctx.WriteString(fmt.Sprintf("%v \n", err1))
		return
	}
	// 构造文件
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2016_01_02_15_04_05" + randNum)))
	fileName := fmt.Sprintf("%x", hashName) + ext
	fpath := uploadDir + fileName
	err = u.SaveToFile("myfile", fpath)
	if err != nil {
		u.Ctx.WriteString(fmt.Sprintf("%v", err))
	}
	u.Ctx.WriteString("上传成功")
}
