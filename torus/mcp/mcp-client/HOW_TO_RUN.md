# Setup a Python project
You create a new python project with **uv**
```sh
# Create project directory
uv init mcp-client
cd mcp-client

# Create virtual environment
uv venv

# Activate virtual environment
source .venv/bin/activate

# Install required packages
uv add mcp anthropic python-dotenv

# Remove boilerplate files
rm main.py

# Create our main file
touch client.py
```
# Create an API key
Create a key to integrate with the Claude API from the Anthropic Console. You can use the API directly or through a client SDK

# Run the client
To run your client with any MCP server
```sh
uv run client ABSOLUTE_PATH_TO/mcp-server/weather.py
```