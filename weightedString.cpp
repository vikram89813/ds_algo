#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>
#include<unordered_map>

using namespace std;

vector<long long> h;
unordered_map<char,int> h1;

void preprocess() {
    h.push_back(1);
    long long c =1;

    for(int i=1;i<26;i++) {
        long long t = (i+1) * c + c;
        c = t;
        h.push_back(c);
    }

    for(int i = 0;i<26;i++) {
        h1[i + 'A'] = i;
    }
}

string ans_ = "ABCDEFGHIGKLMNOPQRSTUVWXYZ";
long min_l = INT_MAX;



bool helper(vector<char> arr, int m, int n, string ans) {
    if (n == 0)  {
        if(ans.size() < ans_.size()) {
            ans_ = ans;
        }
        return true; 
    }
        
    if (n < 0) 
        return false; 

    if (m <=0 && n >= 1) 
        return false; 


    return helper( arr, m - 1, n, ans ) + helper( arr, m, n-h[arr[n] - 'A'], ans+arr[m-1] ); 
}

string func(int w) {
    preprocess();
    // cout<<h[65-'A']<<endl;
    int idx = 0;
    for(int i=0;i<26;i++) {
        if(h[i] > w) {
            idx = i;
            break;
        }
    }
    

    vector<char> c_arr;
    for(int i=0;i<idx;i++) {
        int g = i + 'A';
        c_arr.push_back(char(g));
    }

    bool s = helper(c_arr, c_arr.size(),w, "");
    sort(ans_.begin(), ans_.end());
    return ans_;
}

int main() {
    int w = 25;
    cout<<func(w)<<endl;
}