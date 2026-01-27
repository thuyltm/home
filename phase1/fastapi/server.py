import os
import sys

current_directory = os.getcwd()
sys.path.append(current_directory+"/phase1/fastapi/")

from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}
