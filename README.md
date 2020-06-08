# Git clone

git-clone is a simple command that as its name mention it clone a git repository into a local directory.

## Parameters

| Environment Variable      | Description                                                                                      | Default |
|---------------------------|--------------------------------------------------------------------------------------------------|---------|
| GIT_CLONE_REPO            | the git repository to clone                                                                      | ""      |
| GIT_CLONE_REF             | the git revision (branch, tag or hash) to check out                                              | "main"  |
| GIT_CLONE_ROOT_DIR        | the root directory for git-sync operations, under which --dest will be created                   | "/tmp"  |
| GIT_CLONE_DEST_DIR        | the name of a directory in which to checkout files under root (defaults to the leaf dir of repo) | ""      |
| GIT_CLONE_USERNAME_SECRET | the username to use for git auth                                                                 | ""      |
| GIT_CLONE_PASSWORD_SECRET | the password to use for git auth (users should prefer env vars for passwords)                    | ""      |
| GIT_CLONE_TIMEOUT         | the max number of seconds allowed for the clone to complete                                      | 120     |
| GIT_CLONE_DEBUG           |                                                                                                  | false   |
