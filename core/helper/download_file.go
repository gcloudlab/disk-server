package helper

import (
	"context"
	"gcloud/core/define"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func CosDownload(r *http.Request, dowload_path string, fileName string) ([]byte, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := define.CosFolderName + "/" + dowload_path[61:]
	resp, err := client.Object.Get(context.Background(), key, nil)
	if err != nil {
		panic(err)
	}

	bs, _ := ioutil.ReadAll(resp.Body)

	// _, err = client.Object.GetToFile(context.Background(), key, fileName, nil)
	// if err != nil {
	// 	panic(err)
	// }

	return bs, err
}
