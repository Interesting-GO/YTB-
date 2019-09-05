package test

import (
	"github.com/Interesting-GO/youtubetools/video_dow"
	"testing"
)

func TestDow(t *testing.T) {
	dow := video_dow.YoutubeDow("https://www.youtube.com/watch?v=kkHohsnL8G0", "hello.mp4", "127.0.0.1:8001")
	if dow != nil {
		panic(dow)
	}
}
