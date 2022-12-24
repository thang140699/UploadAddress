package ultilities

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

// env
func Upload(fileInput []byte, fileName string, flName string) error {

	id := uuid.New()
	pathJSON := "/home/thang/.local/share/Trash/files/Wedding_Utilities/utilities/firebase/config.json"
	configStorage := &firebase.Config{
		StorageBucket: "wedding-event-5e665.appspot.com",
	}
	opt := option.WithCredentialsFile(pathJSON)
	app, err := firebase.NewApp(context.Background(), configStorage, opt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	ctx := context.Background()
	client, err := app.Storage(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return err
	}
	var fleName string
	if strings.Split(flName, "/")[0] == "flie" {
		fleName = "Image/"
	} else {
		fleName = "Txt"
	}

	docName := fleName + fileName
	object := bucket.Object(docName)
	writer := object.NewWriter(ctx)

	//Set the attribute
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(fileInput)); err != nil {
		return err
	}
	_ = object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader)

	return nil
}
