# Setup a Python project
You install **uv** and set up a Python project and environment
```sh
# Create a new directory for our project
uv init weather
cd weather

# Create virtual environment and activate it
uv venv
source .venv/bin/activate

# Install dependencies
uv add "mcp[cli]" httpx

# Create our server file
touch weather.py
```
# Configure MCP server for Claude for Desktop
We need to configure Claude for Desktop for which MCP servers you want to use. To do this, edit the file **claude_desktop_config.json**
```sh
code /home/thuy/.config/Claude/claude_desktop_config.json
```
You then add your servers in the **mcpServers** key
```json
{
  "mcpServers": {
    "weather": {
      "command": "uv",
      "args": [
        "--directory",
        "/ABSOLUTE/PATH/TO/PARENT/FOLDER/weather",
        "run",
        "weather.py"
      ]
    }
  }
}
```
This tells Claude for Desktop to launch an MCP server named "weather" by running **_uv --directoy /ABSOLUTE/PATH/TO/PARENT/FOLDER/weather run weather.py_**

# RUN
Since this is the US National Weather service, the queries will only work for US locations. You can test your server by running the following commands in Claude for Desktop:
- What's the weather in Sacramento?
- What are the active weather alers in Texas?