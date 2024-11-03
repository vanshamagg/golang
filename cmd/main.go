package main

import "random/channels"

// s "random/structures"
// "random/goroutines"

// func main() {

// 	cfg, err := config.LoadDefaultConfig(context.TODO())

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := s3.NewFromConfig(cfg)

// 	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
// 		Bucket: aws.String("dev-cropwise-planting-files"),
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, object := range output.Contents {
// 		log.Printf("key=%s, size=%d", aws.ToString(object.Key), object.Size)
// 	}

// 	object, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
// 		Bucket: aws.String("dev-cropwise-planting-files"),
// 		Key: aws.String("ECv2/test.csv"),
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	buf := make([]byte, 2048)

// 	for {
// 		n, err := object.Body.Read(buf)

// 		if err != nil && err != io.EOF {
// 			log.Fatal(err)
// 		}

// 		if  err == io.EOF {
// 			log.Println("End of file baba")
// 			break
// 		}

//			fmt.Print(string(buf[:n]))
//		}
//	}
// func main() {
// 	user := s.CreateUser(&s.User{
// 		FirstName: "Vansham",
// 		LastName:  "Aggarwal",
// 		Address:   "Amritsar, Punjab",
// 		Age:       27,
// 	})

//   log.Println(user)
// }

func main() {
	channels.BufferedChannelExample()
}
