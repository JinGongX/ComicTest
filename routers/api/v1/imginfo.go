package v1

import (
	"fmt"
	"gocmictest/pkg/e"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"bytes"
	"encoding/json"
)

// 上传文件
func Uploadimgfile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败："+err.Error())
		return
	}
	var uploadir string
	uploadir = "./files/"
	_, err = os.Stat(uploadir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadir, os.ModePerm)
	}

	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := uploadir + newFileName
	uplouderr := c.SaveUploadedFile(file, dst)

	if uplouderr != nil {
		fmt.Println(uplouderr)
		c.JSON(500, gin.H{"err": "file saving failed:" + uplouderr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "suceess",
		"data": map[string]interface{}{
			"url":  "http://127.0.0.1:2025/files/" + newFileName,
			"name": newFileName, //file.Filename,
			"size": file.Size,
		},
	})

}

func SendComic(c *gin.Context) {
	var requestData struct {
		Filename string `json:"filename"`
	}

	// 解析 JSON 数据
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "JSON 解析失败: " + err.Error()})
		return
	}

	//fmt.Println("接收到的文件名:", requestData.Filename)
	log.Printf("err.key: %s, filename: %s", requestData.Filename, requestData.Filename)
	data := map[string]string{
		"arg1": requestData.Filename,
		//"arg2": "dog",
	}

	// 序列化 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	log.Printf("err.key: %s, filename: %s", jsonData, jsonData)

	// 发起 HTTP POST 请求
	resp, err := http.Post("http://localhost:5600/runcomicimage", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(500, gin.H{"err": "file saving failed:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  result["status"],
		"data": "http://127.0.0.1:2025/files/output/" + result["result"],
	})

}
