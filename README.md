
# Ci/Cd Simple Website Menggunakan Golang,Docker,dan GKE

Dalam proyek ini saya membuat aplikasi golang yang menampilkan "helloworld" pada path ("/") dengan menggunakan library dari mux dan juga menampilkan halaman portofolio di path ("/portofolio/") dan fungsi github Action golang yang dibuat memiliki fungsi tes yang akan menguji koneksi pada setiap path yang saya buat sebelumnya.



# Demo
- Saat aplikasi push ke Github, aplikasi akan memanggil Github Action dan membaca alur kerja kode dari file ci-cd.yml.

- Alur kerja file ci-cd.yaml menentukan platform yang akan di jalankan, dalam hal ini lingkungan berjalan di Ubuntu misalnya dan memicu fungsi untuk checkout github.

- Untuk menunjukkan versi golang yang akan digunakan di sini, proyek yang saya buat bekerja dengan golang versi 1.19.

- Pada langkah ini, alur kerja menjalankan fungsi "go build -v ./.." dan memanggil fungsi main.go dan akan mengunduh repositori yang digunakan dalam aplikasi yang saya gunakan setelah membangun aplikasi ini dan juga menjalankan perintah "go test ". - v ./." lalu golang akan mencari file test case yang sebelumnya dibuat pada file main_test.go.

- Pada titik ini, alur kerja menjalankan perintah untuk masuk ke Docker Hub. Di sini saya menggunakan github secret untuk menyembunyikan file nama pengguna dan kata sandi saya.

- Pada titik ini, Dockerfile bulid yang dibuat sebelumnya dan Dockerfile akan segera dieksekusi. /app dan juga bangun proyek golang sebagai "utama" yang dapat dieksekusi dan

