#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>
#include<unordered_map>

using namespace std;

string remove(string str) {
    for(int i=0;i<str.size();i++) {
        if(i == 0 || str[i]!= str[i-1]) continue;
        int k = 0;
        while(str[i+k] == str[i] && i+k < str.size()) {
            k++;
        }
        i--;
        k++;
        str.erase(i, k);
        i--;
    }
    return str;
}

int main() {
    string str1 = "geeksforgeeg"; 
    cout << remove(str1) << endl; 
  
    string  str2 = "azxxxzy"; 
    cout << remove(str2) << endl; 
  
    string str3 = "caaabbbaac"; 
    cout << remove(str3) << endl; 
  
    string str4 = "gghhg"; 
    cout << remove(str4) << endl; 
  
    string str5 = "aaaacddddcappp"; 
    cout << remove(str5) << endl; 
  
    string str6 = "aaaaaaaaaa"; 
    cout << remove(str6) << endl; 
  
    string str7 = "qpaaaaadaaaaadprq"; 
    cout << remove(str7) << endl; 
  
    string str8 = "acaaabbbacdddd"; 
    cout << remove(str8) << endl; 
  
    string str9 = "acbbcddc"; 
    cout << remove(str9) << endl; 
  
    return 0; 
}