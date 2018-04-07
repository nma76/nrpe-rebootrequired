# nrpe-rebootrequired
This plugin monitors the need for reboot on Ubuntu/Debian systems. It also shows the reason reboot is required.

This first version is only tested on my developer machine which runt on MacOS 10.13.3, and itäs only tested with simulated files.

Currently this plugin return status:
OK      - when no reboot is required
WARNING - when reboot is required. Reason for the rebbot is shown in the message.
UNKNOWN - when the plugin is unable to get any status