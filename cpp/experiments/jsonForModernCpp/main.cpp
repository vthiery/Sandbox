#include <iostream>
#include <json.hpp>

using Json = nlohmann::json;

const auto find = [](const Json& input, const std::string& node)
{
    static const std::string InvalidData{ "Invalid data" };
    static const std::string MissingData{ "Missing data" };

    try
    {
        return input.at(node);
    }
    catch (const std::domain_error& e)
    {
        std::cout << InvalidData  << std::endl;
        throw;
    }
    catch (const std::out_of_range& e)
    {
        std::cout << MissingData << std::endl;
        throw;
    }
};

int main()
{
    const Json input = R"(
    {
        "a": [
            { "a1": 1 },
            { "a2": 2 }
            ],
        "b": "coco"
    })"_json;

    const Json c1{ find(input, "a") };
    std::cout << c1 << std::endl;
    // prints [[{"a1":1},{"a2":2}]]

    const Json c2 = find(input, "a");
    std::cout << c2 << std::endl;
    // prints [{"a1":1},{"a2":2}]

    Json&& c3{ find(input, "a") };
    std::cout << c3 << std::endl;
    // prints [{"a1":1},{"a2":2}]

    Json&& c4 = find(input, "a");
    std::cout << c4 << std::endl;
    // prints [{"a1":1},{"a2":2}]

    return 0;
}

