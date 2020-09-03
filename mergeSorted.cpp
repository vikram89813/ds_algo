#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>
#include<unordered_map>

using namespace std;

void mergeSorted(vector<int> &a1, vector<int> &a2) {
    int n = a1.size();
    int m = a2.size();

    int i =0;
    int j=0;

    while(i <n && j < m) {
        if(a1[i] > a2[j]) {
            swap(a1[i],a2[j]);
            i++;
        }
    }
}

int main() {
    vector<int> a1 = { 10, 27, 38, 43 ,82 }; 
    vector<int> a2 = { 3,9 };
    mergeSorted(a1,a2);
    return 0;
}