# Brief Proyek: EduLearn API

## Latar Belakang
EduLearn, adalah platform pendidikan berbasis teknologi yang menyediakan layanan pembelajaran online untuk siswa dan pengajar. Kami ingin mengembangkan sebuah RESTful API untuk mengelola kursus, siswa, dan instruktur secara efisien.

## Tujuan Proyek
Membuat API backend untuk platform EduLearn yang memungkinkan siswa untuk mendaftar kursus, instruktur untuk mengelola materi pembelajaran, dan admin untuk mengelola pengguna serta pembayaran.

## Fitur Utama
1. **Manajemen Pengguna**
2. **Manajemen Kursus**
3. **Manajemen Materi Pembelajaran**
4. **Sistem Pembayaran**
5. **Sertifikasi**

## Teknologi yang Digunakan
- **Bahasa Pemrograman**: Golang
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Autentikasi**: JWT
- **Gateway Pembayaran**: Midtrans/Stripe
- **Dokumentasi API**: Swagger

## Struktur Database
### Tabel `users`
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) CHECK (role IN ('student', 'instructor', 'admin')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Tabel `courses`
```sql
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    instructor_id INT REFERENCES users(id),
    price DECIMAL(10,2) DEFAULT 0,
    category VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Tabel `enrollments`
```sql
CREATE TABLE enrollments (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users(id),
    course_id INT REFERENCES courses(id),
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Tabel `materials`
```sql
CREATE TABLE materials (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES courses(id),
    title VARCHAR(255) NOT NULL,
    content TEXT,
    file_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Tabel `payments`
```sql
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users(id),
    course_id INT REFERENCES courses(id),
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('pending', 'completed', 'failed')) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Rute API
### 1️⃣ Authentication
| Method | Endpoint       | Deskripsi               | Role Akses |
|--------|----------------|-------------------------|-------------|
| **POST**   | `/auth/register` | Registrasi pengguna baru | Public      |
| **POST**   | `/auth/login`    | Login ke sistem         | Public      |

### 2️⃣ Manajemen Pengguna
| Method | Endpoint          | Deskripsi             | Role Akses |
|--------|-------------------|-----------------------|-------------|
| **GET**    | `/users`          | List semua pengguna  | Admin       |
| **GET**    | `/users/:id`      | Detail pengguna      | Admin, User |
| **PUT**    | `/users/me`       | Update profil        | User        |
| **DELETE** | `/users/:id`      | Hapus pengguna       | Admin       |

### 3️⃣ Manajemen Kursus
| Method | Endpoint         | Deskripsi               | Role Akses   |
|--------|------------------|-------------------------|---------------|
| **GET**    | `/courses`       | List semua kursus       | Public        |
| **POST**   | `/courses`       | Tambah kursus baru     | Instructor    |
| **GET**    | `/courses/:id`   | Detail kursus          | Public        |
| **PUT**    | `/courses/:id`   | Update kursus          | Instructor    |
| **DELETE** | `/courses/:id`   | Hapus kursus           | Instructor    |

### 4️⃣ Pendaftaran Kursus
| Method | Endpoint               | Deskripsi                   | Role Akses |
|--------|------------------------|-----------------------------|-------------|
| **POST**   | `/courses/:id/enroll`   | Daftar ke kursus              | Student     |
| **GET**    | `/users/:id/courses`    | List kursus yang diikuti      | Student     |

### 5️⃣ Materi Pembelajaran
| Method | Endpoint                 | Deskripsi                  | Role Akses   |
|--------|--------------------------|----------------------------|---------------|
| **GET**    | `/courses/:id/materials` | List semua materi kursus    | Public        |
| **POST**   | `/courses/:id/materials` | Tambah materi baru         | Instructor    |
| **GET**    | `/materials/:id`         | Detail materi              | Public        |
| **PUT**    | `/materials/:id`         | Update materi              | Instructor    |
| **DELETE** | `/materials/:id`         | Hapus materi               | Instructor    |

### 6️⃣ Pembayaran
| Method | Endpoint              | Deskripsi                   | Role Akses |
|--------|-----------------------|-----------------------------|-------------|
| **POST**   | `/payments`            | Proses pembayaran          | Student     |
| **GET**    | `/payments/:id`        | Detail pembayaran          | Student     |
| **GET**    | `/users/:id/payments`  | Riwayat pembayaran pengguna | Student     |

## Dokumentasi API

### Cara Menjalankan Project
```bash
# Clone repository
$ git clone https://github.com/adamfarizi/edu-learn.git
$ cd backend-api

# Install dependencies
$ go mod tidy

# Jalankan server
$ go run main.go
```

### Setup Database
- Buat database PostgreSQL baru: `edulearn`
- Import `schema.sql` untuk membuat tabel.
- Import `data.sql` untuk data dummy.

Setting environment database:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=edulearn
JWT_SECRET=yourjwtsecret
```

### Contoh Request dan Response

**Register User**
```
POST /auth/register
Content-Type: application/json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "student"
}
```

Response:
```json
{
  "message": "User registered successfully",
  "user_id": 1
}
```

**Login User**
```
POST /auth/login
Content-Type: application/json
{
  "email": "john@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "token": "<jwt_token>"
}
```

**Get All Courses**
```
GET /courses
Authorization: Bearer <jwt_token>
```

Response:
```json
[
  {
    "id": 1,
    "title": "Belajar Golang",
    "description": "Dasar hingga lanjutan",
    "price": 199.99,
    "category": "Programming"
  },
  {
    "id": 2,
    "title": "Design UI/UX",
    "description": "Teori dan Praktik",
    "price": 99.99,
    "category": "Design"
  }
]
```

## Deliverables
1. **Kode sumber API** dengan dokumentasi lengkap.
2. **Database schema** dan data dummy untuk pengujian.
3. **Dokumentasi API** menggunakan Swagger.
4. **Unit testing** minimal 80% coverage.

## Timeline Proyek
- **Minggu 1-2**: Perancangan database dan arsitektur API.
- **Minggu 3-5**: Implementasi fitur utama.
- **Minggu 6**: Pengujian dan debugging.
- **Minggu 7**: Deployment dan dokumentasi akhir.

## Catatan Tambahan
API ini akan digunakan oleh aplikasi web dan mobile EduLearn, sehingga harus dibuat scalable dan aman terhadap serangan siber seperti SQL Injection dan XSS.

---
Jika ada revisi atau tambahan fitur, silakan diskusikan dengan tim pengembang.

