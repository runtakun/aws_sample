package main

import "fmt"
import "github.com/aws/aws-sdk-go/aws"
//import "github.com/aws/aws-sdk-go/service/s3"
import "github.com/aws/aws-sdk-go/service/kinesis"
import "github.com/aws/aws-sdk-go/aws/awsutil"
import "bytes"
import "gopkg.in/vmihailenco/msgpack.v2"
import "code.google.com/p/go-uuid/uuid"
import "net/http"

func main() {
  config := aws.NewConfig().WithRegion("ap-northeast-1").WithHTTPClient(&http.Client{})
/*
  svc := s3.New(config)

  var params *s3.ListBucketsInput
  resp, err := svc.ListBuckets(params)
*/
  svc := kinesis.New(config)

  message := map[string]interface{}{
    "id":           uuid.NewRandom().String(),
  }

  buf := &bytes.Buffer{}
  _ = msgpack.NewEncoder(buf).Encode(message)

  record := &kinesis.PutRecordsRequestEntry{
    Data:         buf.Bytes(),
    PartitionKey: aws.String("partition_key-1"),
  }

  var records []*kinesis.PutRecordsRequestEntry
  records = append(records, record)

  input := &kinesis.PutRecordsInput{
    Records:    records,
    StreamName: aws.String("sample-stream"),
  }

  resp, err := svc.PutRecords(input)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(awsutil.Prettify(resp))
}
