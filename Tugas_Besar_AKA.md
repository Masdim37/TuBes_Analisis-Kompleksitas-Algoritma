# <h1 align="center">TUGAS BESAR ANALISIS KOMPLEKSITAS ALGORITMA (AKA)</h1>
# <h2 align="center">Analisis perbandingan Algoritma Breadth-First Search (BFS) dan Depth-First Search (DFS) dalam penelusuran Graf</h2>

### Anggota Kelompok 
Muhammad Dhimas Hafizh Fathurrahman - 2311102151 <br/>
Irfan Thoriq Habibi - 2311102131 <br/>
Kelas S1IF-11-04 <br/>
Telkom University Purwokerto <br/>

### Study Case
Traversal (penelusuran) adalah proses mengunjungi setiap simpul (node) dalam suatu pohon atau graf secara sistematis untuk mengetahui struktur dari graf tersebut. Terdapat 2 algoritma yang digunakan dalam traversal graf, yaitu BFS dan DFS. 

Algoritma Breadth-First Search (BFS) dan Algoritma Depth-First Search (DFS) adalah dua metode yang umum digunakan dalam penelusuran graf, masing-masing menggunakan pendekatan yang berbeda. Pada dasarnya, BFS menggunakan struktur data antrian (queue) untuk menjelajahi simpul secara level (Iteratif), dimulai dari simpul awal dan menambahkan semua tetangga yang belum dikunjungi ke queue. Di sisi lain, DFS mengunjungi simpul-simpul pohon atau graf secara berulang-ulang (Rekursif) mulai dari simpul awal yang dipilih, dan melanjutkan eksplorasi ke tetangga atau node terdekat yang belum dikunjungi.

Pada study case ini, dilakukan analisis perbandingan waktu eksekusi algoritma BFS dan DFS untuk melakukan traversal tree terhadap model graf atau pohon yang telah ditentukan.

### Algoritma BFS <br/>
Algoritma Breadth-First Search (BFS) atau yang disebut juga sebagai algoritma pencarian melebar adalah algoritma pencarian dan penjelajahan yang bekerja dengan mengunjungi simpul-simpul pohon atau graf secara bertingkat dari simpul awal. BFS memulai dari simpul awal dan mengunjungi semua tetangganya terlebih dahulu sebelum melanjutkan ke tingkat berikutnya.

Algoritma BFS dijalankan secara iteratif dan menggunakan antrian (queue) untuk memastikan urutan pengunjungannya, sehingga memungkinkan setiap tingkat dari pohon atau graf untuk dijelajahi sepenuhnya sebelum beralih ke tingkat berikutnya. Proses ini berulang sampai semua simpul yang dapat diakses telah dijelajahi atau tidak ada lagi simpul yang dapat diakses dari antrian.

Berikut merupakan cara kerja algoritma BFS :
1. Pilih simpul awal dalam tree (simpul A). 
2. Kunjungi semua simpul yang bertetanggaan dengan simpul A, biasanya di tingkat pertama. 
3. Kunjungi semua simpul yang bertetangga dengan simpul-simpul pada tingkat sebelumnya di tingkat berikutnya.
4. Ulangi proses hingga semua simpul dalam tree terjelajahi.

