#### GENERAL
|||
|-|-|
`git config --global alias.pm "push origin master"`|Make git alias, to use: ($ git pm).
`git config core.pager 'less -r'`|Set pager with no wrap lines.
`git show some-branch:some-file.txt`|Show file from another branch.
`git config --list`|See git settings.
`git clone -b <branch_name> https://url.com/user/repository.git`|Clone repository on specific branch.
`git rev-parse --show-toplevel`|Get the git root directory
`git commit --amend -m "New commit message"`|Change message of last commit.
`git checkout [branch]`|Change current branch.
`git checkout [file]`|Reverses changed file the latest version in the repository.
`git rm $(git ls-files --deleted)`|Remove from repository all locally deleted files.
`git blame <file>`|Show who authored each line in file.
`git clean -df`|Delete not monitored files.
`git commit --amend`|Fix your previous commit, instead of making a new commit.
`git config --global user.name "YOUR NAME"; git config --global user.email "YOUR EMAIL"`|Set your details on commits.
`git show COMMIT_ID`|See details (log message, text diff) of a commit.
#### BASICS
|||
|-|-|
`git clone <repo>`|Clone repo located at <repo> onto local machine.
`git status`|List which files are staged, unstaged, and untracked.
`git add <directory/file>`|Stage all changes in <directory> for the next commit.
`git init <directory>`|Create empty Git repo in specified directory. Run with no arguments to initialize the current directory.
#### SEARCH
|||
|-|-|
`git grep 'SEARCH_WORD' $(git ls-remote . 'refs/remotes/*' | cut -f 2)`|Find text in all branches.
#### PUSH
|||
|-|-|
`git push <remotename> <commit SHA>:<remotebranchname>`|Pushing specific commit to a remote, and not the previous commits.
#### TAG
|||
|-|-|
`git tag <tag_name>`|Create a tag.
`git push origin <tag_name>`|Push a single tag.
`git push --tags`|Push all tags.
`git describe --tags `git rev-list --tags --max-count=1``|See closest tag.
#### STASHING
|||
|-|-|
`git stash save [msg]`|Save your local modifications to a new stash, and run ($ git reset –hard) to revert them.
`git stash list`|List all current stashes.
`git stash apply [stash]`|Restore the changes recorded in the stash on top of the current working tree state. 
`git stash pop [stash]`|Remove a single stashed state from the stash list and apply on top of the current working tree state.
`git stash clear`|Remove all the stashed states.
`git stash drop [stash]`|Remove a single stashed state from the stash list.
`git stash branch <new-branch> [stash]`|Creates and checks out a new branch.
#### REMOTES
|||
|-|-|
`git remote add remote <url>`|Adds a remote named remote for the repository at url.
`git rm remote <url>`|Remove reference to remote repository.
`git remote set-url origin https://seuusuario@github.com/usuario/repositorio.git`|Resolves error: 403 while accessing.
#### HISTORIC
|||
|-|-|
`git log --follow [file]`|Historical list of a file, including renames.
`git log -p [arquivo]`|History of changes for file.
`git log --oneline --graph`|Parameters for better logging.
`git log --author="AUTHOR NAME" --after="1 week ago"`|Show all commits that happened after certain date and by specific author.
`git log --author="AUTHOR NAME" --before="1 week ago"`|Show all commits that happened before certain date and by specific author.
`git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit`|Beautiful git log.
`git log`|Version history lists for the current branch.
#### DIFFERENCES
|||
|-|-|
`git diff [file]`|Shows alterations file in stage.
`git diff COMMIT1_ID COMMIT2_ID`|See differences between two commits.
`git diff --name-only COMMIT1_ID COMMIT2_ID`|See the files changed between two commits.
#### RESET
|||
|-|-|
`git reset [commit]`|Undoes all after commits [commit], preserves locally alterations.
`git reset --hard [commit]`|Discards all changes and returns to specific commit.
#### BRANCH
|||
|-|-|
`git checkout -b <new_branch_name> <name_of_an_existing_branch>`|Create new branch from existing branch.
`git branch -a`|See all branches.
`git fetch origin`|Get all branches.
`git checkout -b <branch_name>`|Create a new branch.
`git checkout -d <branch_name>`|Delete branch.
#### MERGE
|||
|-|-|
`git merge <another_branch>`|Merge current branch with another branch.
`git merge <branch_A> <branch_B>`|Get the changes from <branch_A> and <branch_B> back to current branch.
`git merge <branch-name> --no-commit --no-ff`|Merge a branch without committing.
