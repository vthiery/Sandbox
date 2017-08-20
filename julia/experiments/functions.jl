using Base.Test

@test map(x -> x^2 - x + 1, [2, 3, -1, 0, 4]) == [3, 7, 3, 1, 13]

###################

function exponential(exponent)
    a -> a ^ exponent
end

toThe2 = exponential(2)
toThe3 = exponential(3)

@testset "toTheX Tests" begin
    @testset "Positive" for i in 1:100
        @test toThe2(i) == i ^ 2
        @test toThe3(i) == i ^ 3
    end
    @testset "Negative" for i in 1:100
        @test toThe2(i) == i ^ 2
        @test toThe3(i) == i ^ 3
    end
end

###################

function greet(x)
    str = x()
    return "Hello, $str\!"
end

function country()
    return("Switzerland")
end

@test greet(country) == "Hello, Switzerland!"

###################

function oper(a, b ,c)
    return a(b, c)
end

@test oper(*, 3.2, 10.0) == 32.0
@time oper(*, 3.2, 10.0)

###################

function dotproduct(x::Array, y::Array)
    result = zero(eltype(x))
    for i = 1:length(x)
        @inbounds result += x[i] * y[i]
    end
    return result
end

@time(dotproduct([1234, 5747, 2243243, 535345, 76345, 2346], [23468, 4563, 2457, 124556, 27563, 63456]))