Berikut Merupakan Program Traversal Tree Menggunakan Algoritma BFS :
```C++
#include <iostream>
#include <iomanip>
#include <queue>
#include <chrono> //library untuk mencatat waktu eksekusi program
using namespace std;

//Mulai timer
auto start = chrono::high_resolution_clock::now();

struct Pohon{
    char data;
    Pohon *left, *right, *parent;
};
Pohon *root, *baru;

void init(){
    root = NULL;
}

bool isEmpty(){
    return root == NULL;
}

void buatNode(char data){
    if (isEmpty())
    {
        root = new Pohon();
        root->data = data;
        root->left = NULL;
        root->right = NULL;
        root->parent = NULL;
    }
    else
    {
        cout << "\n Tree sudah ada!" << endl;
    }
}

Pohon *insertLeft(char data, Pohon *node)
{
    if (isEmpty())
    {
        cout << "\n Buat tree terlebih dahulu!" << endl;
        return NULL;
    }
    else
    {
        if (node->left != NULL)
        {
            cout << "\n Node " << node->data << " sudah ada child kiri !" << endl;
                return NULL;
        }
        else
        {
            Pohon *baru = new Pohon();
            baru->data = data;
            baru->left = NULL;
            baru->right = NULL;
            baru->parent = node;
            node->left = baru;
                return baru;
        }
    }
}

Pohon *insertRight(char data, Pohon *node)
{
    if (isEmpty())
    {
        cout << "\n Buat tree terlebih dahulu!" << endl;
        return NULL;
    }
    else
    {
        if (node->right != NULL)
        {
            cout << "\n Node " << node->data << " sudah ada child kanan !" << endl;
                return NULL;
        }
        else
        {
            Pohon *baru = new Pohon();
            baru->data = data;
            baru->left = NULL;
            baru->right = NULL;
            baru->parent = node;
            node->right = baru;
                return baru;
        }
    }
}

void BFS(Pohon *node){
    if(node == NULL){ //jika node kosong, maka
        return; //keluar dari fungsi BFS
    }

    queue<Pohon*> ListNode; //buat queue dengan nama ListNode
    ListNode.push(node); //menambahkan node awal kedalam queue ListNode

    while(!ListNode.empty()){ //ketika queue tidak kosong, maka lakukan perulangan
        Pohon* current = ListNode.front(); //buat node current yang berisi elemen terdepan dalam queue
        ListNode.pop(); //menghapus elemen terdepan dalam queue

        cout << current->data << " "; //mencetak data node current

        if(current->left != nullptr){ //jika node kiri dari current tidak kosong, maka lakukan
            ListNode.push(current->left); //tambahkan node kiri dari current ke queue
        }

        if(current->right != nullptr){ //jika node kanan dari current tidak kosong, maka lakukan
            ListNode.push(current->right); //tambahkan node kanan dari current ke queue
        }
    }
}

int main(){
    buatNode('A'); //buat node A sebagai root
    Pohon *nodeB, *nodeC, *nodeD, *nodeE, *nodeF, *nodeG, *nodeH, *nodeI, *nodeJ, *nodeK, *nodeL, *nodeM;
    nodeB = insertLeft('B', root), //tambah node B sebagai child kiri node A
    nodeC =  insertRight('C', root), //tambah node C sebagai child kanan node A
    nodeD = insertLeft('D', nodeB), //tambah node D sebagai child kiri node B
    nodeE = insertRight('E', nodeB), //tambah node e sebagai child kanan node B
    nodeF = insertRight('F', nodeC), //tambah node F sebagai child kanan node C
    nodeG = insertLeft('G', nodeD), //tambah node G sebagai child kiri node E
    nodeH = insertRight('H', nodeD), //tambah node H sebagai child kanan node D
    nodeI = insertRight('I', nodeE), //tambah node I sebagai child kanan node E
    nodeJ = insertRight('J', nodeF); //tambah node J sebagai child kanan node F
    nodeK = insertLeft('K', nodeH); //tambah node K sebagai child kiri node H
    nodeL = insertRight('L', nodeH); //tambah node L sebagai child kanan node H
    nodeM = insertLeft('M', nodeJ); //tambah node M sebagai child kiri node J

    cout << "--- Traversal BFS ---" << endl;
    BFS(root);
    cout << endl;

    //Akhiri timer
    auto end = chrono::high_resolution_clock::now();

    //Hitung waktu yang dibutuhkan
    chrono::duration<double> duration = end - start;
    cout << fixed << setprecision(8);
    cout << "Waktu eksekusi = " << duration.count() << " detik" << endl;

    return 0;
}
```

### Algoritma DFS <br/>
Algoritma Depth-First Search (DFS) atau yang disebut juga sebagai algoritma pencarian mendalam adalah algoritma pencarian dan penjelajahan yang bekerja dengan mengunjungi simpul-simpul pohon atau graf secara berulang-ulang (rekursif) mulai dari simpul awal yang dipilih, dan melanjutkan eksplorasi ke tetangga terdekat yang belum dikunjungi.

Algoritma DFS dijalankan secara rekursif dan menekankan pada eksplorasi mendalam dengan cara mengikuti setiap cabang pohon hingga mencapai ujungnya sebelum kembali ke simpul sebelumnya dan mengeksplorasi cabang lainnya. Proses ini berulang sampai semua simpul yang dapat diakses telah dieksplorasi atau tidak ada lagi simpul yang dapat diakses dari simpul yang telah dikunjungi.

