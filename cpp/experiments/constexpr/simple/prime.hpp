
constexpr bool isPrimeRecursive(const unsigned int number, const unsigned int ref)
{
    return (ref * ref > number) ? true : (number % ref == 0) ? false : isPrimeRecursive(number, ref + 1);
}

constexpr bool isPrime(const unsigned int number)
{
    return (number < 2) ? false : isPrimeRecursive(number, 2);
}

