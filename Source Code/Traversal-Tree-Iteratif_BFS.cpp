#include <iostream>
#include <iomanip>
#include <queue>
using namespace std;

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
    if(node == NULL){
        return;
    }

    queue<Pohon*> ListNode; //buat queue dengan nama ListNode
    ListNode.push(node); //menambahkan node root kedalam queue ListNode

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

    return 0;
}