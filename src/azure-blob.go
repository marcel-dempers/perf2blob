package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"github.com/Azure/azure-storage-blob-go/azblob"
)

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok { // This error is a Service-specific
			switch serr.ServiceCode() { // Compare serviceCode to ServiceCodeXxx constants
			case azblob.ServiceCodeContainerAlreadyExists:
				fmt.Println("Received 409. Container already exists")
				return
			}
		}
		log.Fatal(err)
	}
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func UploadFiles() {

	accountName, accountKey,containerName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"), os.Getenv("AZURE_STORAGE_ACCOUNT_KEY"), os.Getenv("AZURE_STORAGE_CONTAINER")
	if len(accountName) == 0 || len(accountKey) == 0 || len(containerName) == 0 {
		log.Fatal("Either the AZURE_STORAGE_ACCOUNT_NAME or AZURE_STORAGE_ACCOUNT_KEY or AZURE_STORAGE_CONTAINER environment variable is not set")
	}

	// Create a default request pipeline using your storage account name and account key.
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", accountName))
	serviceURL := azblob.NewServiceURL(*u, p)
	ctx := context.Background()

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL := serviceURL.NewContainerURL(containerName)

	// Create the container
	fmt.Printf("Creating a container named %s\n", containerName)
	 // This example uses a never-expiring context
	_, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	handleErrors(err)

	fileName := "/out/perf.data"
	check(err)
	
	blobURL := containerURL.NewBlockBlobURL(fileName)
	file, err := os.Open(fileName)
	handleErrors(err)

	fmt.Printf("Uploading the file with blob name: %s\n", fileName)
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
	handleErrors(err)

}
