#include <iostream>
#include "inline_class.hpp"

int main()
{
    std::cout << "Welcome to your beloved sandbox!" << std::endl;

    InlineClass c{ 5 };
    std::cout << c.myVar() << std::endl;

    return 0;
}

