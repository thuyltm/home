cmake_minimum_required(VERSION 3.28)
project(OpenGLProject C)

# Debug only the finding of a specific package
set(CMAKE_FIND_DEBUG_MODE ON)
# Find OpenGL (core library)
find_package(OpenGL REQUIRED)
set(CMAKE_FIND_DEBUG_MODE OFF)

# Use PkgConfig to find GLFW and GLEW (windowing and extension libraries)
find_package(PkgConfig REQUIRED)
pkg_check_modules(GLFW REQUIRED IMPORTED_TARGET glfw3)
pkg_check_modules(GLEW REQUIRED IMPORTED_TARGET glew)

# Add your executable target
add_executable(${PROJECT_NAME} helloopengl.c)

# Link the libraries to your executable
target_link_libraries(${PROJECT_NAME} PRIVATE 
    OpenGL::GL 
    PkgConfig::GLFW 
    PkgConfig::GLEW
)
