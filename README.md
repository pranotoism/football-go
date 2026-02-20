# Football Team Management API

REST API backend untuk pengelolaan tim sepakbola amatir, dibangun menggunakan **Go**, **Gin Framework**, **GORM**, dan **PostgreSQL**.

## Tech Stack

- Go 1.21+
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL (Database)
- JWT (Authentication)
- bcrypt (Password hashing)

## Fitur

1. **Manajemen Tim** - CRUD data tim sepakbola (nama, logo, tahun berdiri, alamat, kota)
2. **Manajemen Pemain** - CRUD data pemain dengan validasi nomor punggung unik per tim
3. **Jadwal Pertandingan** - CRUD jadwal pertandingan antar tim
4. **Hasil Pertandingan** - Pelaporan hasil pertandingan (skor, pencetak gol, waktu gol)
5. **Laporan Pertandingan** - Report lengkap: skor, status, top scorer, akumulasi kemenangan
6. **Autentikasi** - JWT-based authentication
7. **Soft Delete** - Semua penghapusan data menggunakan mekanisme soft delete

## Setup & Installation

### Prerequisites

- Go 1.21 atau lebih baru
- PostgreSQL

### 1. Clone repository

```bash
git clone https://github.com/pranotoism/football-go.git
cd football-go
```

### 2. Setup database

Buat database PostgreSQL:

```sql
CREATE DATABASE football_go;
```

### 3. Konfigurasi environment

```bash
cp .env.example .env
```

Edit `.env` sesuai konfigurasi database Anda:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=football_go
JWT_SECRET=your-secret-key-here
APP_PORT=8080
```

### 4. Jalankan aplikasi

```bash
go mod tidy
go run main.go
```

Server akan berjalan di `http://localhost:8080`. Tabel database akan dibuat otomatis saat pertama kali dijalankan.

### 5. Testing dengan Postman

File Postman Collection sudah disediakan di root project: `Football_API.postman_collection.json`

1. Buka Postman → klik **Import** → pilih file `Football_API.postman_collection.json`
2. Jalankan request **Register** lalu **Login** terlebih dahulu
3. Token JWT akan otomatis tersimpan setelah Login berhasil
4. Semua request lainnya sudah dikonfigurasi untuk menggunakan token tersebut

## API Endpoints

### Auth (Public)

| Method | Endpoint                | Deskripsi                    |
| ------ | ----------------------- | ---------------------------- |
| POST   | `/api/v1/auth/register` | Register user baru           |
| POST   | `/api/v1/auth/login`    | Login, mendapatkan JWT token |

### Tim (Protected - Memerlukan JWT)

| Method | Endpoint            | Deskripsi                    |
| ------ | ------------------- | ---------------------------- |
| POST   | `/api/v1/teams`     | Tambah tim baru              |
| GET    | `/api/v1/teams`     | Daftar semua tim (paginated) |
| GET    | `/api/v1/teams/:id` | Detail tim beserta pemain    |
| PUT    | `/api/v1/teams/:id` | Update informasi tim         |
| DELETE | `/api/v1/teams/:id` | Hapus tim (soft delete)      |

### Pemain (Protected)

| Method | Endpoint                    | Deskripsi                  |
| ------ | --------------------------- | -------------------------- |
| POST   | `/api/v1/teams/:id/players` | Tambah pemain ke tim       |
| GET    | `/api/v1/teams/:id/players` | Daftar pemain dalam tim    |
| GET    | `/api/v1/players/:id`       | Detail pemain              |
| PUT    | `/api/v1/players/:id`       | Update data pemain         |
| DELETE | `/api/v1/players/:id`       | Hapus pemain (soft delete) |

### Pertandingan (Protected)

| Method | Endpoint                     | Deskripsi                        |
| ------ | ---------------------------- | -------------------------------- |
| POST   | `/api/v1/matches`            | Tambah jadwal pertandingan       |
| GET    | `/api/v1/matches`            | Daftar pertandingan (paginated)  |
| GET    | `/api/v1/matches/:id`        | Detail pertandingan              |
| PUT    | `/api/v1/matches/:id`        | Update jadwal pertandingan       |
| DELETE | `/api/v1/matches/:id`        | Hapus pertandingan (soft delete) |
| POST   | `/api/v1/matches/:id/result` | Laporkan hasil pertandingan      |

