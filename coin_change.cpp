#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>

using namespace std;

int dp[4][5];

int count(int arr[], int m, int n) {
    if(n<0) return 0;
    if(n==0) return 1;
    if(m<=0 && n>=1) return 0;
    if(dp[m][n]) return dp[m][n];

    dp[m][n] = count(arr,m-1,n) + count(arr,m,n-arr[m-1]);
    return dp[m][n];
}

int main() {
    int i, j; 
    int arr[] = {1, 2, 3}; 
    int m = sizeof(arr)/sizeof(arr[0]); 
    cout<<count(arr, m, 4)<<endl;
    return 0;
}