#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>
#include<unordered_map>
#include<list>
#include<sstream>

 
using namespace std;

class LRUCache { 
    list<string> dq; 
    unordered_map<string, list<string>::iterator> ma; 
    int csize;
  
public: 
    LRUCache(int); 
    void refer(string, string); 
    //void display();
    bool isPresent(string str, string s2); 
}; 

LRUCache::LRUCache(int n) 
{ 
    csize = n; 
} 
  
void LRUCache::refer(string x, string y) 
{ 
    if (ma.find(x) == ma.end()) { 
        if (dq.size() == csize) { 
            string last = dq.back(); 
            dq.pop_back(); 
            ma.erase(last); 
        } 
    } 
  
    else
        dq.erase(ma[x]); 

    dq.push_front(y); 
    ma[x] = dq.begin(); 
} 

bool LRUCache::isPresent(string str, string s2) 
{ 
    if (ma.find(str) != ma.end()) {
        auto it = ma[str];
        if(*it == s2) return true;
        else return false;
    } else {
        return false;
    }
} 
  
// void LRUCache::display() 
// { 
//     for (auto it = dq.begin(); it != dq.end(); 
//          it++) 
//         cout << (*it) << " "; 
//     cout << endl; 
// } 


class NCache {
    public:
    int levels;
    vector<int> capacity;
    vector<int> readTimes;
    vector<int> writeTimes;
    unordered_map<int, LRUCache*> objectHash;
    NCache(int l, vector<int> &c, vector<int> &r, vector<int> &w, vector<LRUCache*> objs) {
        levels = l;

        for(int i : c) {
            capacity.push_back(i);
        }

        for(int i : r) {
            readTimes.push_back(i);
        }

        for(int i : w) {
            writeTimes.push_back(i);
        }

        for(int i = 0;i<objs.size();i++) {
            objectHash[i] = objs[i];
        }
    }
};
    
int main() 
{ 
    /*
    3
    2 5 10
    1 3 9
    5 10 15 */

    int n;
    cin>>n;
    vector<int> capacity;
    vector<int> readTimes;
    vector<int> writeTimes;
    for(int i=0;i<n;i++) {
        int t;
        cin>>t;
        capacity.push_back(t);
    }
    for(int i=0;i<n;i++) {
        int t;
        cin>>t;
        readTimes.push_back(t);
    }
        for(int i=0;i<n;i++) {
        int t;
        cin>>t;
        writeTimes.push_back(t);
    }

    vector<LRUCache*> objs;

    for(int i=0;i<n;i++) {
        LRUCache *l = new LRUCache(capacity[i]);
        objs.push_back(l);
    }
    
    NCache ca(n, capacity, readTimes, writeTimes, objs);
   string str;
   cin.ignore();
   while(getline(cin, str)) {

        if(str.empty()) break;

    istringstream ss(str);
    vector<string> w;
    do {
        string t;
        ss>>t;
        w.push_back(t);
    }while(ss);
    if(w[0] == "WRITE") {
            int writeTime = 0;
            int readTime = 0;
            int totalTime = 0;

            for(int i=0;i<n;i++) {
                auto o = ca.objectHash[i];
                if(o->isPresent(w[1], w[2])) {
                    readTime += ca.readTimes[i];
                } else {
                    writeTime += ca.writeTimes[i];
                    o->refer(w[1],w[2]);
                    //cout<<"here"<<endl;
                }
            }
            totalTime += readTime + writeTime;
            cout<<totalTime<<endl;
    }

    if(w[0] == "READ") {
            int writeTime = 0;
            int readTime = 0;
            int totalTime = 0;

            for(int i=0;i<n;i++) {
                auto o = ca.objectHash[i];
                if(o->isPresent(w[1],w[2])) {
                    readTime += ca.readTimes[i];
                    for(int j = i-1;j>=0;j--) {
                        writeTime += ca.writeTimes[j];
                        o->refer(w[1],w[2]);
                    }
                } else {
                    continue;
                }
            }
            totalTime += readTime + writeTime;
            cout<<totalTime<<endl;
    }

    if(w[0] == "STATS") {
        cout<<"GET STATS"<<endl;
    }

   }
    return 0; 
} 

