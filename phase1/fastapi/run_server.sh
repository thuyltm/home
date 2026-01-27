#!/bin/sh

# fastapi run phase1/fastapi/server.py --port 80 --workers 2
pwd
ls -R
cd phase1/fastapi
uvicorn server:app --reload
fastapi run server.py --port 80 --workers 2