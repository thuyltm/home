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
### Git hang and reject because of a large file
To fix a _git push_ that is hanging or rejected because of a large file, simply deleting the file normally and making a new commit will not work, because the large file remains hidden in your previous unpushed commit

This approach resets your repository back to match the remote server, then remove the large file from Git's memory and lastly recommit your valid changes
```sh
git reset --soft HEAD
# Untrack and remove the large file from Git's memory
git rm --cached path/to/large_file.ext
# Find files larger than 50MB and append them to .gitignore
find . -type f -size +50M -not -path "./.git/*" | sed 's|^\./||g' >> .gitignore
# Recommit your valid changes
git add .
git commit -m "Recommit changes"
git push origin main
```
### View the list of commits
Check which commmits will be sent to the remote repository when you run _git push origin main
```sh
git log origin/main..main --oneline
```