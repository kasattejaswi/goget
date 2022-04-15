package downloader

import (
	"fmt"
	"github.com/gosuri/uilive"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type DownloadOptions struct {
	Url      string
	Threads  int
	Output   string
	FileName string
}

func (d *DownloadOptions) Download() error {
	rs, cl, err := isRangeSupported(d.Url)
	if err != nil {
		return err
	}
	if rs {
		fmt.Println("File size in bytes:", cl)
		var wg sync.WaitGroup
		eachChunk := cl / d.Threads
		lastChunk := cl % d.Threads
		var startByte int
		var endByte int
		for i := 1; i <= d.Threads; i++ {
			if i == 1 {
				endByte = eachChunk
			} else {
				startByte = eachChunk*(i-1) + 1
				endByte = eachChunk * i
			}
			wg.Add(1)
			filePath := filepath.Join(d.Output, d.FileName+"."+strconv.Itoa(endByte))
			go func() {
				err := downloadFileForRange(&wg, filePath, d.Url, strconv.Itoa(startByte), strconv.Itoa(endByte))
				if err != nil {
					fmt.Println(err)
					return
				}
			}()
			if err != nil {
				return err
			}
		}
		if lastChunk != 0 {
			wg.Add(1)
			filePath := filepath.Join(d.Output, d.FileName+"."+strconv.Itoa(endByte+lastChunk))
			go func() {
				err := downloadFileForRange(&wg, filePath, d.Url, strconv.Itoa(endByte+1), strconv.Itoa(endByte+lastChunk))
				if err != nil {
					fmt.Println(err)
					return
				}
			}()
			if err != nil {
				return err
			}
		}
		wg.Wait()
	}
	return nil
}

func downloadFileForRange(wg *sync.WaitGroup, path string, url string, r1 string, r2 string) error {
	defer wg.Done()
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	fmt.Fprintln(writer, fmt.Sprintf("Downloading for range %v to range %v", r1, r2))
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	request, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return err
	}
	request.Header.Add("Range", fmt.Sprintf("bytes=%v-%v", r1, r2))
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintln(writer, fmt.Sprintf("Download compleated for range %v to range %v bytes", r1, r2))
	return nil
}

func isRangeSupported(url string) (bool, int, error) {
	request, err := http.NewRequest("HEAD", url, strings.NewReader(""))
	if err != nil {
		return false, 0, err
	}
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return false, 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	fmt.Println(res.StatusCode == 200)
	if res.StatusCode != 200 && res.StatusCode != 206 {
		return false, 0, fmt.Errorf("response code is not 200 or 206")
	}
	clh := res.Header.Get("Content-Length")
	cl, err := strconv.Atoi(clh)
	if err != nil {
		return false, 0, err
	}
	if res.Header.Get("Accept-Ranges") == "bytes" {
		return true, cl, nil
	}
	return false, cl, nil
}

func (d *DownloadOptions) DownloadWithoutProgress() error {
	filePath := filepath.Join(d.Output, d.FileName)
	fmt.Println(d.FileName)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)
	res, err := http.Get(d.Url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", res.Status)
	}
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}
