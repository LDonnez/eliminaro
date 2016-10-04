# eliminaro
Remove log files from your projects with a blink of an eye.

Log files from projects can take alot of disk space eventually, to tackle this issue I created eliminaro. A simple little go program to remove them.

Besides removing log files, Eliminaro also removes the already merged git branches from your project.

## How to install

1. Download the binary at https://github.com/LDonnez/eliminaro/releases
2. Give the binaro file execution permissions -> `chmod +x eliminaro`
3. mv the binary to `/usr/bin/eliminaro` for installing for all users otherwise move it to `/usr/local/bin/eliminaro`

## How to use

When going to a project with a log directory just type `eliminaro` in the terminal and all the files from the log directory will be removed.

### Currently available flags

`-dir=<directory` Specify a root directory

`-log-dir=<log dir>` Specify a log directory where files needs to be removed.

### Currently implemented features

- [x] Remove log files from the log directory in your project
- [x] Remove merged git branches
