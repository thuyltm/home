[Command and Options](https://bazel.build/docs/user-manual)

#### --copt=cc-option

This option takes an argument which is to be passed to the compiler. The argument will be passed to the compiler whenever it is invoked for preprocessing, compiling, and/or assembling C, C++, or assembler code. It will not be passed when linking.
```sh
% bazel build --copt="-g0" --copt="-fpic" //foo
```

#### --cxxopt=cc-option

This option takes an argument which is to be passed to the compiler when compiling C++ source files.

This is similar to --copt, but only applies to C++ compilation, not to C
```sh
% bazel build --cxxopt="-fpermissive" --cxxopt="-Wno-error" //foo/cruddy_code
```

#### --linkopt=linker-option

This option takes an argument which is to be passed to the compiler when linking.

This is similar to --copt, but only applies to linking, not to compilation