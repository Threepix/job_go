#include <iostream>
#include <ctime>

using namespace std;

int main(){

    int a = 0, b = 3, c = 3;

    unsigned int start_time = clock();
    for (int i = 0; i <100000000; i++){
       a += 2*b + c - i; 
    }
    unsigned int end_time = clock();
    unsigned int search_time = end_time - start_time;
    cout << "time";
    cout << search_time << "ms\n";
    system("pause");
    return 0;
}