### Laporan (Protected)

| Method | Endpoint                     | Deskripsi                  |
| ------ | ---------------------------- | -------------------------- |
| GET    | `/api/v1/matches/:id/report` | Laporan satu pertandingan  |
| GET    | `/api/v1/reports/matches`    | Laporan semua pertandingan |

## Contoh Penggunaan API

### 1. Register

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name": "Admin", "email": "admin@xyz.com", "password": "password123"}'
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@xyz.com", "password": "password123"}'
```

Response:

```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOi..."
  }
}
```

### 3. Tambah Tim

```bash
curl -X POST http://localhost:8080/api/v1/teams \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "name": "Persib Bandung",
    "logo_url": "https://example.com/persib.png",
    "founded_year": 1933,
    "hq_address": "Jl. Sulanjana No. 17",
    "hq_city": "Bandung"
  }'
```

### 4. Tambah Pemain

```bash
curl -X POST http://localhost:8080/api/v1/teams/1/players \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "name": "Ciro Alves",
    "height_cm": 186,
    "weight_kg": 80,
    "position": "penyerang",
    "jersey_number": 9
  }'
```

### 5. Tambah Jadwal Pertandingan

```bash
curl -X POST http://localhost:8080/api/v1/matches \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "match_date": "2026-03-15",
    "match_time": "19:00:00",
    "home_team_id": 1,
    "away_team_id": 2
  }'
```

### 6. Laporkan Hasil Pertandingan

```bash
curl -X POST http://localhost:8080/api/v1/matches/1/result \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "home_score": 2,
    "away_score": 1,
    "goals": [
      {"player_id": 1, "team_id": 1, "minute": 23},
      {"player_id": 1, "team_id": 1, "minute": 67},
      {"player_id": 3, "team_id": 2, "minute": 45}
    ]
  }'
```

### 7. Lihat Laporan Pertandingan

```bash
curl http://localhost:8080/api/v1/matches/1/report \
  -H "Authorization: Bearer <token>"
```

Response:

```json
{
  "status": "success",
  "message": "Match report retrieved successfully",
  "data": {
    "match_id": 1,
    "match_date": "2026-03-15",
    "match_time": "19:00:00",
    "home_team": { "id": 1, "name": "Persib Bandung" },
    "away_team": { "id": 2, "name": "Persija Jakarta" },
    "home_score": 2,
    "away_score": 1,
    "status": "Home Win",
    "goals": [
      {
        "player_name": "Ciro Alves",
        "team_name": "Persib Bandung",
        "minute": 23
      },
      {
        "player_name": "Marko Simic",
        "team_name": "Persija Jakarta",
        "minute": 45
      },
      {
        "player_name": "Ciro Alves",
        "team_name": "Persib Bandung",
        "minute": 67
      }
    ],
    "top_scorer": { "player_name": "Ciro Alves", "goals": 2 },
    "cumulative_home_wins": 5,
    "cumulative_away_wins": 3
  }
}
```

## Pagination

Semua endpoint list mendukung pagination:

```
GET /api/v1/teams?page=1&per_page=10
```

## Posisi Pemain

Pilihan posisi pemain: `penyerang`, `gelandang`, `bertahan`, `penjaga_gawang`

## Asumsi

1. Nomor punggung pemain unik dalam satu tim (pemain aktif/belum dihapus)
2. Satu pertandingan hanya bisa dilaporkan hasilnya sekali
3. Jumlah gol yang dilaporkan harus sesuai dengan skor akhir
4. Saat tim dihapus (soft delete), semua pemain dalam tim juga ikut di-soft-delete
5. Semua endpoint selain auth memerlukan JWT token

## Struktur Proyek

```
football-go/
├── main.go              # Entry point
├── config/              # Konfigurasi (env vars)
├── database/            # Koneksi database & migrasi
├── model/               # GORM models
├── dto/                 # Data Transfer Objects (request/response)
├── repository/          # Data access layer
├── service/             # Business logic
├── handler/             # HTTP handlers
├── middleware/           # Auth & error middleware
├── router/              # Route definitions
└── util/                # Helpers (JWT, response)
```
