using Base.Test
using Currency

cp1 = CurrencyPair("EUR", "USD", 1.1)
@test cp1.spot == 1.1
@test toString(cp1) == "EURUSD"
@test_throws ArgumentError CurrencyPair("EURI", "USD", 1.1)
@test_throws ArgumentError CurrencyPair("EUR", "USDI", 1.1)
@test_throws ArgumentError CurrencyPair("EUR", "EUR", 1.1)
@test_throws ArgumentError CurrencyPair("EUR", "USD", 0.0)
@test_throws ArgumentError CurrencyPair("EUR", "USD", -1.1)

setSpot!(cp1, 1.2)
@test cp1.spot == 1.2
@test_throws ArgumentError setSpot!(cp1, 0.0)
@test_throws ArgumentError setSpot!(cp1, -1.2)

cp2 = CurrencyPair("USD", "JPY", 100.0)
@test cp2.spot == 100.0
@test toString(cp2) == "USDJPY"

cp12 = cp1 * cp2
@test cp12.spot == 120.0
@test toString(cp12) == "EURJPY"

cp3 = CurrencyPair("JPY", "EUR", 0.01)
cp13 = cp1 * cp3
@test cp13.spot == 0.012
@test toString(cp13) == "JPYUSD"