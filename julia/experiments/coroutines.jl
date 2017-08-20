using Base.Test

function yield()
    produce(0)
    for n = 1:5
        produce(3n)
    end
end

k = 0
for n in Task(yield)
    @test n == 3k
    k += 1
end

#######################

function findPi(n)
    inside = 
        @parallel (+) for i = 1:n
            x, y = rand(2)
            x^2 + y^2 <= 1 ? 1 : 0
        end
    4 * inside / n
end

println(findPi(10000))
