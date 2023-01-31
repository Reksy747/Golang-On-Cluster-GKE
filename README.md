
# Ci/Cd Simple Website Menggunakan Golang,Docker,dan GKE

Dalam proyek ini saya membuat aplikasi golang yang menampilkan "helloworld" pada path ("/") dengan menggunakan library dari mux dan juga menampilkan halaman portofolio di path ("/portofolio/") dan fungsi github Action golang yang dibuat memiliki fungsi tes yang akan menguji koneksi pada setiap path yang saya buat sebelumnya.



# Demo
- setup enviroment kubernetes disini saya menggunakann GKE untuk menjalankan cluster dan simpan di secret github dan juga simpan secret dari nama project yang di buat kedalam secret
    - membuat cluster 

```shell
 gcloud container clusters create $GKE_CLUSTER \
	--project=$GKE_PROJECT \
	--zone=$GKE_ZONE
```
   - menyalakan api gcloud

```shell
gcloud services enable \
	containerregistry.googleapis.com \
	container.googleapis.com
```
   - membuat user untuk menjalankan github action

```shell
gcloud iam service-accounts create $SA_NAME

gcloud iam service-accounts list

gcloud projects add-iam-policy-binding $GKE_PROJECT \
	--member=serviceAccount:$SA_EMAIL \
	--role=roles/container.admin
gcloud projects add-iam-policy-binding $GKE_PROJECT \
	--member=serviceAccount:$SA_EMAIL \
	--role=roles/storage.admin
gcloud projects add-iam-policy-binding $GKE_PROJECT \
	--member=serviceAccount:$SA_EMAIL \
	--role=roles/container.clusterViewer

gcloud iam service-accounts keys create key.json --iam-account=$SA_EMAIL

export GKE_SA_KEY=$(cat key.json | base64)

echo $GKE_SA_KEY
```   
   - kita setup file yaml awal disini saya akan apply file yaml deploy.yaml dan loadbalancer.yaml dan akan berjalan pada namespace staging

```shell
kubectl create ns staging
kubectl apply -f deploy.yaml
kubectl apply -f loadbalancer.yaml
```
    
   - lalu kita cek apakah sudah berjalan dengan perintah 

```shell
kubectl get services --field-selector metadata.name=reksy-web
```

jika bagian cloud gke sudah di setup maka kita akan langsung ke alur Ci/Cd

- Saat aplikasi push ke Github, aplikasi akan memanggil Github Action dan membaca alur kerja kode dari file ci-cd.yml.

- Alur kerja file ci-cd.yaml menentukan platform yang akan di jalankan, dalam hal ini lingkungan berjalan di Ubuntu misalnya dan memicu fungsi untuk checkout github.

- Untuk menunjukkan versi golang yang akan digunakan di sini, proyek yang saya buat bekerja dengan golang versi 1.19.

- workflow menjalankan fungsi "go build -v ./.." dan memanggil fungsi main.go dan akan mengunduh repositori yang digunakan dalam aplikasi yang saya gunakan setelah membangun aplikasi ini dan juga menjalankan perintah "go test ". - v ./." lalu golang akan mencari file test case yang sebelumnya dibuat pada file main_test.go.

- workflow menjalankan perintah untuk masuk ke Docker Hub. Di sini saya menggunakan github secret untuk menyembunyikan file nama pengguna dan kata sandi saya.

-  Dockerfile yang dibuat sebelumnya akan di build lagi dan akan menjadi file "main" yang akan di pindahkan ke dalam /app dan juga akan mengcopy beberapa folder termaksud folder/portofolio/ jika docker sudah di build dengan sukses makan akan menjalankan fungsi docker push ke docker hub milik saya https://hub.docker.com/u/reksy737

- workflow akan menjalankan fungsi untuk login pada gcloud cli dan akan mejalankan fungsi mendapatkan fungsi GKE credential yang telah di simpan pada secret github sehingga workflow dapat menjalankan fungsi kubectl

- workflow akan menjalankan fungsi untuk menganti image docker yang telah terakhir di build sehingga kita tidak perlu untuk apply yaml lagi jika ada perubahan code pada aplikasi golang yang kita build kedepannya

```shell
kubectl set image deployment/web-reksy app=reksy737/simpelgoweb:build${{github.run_number}}  -n staging
```

- pada tahap terakhir ini kita akan cek apakah website telah berjalan dengan menggunakan dns dari sslip.io dengan menggunakan perintah pada workflow
```shell
EXTERNAL_IP=$(kubectl get service reksy-web -n staging -ojsonpath="{.status.loadBalancer.ingress[0].ip}")
DOMAIN="${EXTERNAL_IP}.sslip.io"
curl $DOMAIN
echo $DOMAIN
```

