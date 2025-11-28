# 📝 Todo CLI App (Golang + Cobra)

Todo CLI App ini adalah aplikasi simple tapi elegan yang ngebantu lu nyimpen, baca, update, hapus, dan nyari task lewat command line doang. Semua data disimpen di file JSON, jadi ga perlu database ribet.

## 🚀 Fitur

Aplikasi ini punya beberapa fitur yang bisa diakses via command:

• **add** – nambahin task baru
• **list** – nampilin semua task
• **update** – update status task
• **delete** – hapus task berdasarkan ID
• **search** – cari task berdasarkan nama (partial match supported 😎)

Data task otomatis disimpen ke file:

```
data/todo-lists.json
```

---

# 📂 Struktur Folder (Biar ga nyasar)

```
mini_project2/
│── cmd/
│   └── command CLI
│── service/
│   └── logic fitur todo
│── dto/
│   └── struct request & response
│── model/
│   └── struct data todo
│── utils/
│   └── file handling JSON
│── data/
│   └── todo-lists.json (auto generated)
│── main.go
```

---

# 🧠 Cara Install & Jalanin

## 1. Install Dependencies

Pastikan udah install Go minimal versi 1.20 ke atas.

## 2. Clone Project

```
git clone <url-repo-lu>
cd mini_project2
```

## 3. Run App

```
go run .
```

Atau build:

```
go build -o todo
./todo
```

---

# 🎮 Cara Pakai Command

## ➕ Add Task

Tambah task baru:

```
todo add --task "Belajar Go" --priority high
```

Kalo berhasil:

```
Added todo: Belajar Go (priority: high)
```

---

## 📋 List Semua Task

```
todo list
```

Output bakal jadi tabel kece:

```
+----+----------------------+---------+---------+
| ID | TASK                 | PRIORITY| STATUS  |
+----+----------------------+---------+---------+
| 1  | Belajar Go           | high    | pending |
+----+----------------------+---------+---------+
```

---

## 🔄 Update Status

Status cuma boleh: `pending`, `progress`, `done`

```
todo update --id 1 --status done
```

---

## ❌ Delete Task

```
todo delete --id 1
```

---

## 🔍 Search Task

Search by keyword, tidak harus full name (case insensitive):

```
todo search --task belajar
```

Output:

```
+----+----------------------+---------+---------+
| ID | TASK                 | PRIORITY| STATUS  |
+----+----------------------+---------+---------+
| 1  | Belajar Go           | high    | done    |
+----+----------------------+---------+---------+
```

---

# 🧩 Arsitektur Singkat

## Service Layer

`TodoService` ngurus semua logic:

* Create
* List
* Update
* Delete
* Search

## Utils

Ngurus file JSON:

* Auto bikin file kalo belum ada
* Read
* Write

## DTO

Nampung request & response biar rapi.

## Model

Struct `Todo` + `Base` termasuk ID, waktu created, waktu updated.

---

# ✨ Fitur Unik

• ID auto-increment meskipun file ga kosong
• Search dengan partial match
• Status warna-warni di terminal
• JSON file selalu valid karena auto-create

---

# 🧑‍💻 Contoh File JSON

Setelah isi task, isi file mirip kayak gini:

```json
[
 {
  "Id": 1,
  "CreatedAt": "2025-11-28T10:00:00Z",
  "UpdatedAt": "",
  "Task_name": "Belajar Go",
  "Status": "pending",
  "Priority": "high"
 }
]
