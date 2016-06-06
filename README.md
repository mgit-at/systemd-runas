# systemd-runas

Transfer PID to new systemd scope and slice for different ressource accounting, callable as user.

## How it works

systemd-runas is called with the name of the slice the program want's the pid be transferred

```
sudo systemd-runas convert 1234
```

This call transferes PID 1234 into scope convert-1234.scope and applies ressource accounting from convert.slice to it.

## Example sudoers entry

```
tomcat ALL = NOPASSWD: /usr/local/bin/systemd-runas convert [0-9]*
```
