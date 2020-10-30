package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"path"
	"os"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
    r := gin.Default()
    //限制上传最大尺寸
    r.MaxMultipartMemory = 8 << 20
    r.POST("/upload", func(c *gin.Context) {
		apptype := c.DefaultQuery("app", "go") //APP文件类型
        file, err := c.FormFile("file")
        if err != nil {
            c.String(500, "上传图片出错")
		}
		//服务器图片地址 /mnt/images
		// savePath := "/mnt/images/"+apptype
		savePath := "./"+apptype
		baseUrl := "http://images.tomcat97.cn/"+apptype+"/"+file.Filename
		exist,err := PathExists(savePath)
		if err !=nil {
			c.JSON(500, gin.H{"message": "error","app":apptype,"url":baseUrl,"path":savePath})
		}
		if !exist {
			err := os.Mkdir(savePath,os.ModePerm)
			if err !=nil{
				fmt.Printf("mkdir failed![%v]\n", err)
			}else{
				fmt.Printf("mkdir success!\n")
			}
		}
		dst := path.Join(savePath, file.Filename)
        // c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, dst)
	
		if  http.StatusOK == 200 {
			c.JSON(200, gin.H{"message": "success","app":apptype,"url":baseUrl,"path":savePath})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":"error","app":apptype,"url":baseUrl})
		}
        //c.String(http.StatusOK, file.Filename)
	})
	
	r.GET("/images", func(c *gin.Context) {
		savePath := "."
		files := scanDir(savePath)
		var fileStr string
		for _,file := range files{
			fileStr += file+"</br>"
		}
		c.Header("Content-Type","text/html;charset=utf-8")

        c.String(http.StatusOK, fileStr)
    })
    r.Run(":8001")
}

//判断文件夹是否存在
func PathExists(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil{
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}

//扫描当前目录下的文件，不递归扫描
func scanDir(dirName string) []string{
	files,err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}
	var fileList []string
	for _,file := range files{
		// fileList = append(fileList,dirName + string(os.PathSeparator)+file.Name())
		if file.IsDir(){ //判断是否是目录
			fileList = append(fileList,scanDir(dirName + string(os.PathSeparator)+file.Name())...)
		}else{
			fileList = append(fileList,dirName + string(os.PathSeparator)+file.Name())
		}
	}
	return fileList
}
