package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	// inisialisasi router menggunakan fungsi konstruktor
	r := newRouter()

	// Buat server baru menggunakan metode `NewServer` dengan library "httptest".
	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	// Server mock yang kita buat menjalankan server dan memperlihatkan lokasinya di
	// atribut URL
	// dan membuat permintaan GET ke rute "/" yang tentukan di router
	resp, err := http.Get(mockServer.URL + "/")

	// Handle jika terdapat error
	if err != nil {
		t.Fatal(err)
	}

	// mengecek apakah status 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// Dalam beberapa baris berikutnya, body merespons read, dan diubah menjadi string
	defer resp.Body.Close()
	// membaca body menjadi sekelompok bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// mengubah bytes ke string
	respString := string(b)
	expected := "Hello World!"

	// jika respons cocok dengan yang ditentukan dalam handler.
	// dan outputnya ternyata "Helloworld", maka itu berarti, bahwa
	// rute sesuai
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	//cek route yang dimana kita definisikan, seperti `POST /` route.
	resp, err := http.Post(mockServer.URL+"/", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// jika ingin status 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// Kode untuk menguji body jika ternyata body kosong
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// hit url ke `GET /portofolio/` untuk mencari index.html
	resp, err := http.Get(mockServer.URL + "/portofolio/")
	if err != nil {
		t.Fatal(err)
	}

	// cek apakah status 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// menguji bahwa header tipe konten adalah "text/html; charset=utf-8"
	// agar tahu bahwa file html telah ditampilkan
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}
