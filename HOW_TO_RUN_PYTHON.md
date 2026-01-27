1. Create a virtual environment using Python 3.13 named 'test'
2. Activate the 'test' virtual environment
3. Install library using pip and create a requirements_lock.txt as well
```sh
pip install "fastapi[standard]"
pip freeze > requirements_lock.txt
```
4. Adding dependencies in BUILD.bazel
```sh
py_library(
    name = "---",
    srcs = ["-----"],
    deps = [
        "@pip//fastapi"
    ],
)
```