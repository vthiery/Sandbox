#include <iostream>
// Sum with no + ?!

unsigned int sum(const unsigned int a, const unsigned int b)
{
    const unsigned int exclOr{ a ^ b };
    const unsigned int carry{ (a & b) << 1 };
    return carry == 0 ? exclOr : sum(exclOr, carry);
}

int main()
{
    std::cout << "4 + 5 = " << sum(4, 5) << std::endl;
    std::cout << "10 + 56 = " << sum(10, 56) << std::endl;
}
