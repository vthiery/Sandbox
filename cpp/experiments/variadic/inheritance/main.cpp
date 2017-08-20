#include <iostream>

class A
{
public:
    inline void helloA() const noexcept { std::cout << "A" << std::endl; }
};

class B
{
public:
    inline void helloB() const noexcept { std::cout << "B" << std::endl; }
};

template <class... T>
class C : public T... {};

int main()
{
    std::cout << "Welcome to your beloved sandbox!" << std::endl;

    C<A, B> c;
    c.helloA();
    c.helloB();

    return 0;
}

