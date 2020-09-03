/*
Further Improvements on the Program
1. Program is not meant for concurrency.
2. Will have to introduce threads for concurrency.
3. Will have to manage critical section to avoid data race, race condition and deadlock if using multithreading.
4. The data can be moved to databases to introduce ACID and handle concurrency.
5. Metadata can be stored in a NoSQL data store.
6. Analytic data can be stored on persistent stores like Redis.
7. uuid or a predefined key store can be used for user/merchant id instead of integer in interest of scaling.
8. more contraints like handeling negative credit limits can be imposed on user model. 
*/

#include<iostream>
#include<vector>
#include<string>
#include<stack>
#include<algorithm>
#include<queue>
#include<math.h>
#include<unordered_map>

using namespace std;

// class to store merchant information.
class merchants{
    int id;
    double discount;
    string name;

    public:

    merchants(int i, double d, string name) {
        id = i;
        discount = d;
        name = name;
    }
    void changeDiscount(double d) {
        discount = d;
    }

    double getDiscount() {
        return discount;
    }
};

// class to store users.
class users{
    int id;
    string name;
    string email_id;
    double credit_limit;
    double due_amount;
    public:

    users(int i, string name, string email_id, double c_limit) {
        id = i;
        name = name;
        email_id = email_id;
        credit_limit = c_limit;
        due_amount = 0;
    }

    void updateCreditLimit(double amount) {
        credit_limit -= amount;
    }

    void updateDueAmount(double amount) {
        due_amount += amount;
    }

    double getCredit() {
        return credit_limit;
    }

    double getDue() {
        return due_amount;
    }

    void payDue(double amount) {
        credit_limit += amount;
        due_amount -= amount;
    }
};

// Master class of our service.
class PaylaterService {
    // these can be moved to database or a persistent data store.

    // get merchant id from merchant name.
    unordered_map<string, int> merchant_id_map;
    // get merchant object from merchant id.
    unordered_map<int, merchants *> merchant_map;
    // get all discount of a particulat merchant, used for analytics.
    unordered_map<int, double> merchant_data;

    // get id based on user email
    unordered_map<string, int> user_id_map;
    // get user object from user id.
    unordered_map<int, users *> user_map;
    // get all dues for a user.
    unordered_map<int, double> user_data;

    int merchant_count;
    int user_count;

    public:

    PaylaterService() {
        merchant_count = 0;
        user_count = 0;
    }

    void registerMerchant(string name, double discount) {
        merchants *merchant = new merchants(merchant_count, discount, name);
        merchant_id_map[name] = merchant_count;
        merchant_map[merchant_count] = merchant;
        merchant_data[merchant_count] = 0;
        merchant_count++;
        cout<<"Merchant registered!"<<endl;
    }

    void registerUser(string name, string email_id, double c_limit) {
         users *user = new users(user_count, name, email_id, c_limit);
         user_id_map[email_id] = user_count;
         user_map[user_count] = user;
         user_data[user_count] = 0;
         user_count++;
         cout<<"User registered!"<<endl;
    }

    double getMerchantData(string name) {
        if(merchant_id_map.find(name) == merchant_id_map.end()) {
            cout<<"Merchant not found!"<<endl;
            return -1;
        }
        int merchant_id = merchant_id_map[name];
        return merchant_data[merchant_id];
    }

    double getUserDues(string email) {
        if(user_id_map.find(email) == user_id_map.end()) {
            cout<<"User not found!"<<endl;
            return -1;
        }
        int user_id = user_id_map[email];
        auto user = user_map[user_id];
        return user->getDue();
    }

    double getTotalDues() {
        double total = 0;
        for(auto &it :  user_map) {
            auto user = it.second;
            total += user->getDue();
        }
        return total;
    }

    void transact(string email, string merchant_name, double amount) {
        if(user_id_map.find(email) == user_id_map.end()) {
            cout<<"User not found!"<<endl;
            return;
        }
        if(merchant_id_map.find(merchant_name) == merchant_id_map.end()) {
            cout<<"Merchant not found!"<<endl;
            return;
        }

        //get user
        int user_id = user_id_map[email];
        auto user = user_map[user_id];

        // get merchant
        int merchant_id = merchant_id_map[merchant_name];
        auto merchant = merchant_map[merchant_id];

        if(user->getCredit() < amount) {
            cout<<"User credit less than transcation amount!"<<endl;
            return;
        }

        //get merchant discount.
        double discount = merchant->getDiscount();

        double discount_amount = (discount/100) * amount;
        double actual_amount = amount - discount_amount;

        merchant_data[merchant_id] += discount_amount;

        user->updateCreditLimit(actual_amount);
        user->updateDueAmount(actual_amount);

        user_data[user_id] += actual_amount;
        cout<<"Transcation successful!"<<endl;
    }

    void payDue(string email, double amount) {
        if(user_id_map.find(email) == user_id_map.end()) {
            cout<<"User not found!"<<endl;
            return;
        }

        //get user
        int user_id = user_id_map[email];
        auto user = user_map[user_id];
        user->payDue(amount);
        cout<<"Due paid!"<<endl;
    }

    double getCredit(string email) {
        if(user_id_map.find(email) == user_id_map.end()) {
            cout<<"User not found!"<<endl;
            return 0;
        }

        //get user
        int user_id = user_id_map[email];
        auto user = user_map[user_id];

        return user->getCredit();
    }
};

int main() {
    PaylaterService ps;

    //register merchants
    ps.registerMerchant("PayPal", 10);
    ps.registerMerchant("MasterCard", 5);

    // register users
    ps.registerUser("John", "john123@test.com", 2000.00);
    ps.registerUser("Jane", "jane567@test.com", 3000.00);

    // perform a transcation.
    ps.transact("john123@test.com", "PayPal", 500.00);

    cout<<"data after a transcation"<<endl;
    cout<<"Due amount for user John : ";
    cout<<ps.getUserDues("john123@test.com")<<endl;
    cout<<"Discount amount for PayPal uptil current date: ";
    cout<<ps.getMerchantData("PayPal")<<endl;
    cout<<"Total dues for all users: ";
    cout<<ps.getTotalDues()<<endl;

    cout<<"After paying due amount"<<endl;

    // pay dues
    ps.payDue("john123@test.com", 450.00);
    cout<<"Due amount for user John : ";
    cout<<ps.getUserDues("john123@test.com")<<endl;

    cout<<"Discount amount for PayPal uptil current date: ";
    cout<<ps.getMerchantData("PayPal")<<endl;
    cout<<"Total dues for all users: ";
    cout<<ps.getTotalDues()<<endl;

    cout<<"Transcation of amount greater than credit limit "<<endl;
    cout<<"get credit limit for John : ";
    cout<<ps.getCredit("john123@test.com")<<endl;
    ps.transact("john123@test.com", "PayPal", 4000.00);
}