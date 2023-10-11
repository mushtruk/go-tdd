
package concurrency

import (
	"reflect"
	"testing"
	"time"
)

var websites = []string{
	"http://google.com",
	"http://blog.gypsydave5.com",
	"waat://furhurterwe.geds",
	}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}


func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T)  {
	want := map[string]bool{
		"http://google.com" : 			true,
		"http://blog.gypsydave5.com" : 	true,
		"waat://furhurterwe.geds": 		false,
	}
	
	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
