package main

import (
	"fmt"
	"strconv"
)

type Codec struct {
	TinyURL map[string]string
	target  int
}

func Constructor() Codec {
	return Codec{
		make(map[string]string), 1,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	shortUrl := "http://tinyurl.com/" + strconv.Itoa(this.target)
	this.TinyURL[shortUrl] = longUrl
	this.target++
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.TinyURL[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func main() {
	longUrl := "https://baidu.com/lkjgwoeoiweg"
	obj := Constructor()
	url := obj.encode(longUrl)
	ans := obj.decode(url)
	fmt.Println(url, ans)
	url = obj.encode(longUrl)
	ans = obj.decode(url)
	fmt.Println(url, ans)
	url = obj.encode(longUrl)
	ans = obj.decode(url)
	fmt.Println(url, ans)
}
