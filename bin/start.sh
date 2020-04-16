#!/bin/sh -e

## Helper script for starting a system daemon.
## Used in systemd unit file.
##
## Redirects all daemon's output to the rotated log file within /var/log.

/usr/sbin/logan-server 2>&1 | /usr/bin/reopener -s /var/log/logan/messages.log
