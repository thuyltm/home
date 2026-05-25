Your Flask app works locally but cannot reachable inside a Docker container even though you already established the port mapping. This error occurs because Flask is bound to 127.0.0.1 which only accepts for request originating from inside the container itself

To make it accessible from outside of your host machine, you must bind it to all interface (0.0.0.0)
- Fix in Python code
```sh
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
```
- Fix in Dockerfile
```sh
CMD ["flask", "run", "--host=0.0.0.0"]
```
The console output will appear as show below
```sh
# * Serving Flask app 'rolldice'
# * Debug mode: off
# INFO:werkzeug:WARNING: This is a development server. Do not use it in a production deployment. Use a production WSGI server instead.
# * Running on all addresses (0.0.0.0)
# * Running on http://127.0.0.1:5000
# * Running on http://172.17.0.2:5000
# INFO:werkzeug:Press CTRL+C to quit
```