### Stop Tracking a Folder
If you accidentally committed a folder and want Git to stop tracing its changes from now on without deleting it from your computer
1. Open your .gitignore file and add the folder path
2. Remove the folder from Git's tracking index
```sh
git rm -r --cached path_to_your_folder
```
### Unstaged changes after reset
When you see the message "Unstaged changes after reset", it means Git has successfully removed your changes from the staging area (index).

If your goal was to revert your files back to the last clean commit
```sh
git restore .
```