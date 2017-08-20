using Optim
using Base.Test

# Define the Rosenbrock function
f(x) = (1.0 - x[1]) ^ 2 + 100.0 * (x[2] - x[1] ^ 2) ^ 2
# Define its gradient
function g!(x::Vector, grad::Vector)
    grad[1] = - 2.0 * (1.0 - x[1]) - 400.0 * (x[2] - x[1] ^ 2) * x[1]
    grad[2] = 200.0 * (x[2] - x[1] ^ 2)
end
# Define its hessian
function h!(x::Vector, hess::Matrix)
    hess[1, 1] = 2.0 - 400.0 * x[2] + 1200.0 * x[1] ^ 2
    hess[1, 2] = -400.0 * x[1]
    hess[2, 1] = -400.0 * x[1]
    hess[2, 2] = 200.0
end

tolerance = 1.0e-8

println("Optimize with Nelder-Mead")
@time res = optimize(f, [0.0, 0.0])
@test Optim.minimum(res) < tolerance

println("Optimize with LBFGS, with gradient")
@time res = optimize(f, g!, [0.0, 0.0], LBFGS())
@test Optim.minimum(res) < tolerance

println("Optimize with Newton, with gradient and hessian")
@time res = optimize(f, g!, h!, [0.0, 0.0], Newton())
@test Optim.minimum(res) < tolerance

println("Box optimization")
@time res = optimize(OnceDifferentiable(f), [2.0, 2.0], [1.25, -2.1], [Inf, Inf], Fminbox(), optimizer = GradientDescent)
@test 0.06 < Optim.minimum(res) 
@test Optim.minimum(res) < 0.07