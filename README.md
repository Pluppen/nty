# Nty
A simple (https://ntfy.sh/)[https://ntfy.sh/] wrapper written in Go.

# Use case
```bash
long_running_command && ntfy -msg 'Long running command completed successfully' || ntfy -msg 'Long running command failed.' -priority 'Urget'
```
