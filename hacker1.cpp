#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>

using namespace std;

int getStrength(string password, int weight_a) {
    // Complete the function
        int str = 0;
        for(char ch : password) {
            str+=(ch-'a'+weight_a)%26;
        }
        return str;
    }

int main() {
    cout<<"hello"<<endl;
}