using Base.Test

A = [1 2 3; 4 1 6; 7 8 1]

@test trace(A) == 3
@test det(A) == 104

# Invert
@time inv(A) 
# Eigenvalues/eigenvecs decomposition
@time eigvals(A)
@time eigvecs(A)
# Smart factorization
@time factorize(A)

# Linear solving Ax = b 
b = [1; 2; 3]
x = A\b
x

# See JuliaLang reference for more about operations and factorizations
