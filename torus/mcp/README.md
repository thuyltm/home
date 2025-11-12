# What is Model Context Protocol (MCP)
MCP (Model Context Protocol) is an open-source standard for connecting AI applications to external systems.


Using MCP, AI applications like Claude or ChatGPT can connect to data sources (e.g. local files, databases), 
tools (e.g. search engines, calculators) and workflows (e.g. specialized prompts)--enabling them to access key information
and perform tasks


Modern AI models need context to do intelligent reasoning. MCP helps organize this context in a structured, extensible way:
- What was previously said? (conversation history)
- What tools are available? (tool/function metadata)
- What actions were taken? (and their outputs)
- What the model is expected to do next

![Alt text](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2F7lr9tmc9rnl5f96coirf.png "Optional Title")

The key participants in the MCP architecture are:
- MCP Host: The AI application that coordinates and manages one or multiple MCP clients
- MCP Client: A component that maintains a connection to an MCP server and obtains context from an MCP server for the MCP host to use
- MCP Server: The program that serves context data, regardless of where it runs.MPC servers can execute locally or remotely. For example, the filesystem server, the server runs locally on the same machine because it uses the STDIO transport. This is refered to as a local MCP Server. The Sentry MCP server runs on the Sentry platform, and uses the Streamable HTTP transport. This is refered to as a remote MCP server.
# Data Layer Protocol
## Lifecycle management
MCP is a stateful protocol that requires lifecycle management. The purpose of lifecycle management is to negotiate the capabilities that both client and server support

1. <span style="color:blue">**Initialization (Lifecycle Management)**</span>

The client sends an intialize request to establish the connection and negotiate supported features

**Initialize Request**

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2025-06-18",
    "capabilities": {
      "elicitation": {}
    },
    "clientInfo": {
      "name": "example-client",
      "version": "1.0.0"
    }
  }
}
```
The server response 
- "tools": {"listChanged": true} server can send tools/list_changes notification when its tool list changes
- "resources": {} server also supports the resources primitive
**Initialize Response**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "protocolVersion": "2025-06-18",
    "capabilities": {
      "tools": {
        "listChanged": true
      },
      "resources": {}
    },
    "serverInfo": {
      "name": "example-server",
      "version": "1.0.0"
    }
  }
}
```

## Primitives
They define what clients and servers can offer each other. These primitives specify the types of contextual information that can be shared with AI application and the range of actions


MCP defines 3 core primitives that servers can expose:
### Tools
Excuatable functions that AI applications can invoke to perform actions(e.g. API calls, database queries)
### Resources
Data sources that provide contextual information to AI applications (e.g., file contents, database records, API responses)
### Prompts: 
Reusable templates that help structure interactions with language models (e.g., system prompts, few-shot examples)

2. <span style="color:blue">**Tool Discovery (Primitives)**</span>

**Tools List Request**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list"
}
```
The response contains a **tools** array that provides comprehensive metadata about each available tool
- name: a unique identifier for the tool within the server's namespace
- title: a human-readable display name for the tool that clients can show to users
- description: detailed explanation of what the tool does and when to use it
- inputSchema: a json schema that defines the expected input parameters, enable type validation and providing clear documentation about required and optional parameters

**Tools List Response**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "tools": [
      {
        "name": "calculator_arithmetic",
        "title": "Calculator",
        "description": "Perform mathematical calculations including basic arithmetic, trigonometric functions, and algebraic operations",
        "inputSchema": {
          "type": "object",
          "properties": {
            "expression": {
              "type": "string",
              "description": "Mathematical expression to evaluate (e.g., '2 + 3 * 4', 'sin(30)', 'sqrt(16)')"
            }
          },
          "required": ["expression"]
        }
      },
      {
        "name": "weather_current",
        "title": "Weather Information",
        "description": "Get current weather information for any location worldwide",
        "inputSchema": {
          "type": "object",
          "properties": {
            "location": {
              "type": "string",
              "description": "City name, address, or coordinates (latitude,longitude)"
            },
            "units": {
              "type": "string",
              "enum": ["metric", "imperial", "kelvin"],
              "description": "Temperature units to use in response",
              "default": "metric"
            }
          },
          "required": ["location"]
        }
      }
    ]
  }
}
```
3. <span style="color:blue">**Tool Execution (Primitives)**</span>

After discovering available tools, the client can invoke them with appropriate arguments by excute the **tools/call** method
**Tool Call Request**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "weather_current",
    "arguments": {
      "location": "San Francisco",
      "units": "imperial"
    }
  }
}
```
The response demonstrates MCP's flexible content structure which allows AI applications can integrated into converstation with language model
**Tool Call Response**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Current weather in San Francisco: 68°F, partly cloudy with light winds from the west at 8 mph. Humidity: 65%"
      }
    ]
  }
}
```
## Notifications

The protocol supports real-time notifications to enable dynamic updates between servers and clients. For example, when a server's available tools change--such as when new functionality becomes available or existing tools are modified--the server can send tool update notifications to inform connected clients about these changes.

4. <span style="color:blue">**Real-time Updates (Notification)**</span>


When the server's available tools change--such as when new functionality becomes available, existing tools are modified, or tools become temporarily unavalibale--the server can proactively notify connected clients
```json
{
  "jsonrpc": "2.0",
  "method": "notifications/tools/list_changed"
}
```
Upon reveiving this notification, the client typically reacts by requesting the updated tool list. This creates a refresh cycle that keeps the client;s understanding of available tools current
```json
{
  "jsonrpc": "2.0",
  "id": 4,
  "method": "tools/list"
}
```