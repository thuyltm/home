1. Create a virtual environment using Python 3.13 named 'myenv'
```sh
conda create --name myenv python=3
```
2. Activate the 'myenv' virtual environment
```sh
conda activate myenv
```
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
5. To install dependencies from a requirements_lock file, run the following the command in your terminal
```sh
pip install -r requirements_lock.txt
```