```sh
# Check the existence of these libs before installing
% pkg-config --list-all | grep -e mesa -e glu -e glew -e lfw3
% sudo apt-get install mesa-utils libglu1-mesa-dev freeglut3-dev mesa-common-dev
% sudo apt-get install libglew-dev libglfw3-dev
```
- mesa-*: Provides the open-source OpenGL implementation and utilities.
- libglew-dev, libglfw3-dev: Provide libraries for managing OpenGL extensions (GLEW) and creating windows and contexts (GLFW). 

#### Understand the file HOW_TO_RUN_CMAKE.md
```sh
find_package(OpenGL REQUIRED)
target_link_libraries(${PROJECT_NAME} PRIVATE 
    OpenGL::GL 
)
```

In CMake, the __find_package()__ command is used to locate and load external projects/packages, primarily searching for <PackageName>Config.cmake or Find<PackageName>.cmake files.

Use __target_link_libraries()__ to link your executable or library to the package's targets

**Troubleshooting**
If CMake can't find a package, you may need to:
- Set the __CMAKE_PREFIX_PATH__ environment to point to the package's installation firectory before running CMake
- Set the specific __< PackageName >_DIR__ variable to the directory containing the *Config.cmake file
```sh
CMake Debug Log at CMakeLists.txt:7 (find_package):
  find_package considered the following paths for FindOpenGL.cmake:

  The file was found at

    /usr/share/cmake-3.28/Modules/FindOpenGL.cmake



CMake Debug Log at CMakeLists.txt:11 (find_package):
  find_package considered the following paths for FindPkgConfig.cmake:

  The file was found at

    /usr/share/cmake-3.28/Modules/FindPkgConfig.cmake
```
- Run Cmake with --debug-find to get detailed output on where it is searching for packages
```sh
mkdir build
cd build
cmake --debug-find ..
```

#### Build and Run the Project
```sh
mkdir build
cd build
cmake ..
make
./OpenGLProject
```