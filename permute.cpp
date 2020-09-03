#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>
#include<unordered_map>

using namespace std;

void permute(string str, int l, int r) {
    if(l == r) cout<<str<<endl;
    else {
        for(int i =l;i<=r;i++) {
            swap(str[i],str[l]);
            permute(str,l+1,r);
            swap(str[i],str[l]);
        }
    }
}

// Driver Code  
int main()  
{  
    string str = "ABC";  
    int n = str.size();  
    permute(str, 0, n-1);  
    return 0;  
}  