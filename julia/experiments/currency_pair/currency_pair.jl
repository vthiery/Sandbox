module Currency

export CurrencyPair, setSpot!, toString

type CurrencyPair
    den::String
    num::String
    spot::Float64

    function CurrencyPair(den::String, num::String, spot::Float64)
        if length(den) != 3
            throw(ArgumentError("Denominator ISO must have 3 characters"))
        elseif length(num) != 3
            throw(ArgumentError("Numerator ISO must have 3 characters"))
        elseif den == num
            throw(ArgumentError("Denominator and numerator must be different"))
        end

        if spot <= 0
            throw(ArgumentError("Spot cannot be negative"))
        end
        new(den, num, spot)
    end
end

function *{CurrencyPair}(cp1::CurrencyPair, cp2::CurrencyPair)
    if cp1.num != cp2.den && cp1.den != cp2.num
        throw(ArgumentError("Incompatible currency pairs"))
    elseif cp1.num == cp2.den
        den = cp1.den
        num = cp2.num
        spot = cp1.spot * cp2.spot
        CurrencyPair(den, num, spot)
    else
        den = cp2.den
        num = cp1.num
        spot = cp1.spot * cp2.spot
        CurrencyPair(den, num, spot)
    end
end

function setSpot!(cp::CurrencyPair, newSpot::Float64)
    if newSpot <= 0
        throw(ArgumentError("Spot cannot be negative"))
    end
    cp.spot = newSpot
end

function toString(cp::CurrencyPair)
    return cp.den * cp.num
end

end