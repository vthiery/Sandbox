using TimeSeries
using MarketData

dates  = collect(Date(1999, 1, 1):Date(2000, 12, 31))
mytime = TimeArray(dates, rand(length(dates)))

# print all the columns for open high low close
print(ohlc[1:3])
# print specific columns
print(ohlc["Open", "Close"])
# print head and tail
print(head(ohlc))
print(tail(ohlc))
# filter on mondays
print(when(cl, dayname, "Monday"))
# update the ts
print(update(cl, Date(2002, 1, 1), 111.11))