# eliminaro
Remove log files from your projects with a blink of an eye.

Log files from projects can take alot of disk space eventually, to tackle this issue I created eliminaro. A simple little go program to remove them.

## How to install

1. Download the binary at https://github.com/LDonnez/eliminaro/releases
2. Give the binaro file execution permissions -> `chmod +x eliminaro`
3. mv the binary to `/usr/bin/` for installing for all users otherwise move it to `/usr/local/bin`

## How to use

When going to a project with a log directory just type `eliminaro` in the terminal and all the files from the log directory will be removed.
You can also specify a root dir with the `-dir=<directory>` flag.
You can also specify a log dir with the `-log-dir=<log directory> flag.

