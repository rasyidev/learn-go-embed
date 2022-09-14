## Go Lang Embed
- Fitur untuk mempermudah membaca isi file pada saat compile time secara otomatis dimasukkan ke dalam variable
- Tidak perlu load file secara manual
- File yang disupport adalah file yang berupa teks seperti txt, json, csv, dll.
- **Implementasi**: tambahkan `//go:embed <nama file>` di atas variable yang dituju.
  - Variabel yang dituju tersebut otomatis berisi konten file secara otomatis ketika kode go lang di-compile
  - Variable yang dituju tidak bisa disimpan di dalam function