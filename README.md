# nrpe-rebootrequired
This plugin monitors the need for reboot on Ubuntu/Debian systems. It also shows the reason reboot is required. Should work on Ubuntu 8.04 and newer. Possibly even older systems if you install the update-notifier-common package. Other dists are unknown and untested.  

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
- Make sure you have Go installed  
- Run Go build

**Install**  
Copy the binary (nrpe-rebootrequired) to your Nagios plugin folder. On Ubuntu this is typically located at /usr/lib/nagios/plugins. Make sure it's executable. You can run chmod +x nrpe-rebootrequired to make it executable.  
  
Add a command to your nagios configuration. On Ubuntu systems you usually add commands to /etc/nagios/nrpe_local.cfg. A command has the syntax:   
`command[<Commandname>]=<Executable>`  
  
For this check you could add the following command:  
`command[check_rebootrequired]=/usr/lib/nagios/plugins/nrpe-rebootrequired`  
  
Make sure to restart the nagios/nrpe service before you use this command. On Ubuntu you can execute `/etc/init.d/nagios-nrpe-server restart`
  
On your Nagios/OP5 server, create a new check of the type check_nrpe and enter check_rebootrequired as the command.

**TODO**  
- Implement a cleaner way to handle messages
- Show for how long reboot benn waiting
- Parameterize timespan before status is set to CRITICAL
