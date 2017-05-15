package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/ftp"
)

func ISOImageUpload(ctx Context, params *UploadISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", e)
	}

	// open FTP
	ftpServer, err := api.OpenFTP(p.ID, false)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// upload
	ftpsClient := ftp.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	compChan := make(chan bool)
	errChan := make(chan error)

	spinner := internal.NewProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
		fmt.Sprintf("Upload ISOImage[ID:%d]", params.Id),
		GlobalOption.Progress)
	go func() {
		spinner.Start()
		err = ftpsClient.Upload(params.GetIsoFile())
		if err != nil {
			errChan <- err
		}
		compChan <- true
	}()
upload:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break upload
		case err := <-errChan:
			return fmt.Errorf("ISOImageUpload is failed: %s", err)
		}
	}

	// close FTP
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// read again
	p, err = api.Read(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}
	return ctx.GetOutput().Print(p)
}
