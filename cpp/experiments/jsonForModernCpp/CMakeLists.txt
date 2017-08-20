cmake_minimum_required (VERSION 3.5)
project (Sandbox)

include(build/conanbuildinfo.cmake)
conan_basic_setup()

add_compile_options (-Wall -Wextra -std=c++14)
add_executable (Sandbox main.cpp)
target_link_libraries(Sandbox ${CONAN_LIBS})
