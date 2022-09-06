package test

import (
	"bytes"
	"context"
	"fmt"
	"gcloud/core/define"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse("https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "gcloud/exampleobject.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/1ff6a037-409d-445a-86cc-6dbca2b29c87.jpeg", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "gcloud/exampleobject2.jpg"

	f, err := os.ReadFile("./img/1ff6a037-409d-445a-86cc-6dbca2b29c87.jpeg")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "gcloud/exampleobject.jpeg"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}

	UploadID := v.UploadID // 1653047261484591ac09d1e16e24bc593154fe610f19aa6d43d475b1e3cbdc030bbc6519af
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "gcloud/exampleobject.jpeg"
	UploadID := "1653047261484591ac09d1e16e24bc593154fe610f19aa6d43d475b1e3cbdc030bbc6519af"
	f, err := os.ReadFile("0.chunk") // md5 : 108e92d35fe1695fbf29737d0b24561d
	if err != nil {
		t.Fatal(err)
	}

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag") // md5
	fmt.Println(PartETag)               // 108e92d35fe1695fbf29737d0b24561d
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "gcloud/exampleobject.jpeg"
	UploadID := "1653047261484591ac09d1e16e24bc593154fe610f19aa6d43d475b1e3cbdc030bbc6519af"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "108e92d35fe1695fbf29737d0b24561d"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCosDownload(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := define.CosFolderName + "/" + "01e90e8c-94a5-4ef7-9374-89b54770eb10.jpg"
	// resp, err := client.Object.Get(context.Background(), key, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// bs, _ := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// fmt.Printf("%s\n", string(bs))

	_, err := client.Object.GetToFile(context.Background(), key, "01e90e8c-94a5-4ef7-9374-89b54770eb10.jpg", nil)
	if err != nil {
		panic(err)
	}
	// 返回文件链接
}