Berikut merupakan cara kerja algoritma DFS :
1. Pilih simpul awal dalam tree (simpul A).
2. Kunjungi simpul A.
3. Kunjungi semua simpul yang bertetangga langsung dengan simpul A (misal simpul B).
4. Ulangi proses mulai dari simpul B.
5. Ketika mencapai simpul N sedemikian sehingga semua simpul yang bertetangga dengannya telah dikunjungi, pencarian dirunut-balik (backtrack) ke simpul terakhir yang dikunjungi sebelumnya dan kumjungi simpul lain yang bertetangga dengannya.
6. Traversal berakhir bila tidak ada lagi simpul yang belum dikunjungi yang dapat dicapai dari simpul yang telah dikunjungi.

Berikut Merupakan Program Traversal Tree Menggunakan Algoritma DFS :
```C++
#include <iostream>
#include <iomanip>
#include <queue>
#include <chrono> //library untuk mencatat waktu eksekusi program
using namespace std;

//Mulai timer
auto start = chrono::high_resolution_clock::now();

struct Pohon{
    char data;
    Pohon *left, *right, *parent;
};
Pohon *root, *baru;

void init(){
    root = NULL;
}

bool isEmpty(){
    return root == NULL;
}

void buatNode(char data){
    if (isEmpty())
    {
        root = new Pohon();
        root->data = data;
        root->left = NULL;
        root->right = NULL;
        root->parent = NULL;
    }
    else
    {
        cout << "\n Tree sudah ada!" << endl;
    }
}

Pohon *insertLeft(char data, Pohon *node)
{
    if (isEmpty())
    {
        cout << "\n Buat tree terlebih dahulu!" << endl;
        return NULL;
    }
    else
    {
        if (node->left != NULL)
        {
            cout << "\n Node " << node->data << " sudah ada child kiri !" << endl;
                return NULL;
        }
        else
        {
            Pohon *baru = new Pohon();
            baru->data = data;
            baru->left = NULL;
            baru->right = NULL;
            baru->parent = node;
            node->left = baru;
                return baru;
        }
    }
}

Pohon *insertRight(char data, Pohon *node)
{
    if (isEmpty())
    {
        cout << "\n Buat tree terlebih dahulu!" << endl;
        return NULL;
    }
    else
    {
        if (node->right != NULL)
        {
            cout << "\n Node " << node->data << " sudah ada child kanan !" << endl;
                return NULL;
        }
        else
        {
            Pohon *baru = new Pohon();
            baru->data = data;
            baru->left = NULL;
            baru->right = NULL;
            baru->parent = node;
            node->right = baru;
                return baru;
        }
    }
}

void DFS(Pohon *node){
    if(node == NULL){ //jika node kosong, maka
        return; //keluar dari fungsi DFS
    }

    cout << node->data << " "; //kunjungi node awal
    DFS(node->left); //panggil kembali fungsi DFS dengan parameter node kiri dari node awal

    DFS(node->right); //panggil kembali fungsi DFS dengan parameter node kanan dari node awal
}

int main(){
    buatNode('A'); //buat node A sebagai root
    Pohon *nodeB, *nodeC, *nodeD, *nodeE, *nodeF, *nodeG, *nodeH, *nodeI, *nodeJ, *nodeK, *nodeL, *nodeM;
    nodeB = insertLeft('B', root), //tambah node B sebagai child kiri node A
    nodeC =  insertRight('C', root), //tambah node C sebagai child kanan node A
    nodeD = insertLeft('D', nodeB), //tambah node D sebagai child kiri node B
    nodeE = insertRight('E', nodeB), //tambah node e sebagai child kanan node B
    nodeF = insertRight('F', nodeC), //tambah node F sebagai child kanan node C
    nodeG = insertLeft('G', nodeD), //tambah node G sebagai child kiri node E
    nodeH = insertRight('H', nodeD), //tambah node H sebagai child kanan node D
    nodeI = insertRight('I', nodeE), //tambah node I sebagai child kanan node E
    nodeJ = insertRight('J', nodeF); //tambah node J sebagai child kanan node F
    nodeK = insertLeft('K', nodeH); //tambah node K sebagai child kiri node H
    nodeL = insertRight('L', nodeH); //tambah node L sebagai child kanan node H
    nodeM = insertLeft('M', nodeJ); //tambah node M sebagai child kiri node J

    cout << "--- Traversal DFS ---" << endl;
    DFS(root);
    cout << endl;

    //Akhiri timer
    auto end = chrono::high_resolution_clock::now();

    //Hitung waktu yang dibutuhkan
    chrono::duration<double> duration = end - start;
    cout << fixed << setprecision(8);
    cout << "Waktu eksekusi = " << duration.count() << " detik" << endl;

    return 0;
}
```

