## 📄 Dokumen Perencanaan Proyek Web

**Nama Proyek:** Aplikasi Profil Uji Coba redis dan docker
**Tanggal:** 05 September 2026
**Penulis:** Rendi Damara

### 1. Ringkasan Proyek

Profil adalah aplikasi web agar si pengguna dengan keamanan JWT bisa melakukan Login, Register, dan Read, keamanan dengan OTP.

**Tujuan Bisnis:**

- Menyediakan kemudahan untuk mengelola profil dengan keamanan sebaik mungkin

**Target Pengguna:**

- Individu yang ingin mencoba membuat profil

---

### 2. Daftar Fitur (User Stories)

| ID  | Fitur              |
| --- | ------------------ |
| F1  | Registrasi & Login |
| F2  | Melihat profil     |

### 3. Arsitektur Sistem

- **Frontend:** HTML, CSS, JavaScript dengan framework **React**.
- **Backend:** Golang, dengan framework http fiber
- **Database:** **Mysql** (relasional) untuk menyimpan data pengguna dan tugas, Redis untuk penyimpanan sementara.
- **Autentikasi:** JSON Web Token (JWT) disimpan di cookie httpOnly.

**Diagram alur sederhana:**

```
[ Browser (React) ] <--> [ REST API (Golang) ] <--> [ Mysql]
```

---

### 4. Desain Database (ERD Sederhana)

**Tabel `users`**

```
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'pelajar',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
```

---

### 5. Alur Proses Utama (Flow)

**Alur Login:**

1. User membuka halaman login.
2. Mengisi email dan password.
3. Frontend mengirim request `POST /api/login`.
4. Backend memvalidasi kredensial.
5. Jika valid, backend membuat token JWT dan mengirimkannya ke frontend.
6. Frontend menyimpan token (misal di memory atau cookie) dan mengarahkan ke dashboard.

Alur Register:

1. User membuka halaman register
2. Input nama, username, email, password
3. fronted mengirimkan lewat post POST/api/register
4. Backend menerima lewat POST/api/register memvalidasi dan meneruskan ke DB
5. Backend memberikan info ke frontend lewat POST/api/register

Alur Read:

1. User membuka halaman profil
2. Frontend mengirim req lewat GET/api/profil/:id
3. Backend menerima req lewat GET/api/profil/id
4. Backend mengirimkan data user lewat GET/api/profil/id

---

### 6. Rencana Halaman (Wireframe Deskripsi)

- **Halaman Login/Registrasi** – Form sederhana, link beralih antara login dan daftar.
- **Halaman Dashboard** – Menampilkan statistik (jumlah tugas belum/selesai), daftar tugas terbaru, dan tombol "Tambah Tugas".
- **Halaman Daftar Tugas** – Tabel atau kartu tugas dengan filter status/kategori, dan kolom pencarian.
- **Halaman Detail/Edit Tugas** – Form untuk mengubah detail tugas, tombol hapus.

---

### 7. Testing Plan

- **Unit Testing:** Menguji fungsi validasi input, pembuatan token, dan logika update status.
- **Integration Testing:** Menguji alur API (misal: login -> buat tugas -> tampil di daftar).
- **Manual Testing:** Menguji tampilan responsif di browser mobile dan desktop, mencoba skenario error (input kosong, password salah).

---

### 8. Milestone Pengerjaan

| Tahap    | Aktivitas                                                                    | Estimasi |
| -------- | ---------------------------------------------------------------------------- | -------- |
| Minggu 1 | Setup proyek (repo, environment), desain database, buat API auth.            | 3 hari   |
| Minggu 2 | Implementasi CRUD tugas di backend, buat frontend halaman login & dashboard. | 4 hari   |
| Minggu 3 | Integrasi frontend-backend, tambahkan fitur filter, styling.                 | 3 hari   |
| Minggu 4 | Testing, perbaikan bug, deployment.                                          | 2 hari   |
