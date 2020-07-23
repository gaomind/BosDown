package main

import (
	"flag"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
)

func main()  {
	var start string
	var end string

	flag.StringVar(&start,"s","2020","起始时间")
	flag.StringVar(&end,"e","2021","结束时间")
	flag.Parse()
	fmt.Printf("起始时间=%v,结束时间=%v",start,end)


	ACCESS_KEY_ID, SECRET_ACCESS_KEY := "test","test"
	ENDPOINT := "https://bj.bcebos.com"
	bosClient, err := bos.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)

	if err !=nil{
		//错误处理
		fmt.Println(string("******"))

	}
	args :=new(api.ListObjectsArgs)

	args.Prefix= "/"
	args.MaxKeys =500
//	listObjectResult,err:=bosClient.ListObjects("test-vehicle01-cn",args)

	result :=new(api.ListObjectsResult)
	result,err=bosClient.ListObjects("test-vehicle01-cn",args)



	for _, obj := range result.Contents {
		fmt.Println("Key:", obj.Key, ", ETag:", obj.ETag, ", Size:", obj.Size,
			", LastModified:", obj.LastModified, ", StorageClass:", obj.StorageClass)
	}








}