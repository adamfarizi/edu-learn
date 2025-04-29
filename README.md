# 📘 EduLearn API

RESTful API backend untuk platform pembelajaran online **EduLearn** yang melayani siswa, instruktur, dan admin dalam satu sistem terpusat.

---

## 📚 Ringkasan Proyek

**EduLearn** adalah platform pendidikan berbasis teknologi yang menyediakan layanan pembelajaran online. Sistem ini akan mengelola:
- Registrasi & otentikasi pengguna
- Kursus & materi pembelajaran
- Pendaftaran kursus
- Pembayaran & sertifikasi

---

## 🛠️ Teknologi

- **Golang** + **Gin**
- **PostgreSQL**
- **GORM**
- **JWT** (autentikasi)
- **Swagger** (dokumentasi API)
- **Midtrans/Stripe** (opsional pembayaran)

---

## 🧠 Struktur Database

### `users`
```sql
id SERIAL PRIMARY KEY,
name VARCHAR(255),
email VARCHAR(255) UNIQUE,
password VARCHAR(255),
role VARCHAR(50) CHECK (role IN ('student', 'instructor', 'admin')),
created_at TIMESTAMP,
updated_at TIMESTAMP
```

### `courses`
```sql
id SERIAL PRIMARY KEY,
title VARCHAR(255),
description TEXT,
instructor_id INT REFERENCES users(id),
price DECIMAL(10,2),
category VARCHAR(100),
created_at TIMESTAMP,
updated_at TIMESTAMP
```

### `enrollments`
```sql
id SERIAL PRIMARY KEY,
student_id INT REFERENCES users(id),
course_id INT REFERENCES courses(id),
enrolled_at TIMESTAMP
```

### `materials`
```sql
id SERIAL PRIMARY KEY,
course_id INT REFERENCES courses(id),
title VARCHAR(255),
content TEXT,
file_url VARCHAR(255),
created_at TIMESTAMP
```

### `payments`
```sql
id SERIAL PRIMARY KEY,
student_id INT REFERENCES users(id),
course_id INT REFERENCES courses(id),
amount DECIMAL(10,2),
status VARCHAR(50) CHECK (status IN ('pending', 'completed', 'failed')),
payment_date TIMESTAMP
```

---

## 🚀 Cara Menjalankan

```bash
# 1. Clone project
git clone https://github.com/adamfarizi/edu-learn.git
cd edu-learn

# 2. Install dependency
go mod tidy

# 3. Jalankan server
go run main.go
```

### 🛢️ Setup Database

1. Buat database PostgreSQL: `edulearn`
2. Import `schema.sql`
3. Import `data.sql` (opsional dummy)

### 🔐 `.env` Config
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=edulearn
JWT_SECRET=yourjwtsecret
```

---

## 📌 Rute API & Role Akses

### 🔐 Authentication
| Method | Endpoint         | Deskripsi         | Akses     |
|--------|------------------|-------------------|-----------|
| POST   | `/auth/register` | Registrasi        | Public    |
| POST   | `/auth/login`    | Login             | Public    |

### 👤 Pengguna
| Method | Endpoint      | Deskripsi         | Akses              |
|--------|---------------|-------------------|--------------------|
| GET    | `/users`      | List user         | Admin              |
| GET    | `/users/:id`  | Detail user       | Admin, User        |
| PUT    | `/users/me`   | Update profil     | User               |
| DELETE | `/users/:id`  | Hapus user        | Admin              |

### 📚 Kursus
| Method | Endpoint         | Deskripsi        | Akses       |
|--------|------------------|------------------|-------------|
| GET    | `/courses`       | Semua kursus     | Public      |
| POST   | `/courses`       | Tambah kursus    | Instructor  |
| GET    | `/courses/:id`   | Detail kursus    | Public      |
| PUT    | `/courses/:id`   | Ubah kursus      | Instructor  |
| DELETE | `/courses/:id`   | Hapus kursus     | Instructor  |

### 📝 Enroll
| Method | Endpoint               | Deskripsi          | Akses    |
|--------|------------------------|--------------------|----------|
| POST   | `/courses/:id/enroll`  | Daftar kursus      | Student  |
| GET    | `/users/:id/courses`   | Kursus saya        | Student  |

### 📄 Materi
| Method | Endpoint                | Deskripsi          | Akses      |
|--------|-------------------------|--------------------|------------|
| GET    | `/courses/:id/materials`| Lihat semua materi | Public     |
| POST   | `/courses/:id/materials`| Tambah materi      | Instructor |
| GET    | `/materials/:id`        | Lihat materi       | Public     |
| PUT    | `/materials/:id`        | Ubah materi        | Instructor |
| DELETE | `/materials/:id`        | Hapus materi       | Instructor |

### 💳 Pembayaran
| Method | Endpoint             | Deskripsi             | Akses   |
|--------|----------------------|-----------------------|---------|
| POST   | `/payments`          | Proses bayar kursus   | Student |
| GET    | `/payments/:id`      | Detail pembayaran     | Student |
| GET    | `/users/:id/payments`| Riwayat pembayaran    | Student |

---

## 📥 Contoh Payload

### 🧑‍🎓 Register
```json
POST /auth/register
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "student"
}
```

### 🔐 Login
```json
POST /auth/login
{
  "email": "john@example.com",
  "password": "password123"
}
```

### ➕ Tambah Kursus
```json
POST /courses
{
  "title": "Belajar Golang",
  "description": "Kursus dari dasar hingga mahir",
  "price": 150000,
  "category": "Programming"
}
```

### 🖋️ Tambah Materi
```json
POST /courses/1/materials
{
  "title": "Intro Golang",
  "content": "Penjelasan dasar Golang",
  "file_url": "https://example.com/intro.pdf"
}
```

### 💰 Bayar Kursus
```json
POST /payments
{
  "course_id": 1,
  "amount": 150000
}
```
