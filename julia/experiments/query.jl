using Query, DataFrames

# Simple query with filtering
df = DataFrame(name=["John", "Sally", "Kirk"], age=[23., 42., 59.], children=[3,5,2])

x = @from i in df begin
    @where i.age > 50
    @select {i.name, i.children}
    @collect DataFrame
end

println(x)

# Projecting

data = [1,2,3]

x = @from i in data begin
    @select i^2
    @collect
end

println(x)

# Old friend inner join

df1 = DataFrame(a=[1,2,3], b=[1.,2.,3.])
df2 = DataFrame(c=[2,4,2], d=["John", "Jim","Sally"])

x = @from i in df1 begin
    @join j in df2 on i.a equals j.c
    @select {i.a,i.b,j.c,j.d}
    @collect DataFrame
end

println(x)
