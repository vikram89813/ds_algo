#include<iostream>

using namespace std;

int powerFunc(int number, int power) 
{ 
    int temp; 
    if( power == 0) 
        return 1; 
    temp = powerFunc(number, power/2); 
    if (power%2 == 0) 
        return temp*temp; 
    else
        return number*temp*temp; 
}

int func(int number, int power) {
    if(power <= 0) return number;
    int result = 0;
    if(power % 2) {
        result = number*number*number;
        return func(result, power-1);
    } else {
        int result = number*number;
        return func(result, power-2);
    }
}

int main() {
    cout<<powerFunc(2,3)<<endl;
    cout<<"hello"<<endl;
}