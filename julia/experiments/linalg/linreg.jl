using Base.Test
using Distributions

N = 12

# Prepare some data
x = [1.0:12.0;]
y = [5.5, 6.3, 7.6, 8.8, 10.9, 11.79, 13.48, 15.02, 17.77, 20.81, 22.0, 22.99]

# Regression y = a + bx
a, b = linreg(x, y)

println("a = ", a, "\nb = ", b)
