Sampe 1:

**Input**: { [ ( ) ] }\
**Output**: YES\
**Penjelasan**: Setiap _bracket_ seimbang, antara _bracket_ buka dan _bracket_ tutup. \
_opening_ : { }\
_opening_ : [ ]\
_opening_ : ( )

Sampel 2: 

**Input**: { [ ( ] ) }

**Output**: NO

**Penjelasan**: String  { [ ( ] ) } tidak seimbang untuk karakter yang diapit oleh { dan } yaitu [ ( ] ).  

Sampel 3:\
**Input**: { ( ( [ ] ) [ ] ) [ ] }\
**Output**: YES
**Penjelasan**: Setiap _bracket_ seimbang, antara _bracket_ buka dan _bracket_ tutup, meskipun struktur _bracket_ tidak beraturan.

  

**Aturan:** 1. Tanda _bracket_ / karakter yang diperbolehkan sebagai berikut: **( , ) , { , } , atau [ , ]**.

2. _Bracket_ bisa dipisahkan **dengan** atau **tanpa** **whitespace**.

3. Periksa tanda kurung yang memiliki kecocokan antara _bracket_ buka dan _bracket_ tutup dengan mengembalikan nilai string **YES** atau **NO**.  

  

**Soal:** 1. Buat fungsi untuk menemukan Balanced Bracket dengan kompleksitas paling rendah!

2. Berapa **ukuran** **kompleksitas** kodinganmu? Jelaskan **detail** **kompleksitas** jawaban No.2, cantumkan di README Repo!