# Data username dan password
username_benar = "user123"
password_benar = "pass123"

# Meminta input dari pengguna
username = input("Masukkan username: ")
password = input("Masukkan password: ")

# Memeriksa kebenaran username dan password
if username == username_benar and password == password_benar:
    print("Login berhasil!")
    # Menampilkan pilihan lanjutan
    print("Pilih opsi:")
    print("1. Opsi 1")
    print("2. Opsi 2")
    print("3. Keluar")
    pilihan = input("Masukkan pilihan Anda: ")
    # Lakukan tindakan berdasarkan pilihan
    if pilihan == "1":
        print("Anda memilih Opsi 1")
    elif pilihan == "2":
        print("Anda memilih Opsi 2")
    elif pilihan == "3":
        print("Keluar dari program.")
    else:
        print("Pilihan tidak valid.")
else:
    print("Username atau password salah. Program berhenti.")
	class Library:
    def __init__(self):
        self.users = []
        self.books = []
        self.borrow_history = []

    def display_user_info(self):
        print("Informasi Pengguna:")
        for user in self.users:
            print(user)
    
    def display_books(self):
        print("Daftar Buku:")
        for book in self.books:
            print(book)

    def add_book(self, book_name):
        self.books.append(book_name)
        print(f"Buku '{book_name}' telah ditambahkan.")

    def borrow_book(self, book_name, user):
        if book_name in self.books:
            self.borrow_history.append((user, book_name))
            print(f"{user} telah meminjam buku '{book_name}'.")
        else:
            print("Buku tidak tersedia.")

    def display_borrow_history(self):
        print("Histori Peminjaman Buku:")
        for user, book in self.borrow_history:
            print(f"{user} meminjam '{book}'")


def main():
    library = Library()

    while True:
        print("\nMenu:")
        print("1. Lihat Informasi Pengguna")
        print("2. Lihat Daftar Buku")
        print("3. Tambah Buku")
        print("4. Tambah Peminjaman Buku")
        print("5. Histori Peminjaman Buku")
        print("6. Keluar dari Program")

        choice = input("Pilih menu (1-6): ")

        if choice == '1':
            library.display_user_info()
        elif choice == '2':
            library.display_books()
        elif choice == '3':
            book_name = input("Masukkan nama buku yang akan ditambahkan: ")
            library.add_book(book_name)
        elif choice == '4':
            user = input("Masukkan nama pengguna: ")
            book_name = input("Masukkan nama buku yang akan dipinjam: ")
            library.borrow_book(book_name, user)
        elif choice == '5':
            library.display_borrow_history()
        elif choice == '6':
            print("Keluar dari program.")
            break
        else:
            print("Pilihan tidak valid. Silakan coba lagi.")

if __name__ == "__main__":
    main()
	<!DOCTYPE html>
	<html lang="id">
	<head>
		<meta charset="UTF-8">
		<meta name="daiva" content="width=device-width, initial-scale=1.0">
		<title>Daftar Mahasiswa</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
				padding: 20px;
				background-color: #f4f4f4;
			}
			table {
				width: 100%;
				border-collapse: collapse;
			}
			th, td {
				padding: 10px;
				border: 1px solid #ccc;
				text-align: left;
			}
			th {
				background-color: #007BFF;
				color: white;
			}
		</style>
	</head>
	<body>
	
	<h1>Daftar Username dan Password Mahasiswa</h1>
	
	<table>
		<thead>
			<tr>
				<th>Nama Depan</th>
				<th>Username</th>
				<th>Password (NPM)</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>Alice</td>
				<td>alice</td>
				<td>123456789</td>
			</tr>
			<tr>
				<td>Bob</td>
				<td>bob</td>
				<td>987654321</td>
			</tr>
			<tr>
				<td>Charlie</td>
				<td>charlie</td>
				<td>123123123</td>
			</tr>
			<!-- Tambahkan data mahasiswa lain di sini -->
		</tbody>
	</ta
	# Daftar awal buku
daftar_buku = {
    'Pemrograman': {
        'nama_buku': 'Pemrograman',
        'jumlah': 10
    },
    'Film': {
        'nama_buku': 'Film',
        'jumlah': 5
    },
    'Printing': {
        'nama_buku': 'Printing',
        'jumlah': 20
    }
}

# Fungsi untuk menampilkan daftar buku
def tampilkan_daftar_buku(daftar):
    for kategori, info in daftar.items():
        print(f"{kategori}:")
        print(f"  Nama Buku: {info['nama_buku']}")
        print(f"  Jumlah: {info['jumlah']}")
        print()

# Menampilkan daftar buku
tampilkan_daftar_buku(daftar_buku)
#include <stdio.h>
#include <string.h>

