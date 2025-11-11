### Ambil data dari postgres
![alt text](image.png)
saat redis tidak menyimpan data products, maka data akan diambil dari postgres, yang memakan waktu 13 ms dalam sekali hit.
### Ambil data dari redis
![alt text](image-1.png)
setelah menggunakan redis, data akan disimpan di redis, sehingga data yang dibutuhkan bisa langsung diambil di redis dan hanya memakan waktu 3 ms. ini jauh sangat meringankan server.