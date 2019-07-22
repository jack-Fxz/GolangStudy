package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	resp,err:=http.Get("http://api.kuaidihelp.com")
	if err!=nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	fmt.Println(resp.Header)
	buf:=make([]byte,4096)
	for  {
		n,err:=resp.Body.Read(buf)
		fmt.Print(string(buf[:n]))
		if err!=nil {
			if err==io.EOF {
				break
			} else {
				fmt.Print(err)
			}
		}
	}
}