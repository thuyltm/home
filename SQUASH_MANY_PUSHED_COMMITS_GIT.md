### You can squash multiple submitted (pushed commits) into a single commit using the Git command line

1. Open the rebase editor
```sh
git rebase -i HEAD~3
# OR
git rebase -i <PARENT_COMMIT_SHA>
```
2. Mark commits to squash

Your text editor will open with a list of commits, ordered from oldest (top) to newest (bottom)
- Leave the top commit as pick
- Change _pick_ to _squash_ for all the subsequent commits below it if you want to rewrite the commit message
- OR change _pick_ to _fixup_ for all the subsequent commits below it if you want to discard the old commit messages completely
```sh
pick ...
fixup ...
pick ....
pick .....
```
3. Sava and close
4. Force-push to GitHub
```sh
git push --force-with-lease origin <branch-name>
```

### You can restore a deleted commit that was already pushed and subsequently removed from the remote server
1. Locate the missing commit hash

Even if a commit was deleted from the branch history via a force-push, Git retains a local log of all reference updates
```sh
git reflog
```
Look for the 7-character alphanumeric hash next to the commit message you want to restore
2. Move your local branch pointer
```sh
git reset --hard <commit_hash>
```
3. Force-Push the restored commit remotely
```sh
git push --force origin <branch_name>
```