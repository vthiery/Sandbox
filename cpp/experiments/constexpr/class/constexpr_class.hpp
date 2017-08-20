

class ConstThingy
{
public:
    constexpr ConstThingy(const int someVar);

    constexpr int someVar() const noexcept;

    constexpr unsigned int someVarSquared() const noexcept;

private:
    const int m_someVar;

};

constexpr ConstThingy::ConstThingy(const int someVar)
    : m_someVar(someVar)
{}

constexpr int ConstThingy::someVar() const noexcept
{
    return m_someVar;
}

constexpr unsigned int ConstThingy::someVarSquared() const noexcept
{
    return static_cast<unsigned int>(m_someVar * m_someVar);
}

