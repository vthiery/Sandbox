# Start with julia -n N where N is the number of workers

r = @spawn rand(2, 2)
s = @spawn 1 .+ fetch(r)

fetch(s)

###################

# Good practices

# Bad
# A = rand(1000, 1000)
# B = @spawn A^2
# fetch(B)

# Good
B = @spawn rand(1000, 1000)^2
fetch(B)

###################

nheads = @parallel (+) for i = 1:200000000
    Int(rand(Bool))
end
