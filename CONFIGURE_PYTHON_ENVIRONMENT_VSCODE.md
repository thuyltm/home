### Configure an existing Python enviroment in VSCode
To configure an existing Python environment in VSCode, you use the **Select Interpreter** command to set the editor to run your specific environment's execution

Follow these steps in order to succeed
1. Press **Ctrl+Shift+P** to open VSCode command
2. Type **Python: Select Interpreter** and select it
3. A list of detected environment will appear. Choose one to activate it for the current workspace

### Configure Pylan to be default language server for Python in VSCode
Flow these steps in sequence to achieve success
1. Install the Python Extension: Open the VS Code Marketplace and install Pylance
2. Activate: Set **"python.languageServer=Pylance"**

    2.1. Open the **Command Palette** by pressing **Ctrl + Shift + P**

    2.2. Type **"Open Settings"** and configure **"python.languageServer=Pylance"**

    2.3. Or **Select Preferences: Open User Setting (JSON)** to edit the raw **settings.json** file and add more the line **"python.languageServer=Pylance"**
3. Verify: Open any .py file, Pylance will activate immediately to start analyzing your code

