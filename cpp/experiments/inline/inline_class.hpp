

class InlineClass
{
public:
    explicit InlineClass(const int var);

    inline int myVar() const noexcept;

private:
    const int m_myVar;
};

InlineClass::
InlineClass(const int var)
    : m_myVar(var)
{}

inline int InlineClass::myVar() const noexcept
{
    return m_myVar;
}

