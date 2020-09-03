#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>

using namespace std;

int func(int num, int po) {
    if(po == 0) return 1;
    if (po == 1) return num;
    int res = num * num;
    return func(res, po-1);
}

int main() {
    cout<<func(2,4)<<endl;
    cout<<"hello"<<endl;
}