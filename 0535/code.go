package _535
//15:00
type Codec struct {
	storage map[string]string
	id int
}


func Constructor() Codec {
	return Codec{
		storage: map[string]string{},
		id: 0,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	url := ""
	for i := 0; i < len(longUrl); i++ {
		if longUrl[i] == '/'{
			url += "/" + string(this.id)
			this.id++
			break
		}
		url += string(longUrl[i])
	}
	this.storage[url] = longUrl
	return url
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.storage[shortUrl]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
