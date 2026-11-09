1. To list the files in the specific commit or a local refrence
```sh
git show --name-only <localref>
```
2. To display the content of the specific file at a particular Git revison or local reference
```sh
git show <localref>:<path/to/file>
```
3. Redirect the standard ouput (stdout) to a file using the > operator
```sh
git show <localref>:<path/to/file> > newfile.md
```