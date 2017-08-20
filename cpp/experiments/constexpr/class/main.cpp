#include <iostream>
#include "constexpr_class.hpp"

int main()
{
    std::cout << "Welcome to your beloved sandbox!" << std::endl;

    ConstThingy c{ -2 };
    std::cout << c.someVar() << std::endl;
    std::cout << c.someVarSquared() << std::endl;

    return 0;
}

