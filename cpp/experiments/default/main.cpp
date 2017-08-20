#include <iostream>
#include "implicit.hpp"

int main()
{
    std::cout << "Welcome to your beloved sandbox!" << std::endl;

    B b1, b2;
    // assign by lvalue reference
    b1 = b2;
    // copy by lvalue reference
    B b3{ b1 };
    // assign by rvalue reference
    b3 = B{};
    // copy by lvalue reference
    B b4{ B{} };

    return 0;
}

