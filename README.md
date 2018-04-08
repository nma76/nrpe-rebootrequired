# nrpe-rebootrequired
This plugin monitors the need for reboot on Ubuntu/Debian systems. It also shows the reason reboot is required. Should work on Ubunto 8.04 and newer. Possibly even older systems if you install the update-notifier-common package. Other dists are unknown and untested.  

**Tested platforms**:  
Ubuntu 16.04 LTS  
  
**Currently this plugin return status**  

| Status   | Message                                                                |
| -------- | ---------------------------------------------------------------------- |
| OK       | when no reboot is required                                             |
| WARNING  | when reboot is required. Reason for the rebbot is shown in the message |
| CRITICAL | when reboot has been required for more than 2 days (48 hours)          |
| UNKNOWN  | when the plugin is unable to get any status                            |

**Compile**  
Just make sure yoou have Go installed and run:  
- Go build

**TODO**  
- Implement a cleaner way to handle messages