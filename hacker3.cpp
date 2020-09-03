#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>

using namespace std;

    string getRoundResult(char winning_suit, char suit1, int number1, char suit2, int number2) {
    // Write your code here
        if(winning_suit == suit1 &&winning_suit==suit2) {
            return (number1>number2?"Player 1 wins":(number1<number2?"Player 2 wins":"Draw"));
        }
        else if(winning_suit==suit1){
            return "Player 1 wins";
        }
        else if(winning_suit==suit2){
            return "Player 2 wins";
        }
        else {
            return (number1>number2?"Player 1 wins":(number1<number2?"Player 2 wins":"Draw"));
        }
    }


int main() {
    cout<<"hello"<<endl;
}