### Output
Diketahui model graf yang digunakan dalam traversal (penelusuran) menggunakan algoritma BFS dan DFS pada kedua kode diatas adalah sebagai berikut :

![Gambar_Model_Graf](https://github.com/Masdim37/TuBes_Analisis-Kompleksitas-Algoritma/blob/main/Images/Gambar_Model_Graf.jpg)

Berikut Output algoritma BFS :

![Output_BFS](https://github.com/Masdim37/TuBes_Analisis-Kompleksitas-Algoritma/blob/main/Images/Output_BFS.png)

Berikut Output algoritma DFS :

![Output_BFS](https://github.com/Masdim37/TuBes_Analisis-Kompleksitas-Algoritma/blob/main/Images/Output_DFS.png)

### Analisis
Berikut merupakan diagram batang perbandingan waktu eksekusi algoritma BFS dan algoritma DFS dalam satuan mikro detik.

![Diagram_Batang_Perbandingan_BFS-DFS](https://github.com/Masdim37/TuBes_Analisis-Kompleksitas-Algoritma/blob/main/Images/Diagram_Batang_Perbandingan_BFS-DFS.jpg)

Berdasarkan diagram diatas, terlihat bahwa waktu eksekusi yang dibutuhkan algoritma BFS lebih lama daripada waktu eksekusi yang dibutuhkan algoritma DFS.

Hal ini karena algoritma BFS menggunakan antrian (queue) yang mengimplementasikan iterasi untuk menelusuri simpul-simpul pohon atau graf secara bertingkat. Setiap simpul pada setiap tingkat harus dikunjungi sebelum melanjutkan ke tingkat berikutnya. Proses ini memerlukan waktu yang lebih lama karena BFS tidak hanya menelusuri cabang-cabang mendalam, tetapi juga mengunjungi simpul-simpul secara berturut-turut dari tingkat terendah ke tingkat tertinggi, yang memerlukan lebih banyak langkah dan iterasi.

Di sisi lain, algoritma DFS menggunakan rekursi, yang memungkinkan eksplorasi mendalam dari simpul-simpul pohon atau graf tanpa harus mengikuti urutan tingkat per tingkat. DFS lebih cepat dalam mengeksplorasi model graf tertentu karena ia tidak terhambat oleh urutan tingkat ke tingkat.

Algoritma BFS maupun DFS memiliki kompleksitas waktu yang sama yaitu Θ(n), karena setiap node yang ditelusur hanya diproses sebanyak 1x. Sehingga untuk n buah node, total operasi yang dilakukan adalah sebanyak n kali. Hal ini juga berlaku dalam best case dan worst case, dikarenakan traversal atau penelusuran yang dilakukan mencakup semua node yang ada didalam graf.

### Kesimpulan
Algoritma BFS & DFS memiliki fungsi yang sama dalam hal penelusuran graf, namun memiliki cara kerja yang berbeda. BFS bekerja dengan mengunjungi setiap simpul secara bertingkat dengan mengimplementasikan metode iteratif, sedangkan DFS bekerja dengan mengunjungi simpul secara mendalam dengan mengimplementasikan metode rekursif. 

Dalam hal efisiensi, BFS lebih cocok digunakan untuk graf yang pendek dan tidak memiliki banyak cabang, karena BFS mengimplementasikan queue untuk menyimpan hasil penelusuran sehingga semakin panjang dan semakin banyak cabang, maka penggunaan memori akan semakin tinggi. DFS lebih cocok digunakan untuk graf yang panjang, karena DFS bekerja dengan menelusuri setiap simpul secara mendalam. Namun, semakin panjang graf yang ditelusur akan berdampak pada waktu eksekusi yang diperlukan.

### Referensi 
Inggiantowi, H. (2008). “Perbandingan Algoritma Penelusuran Depth First Search dan Breadth First Search pada Graf serta Aplikasinya”. Program Studi Teknik Informatika, Sekolah Teknik Elektro dan Informatika, Institut Teknologi Bandung.