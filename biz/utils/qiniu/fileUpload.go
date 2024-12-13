package qiniu

import (
	"ARPostcard_server/biz/conf"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"path/filepath"
)

// uploadFileToQiniu 文件上传到七牛云
func uploadFileToQiniu(fileHeader *multipart.FileHeader) (string, string, error) {
	qiniuConf := conf.Conf.Qiniu
	mac := qbox.NewMac(qiniuConf.AccessKey, qiniuConf.SecretKey)
	bucket := qiniuConf.Bucket

	// 生成上传策略
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		Expires:    3600, // 1 小时
		ReturnBody: `{"key": $(key),"hash": $(etag)}`,
	}
	upToken := putPolicy.UploadToken(mac)

	// 生成唯一文件名
	ext := filepath.Ext(fileHeader.Filename)
	uniqueFileName := uuid.New().String() + ext

	// 配置存储区域
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南广东地区
		UseHTTPS:      true,
		UseCdnDomains: false,
	}

	// 创建表单上传器
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		return "", "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 执行上传
	err = formUploader.Put(context.Background(), &ret, upToken, uniqueFileName, file, fileHeader.Size, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to upload file to Qiniu: %w", err)
	}

	// 返回文件 Key 或 URL
	fileURL := fmt.Sprintf("http://%s/%s", qiniuConf.CDN, ret.Key)
	return fileURL, uniqueFileName, nil
}
