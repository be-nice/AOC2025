cmake_minimum_required(VERSION 3.10)
project(AOC C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)
set(CMAKE_C_EXTENSIONS OFF)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -O3")
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR})

file(GLOB DAY1_SOURCES "day_1/*.c")

add_executable(AOC
    main.c
    ${DAY1_SOURCES}
)

target_include_directories(AOC PRIVATE day_1)

target_compile_options(AOC PRIVATE
    -Wall
    -Wextra
    -Wpedantic
    -Wshadow
    -Wconversion
    -Wdouble-promotion
    -Wformat=2
    -Wundef
    -Werror
)
