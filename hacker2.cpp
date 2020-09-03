#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>

using namespace std;

vector<int> getMaxStreaks(vector<string> toss) {
    // Return an array of two integers containing the maximum streak of heads and tails respectively
        string prev="";
        int mxH=0,mxT=0,cH=0,cT=0;
        for(string curr: toss) {
            if("Heads" == curr){
                if(prev == curr){
                    cH++;
                }
                else{
                    cH=1;
                }
                mxH=max(mxH,cH);
            }
            else {
                if(prev == curr){
                    cT++;
                }
                else{
                    cT=1;
                }
                mxT= max(mxT,cT);
            }
            prev=curr;
        }
        vector<int> ans;
        ans.push_back(mxH);
        ans.push_back(mxT);
        return ans;
    }

int main() {
    cout<<"hello"<<endl;
}