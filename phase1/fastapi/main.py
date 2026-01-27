import os
import sys

current_directory = os.getcwd()
sys.path.append(current_directory+"/phase1/fastapi/")

from my_lib import get_array
import requests

def main():
    print(f"Numpy array: {get_array()}")
    print(f"Requests version: {requests.__version__}")

if __name__ == "__main__":
    main()