// Mendefinisikan tipe data struct untuk Buku
struct Buku {
    char judul[100];
    char penulis[50];
    int tahunTerbit;
    float harga;
};

// Fungsi untuk mencetak informasi buku
void cetakBuku(struct Buku b) {
    printf("Judul: %s\n", b.judul);
    printf("Penulis: %s\n", b.penulis);
    printf("Tahun Terbit: %d\n", b.tahunTerbit);
    printf("Harga: %.2f\n", b.harga);
}

int main() {
    // Membuat sebuah objek Buku
    struct Buku buku1;

    // Mengisi data buku
    strcpy(buku1.judul, "Belajar C Pemrograman");
    strcpy(buku1.penulis, "John Doe");
    buku1.tahunTerbit = 2022;
    buku1.harga = 150000.00;

    // Mencetak informasi buku
    cetakBuku(buku1);

    return 0;
}
# List untuk menyimpan sejarah peminjaman
history_peminjaman = []

# Fungsi untuk menambahkan peminjaman
def tambah_peminjaman(nama_buku, peminjam, tanggal_pinjam):
    peminjaman = {
        'nama_buku': nama_buku,
        'peminjam': peminjam,
        'tanggal_pinjam': tanggal_pinjam
    }
    history_peminjaman.append(peminjaman)

# Fungsi untuk menampilkan sejarah peminjaman
def tampilkan_history():
    if not history_peminjaman:
        print("Tidak ada sejarah peminjaman.")
        return
    
    print("Sejarah Peminjaman:")
    for i, peminjaman in enumerate(history_peminjaman, start=1):
        print(f"{i}. Buku: {peminjaman['nama_buku']}, Peminjam: {peminjaman['peminjam']}, Tanggal Pinjam: {peminjaman['tanggal_pinjam']}")

# Contoh penggunaan
tambah_peminjaman("Belajar Python", "Ali", "2024-10-01")
tambah_peminjaman("Data Structures", "Budi", "2024-10-02")

tampilkan_history()
# Daftar pengguna dan buku
pengguna = {
    'nama': 'User1',
    'email': 'user1@example.com'
}

daftar_buku = []
peminjaman_buku = []

def LihatInformasiPenggunaProgram():
    print("Informasi Pengguna:")
    print(f"Nama: {pengguna['nama']}")
    print(f"Email: {pengguna['email']}\n")

def LihatDaftarBuku():
    if not daftar_buku:
        print("Daftar buku kosong.\n")
        return
    print("Daftar Buku:")
    for buku in daftar_buku:
        print(f"- {buku}")
    print()

def TambahDaftarBuku(buku):
    daftar_buku.append(buku)
    print(f"Buku '{buku}' telah ditambahkan ke daftar.\n")

def TambahPeminjamanBuku(buku):
    if buku in daftar_buku:
        peminjaman_buku.append(buku)
        print(f"Buku '{buku}' telah dipinjam.\n")
    else:
	class Library:
    def __init__(self, total_buku):
        self.total_buku = total_buku
        self.buku_pinjam = 0

    def pinjam_buku(self, jumlah):
        if jumlah < 0:
            print("Jumlah pinjaman tidak dapat negatif.")
            return
        
        if self.buku_pinjam + jumlah > self.total_buku:
            print("Tidak dapat meminjam buku. Jumlah yang diminta melebihi stok yang tersedia.")
            return
        
        self.buku_pinjam += jumlah
        print(f"Buku berhasil dipinjam: {jumlah}. Total buku yang dipinjam: {self.buku_pinjam}.")

    def kembalikan_buku(self, jumlah):
        if jumlah < 0:
            print("Jumlah pengembalian tidak dapat negatif.")
            return
        
        if jumlah > self.buku_pinjam:
            print("Jumlah pengembalian melebihi jumlah yang dipinjam.")
            return
        
        self.buku_pinjam -= jumlah
        print(f"Buku berhasil dikembalik
def menu():
    print("Selamat datang di program kami!")
    print("Pilih salah satu menu:")
    print("1. Menu A")
    print("2. Menu B")
    print("3. Keluar")

    while True:
        choice = input("Masukkan pilihan Anda (1-3): ")
        
        if choice == '1':
            print("\nAnda memilih Menu A.")
            # Tambahkan logika untuk Menu A di sini
            print("Terima kasih telah menggunakan Menu A.\n")
        elif choice == '2':
            print("\nAnda memilih Menu B.")
            # Tambahkan logika untuk Menu B di sini
            print("Terima kasih telah menggunakan Menu B.\n")
        elif choice == '3':
            print("Terima kasih! Selamat tinggal!")
            break
        else:
            print("Pilihan tidak valid, silakan coba lagi.")

if __name__ == "__main__":
    menu()

	  
	