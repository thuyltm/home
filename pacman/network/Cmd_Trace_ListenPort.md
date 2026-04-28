Running the command below to list active ports
```sh
sudo lsof -i -P -n | grep LISTEN
# btcd      66463            thuy   10u  IPv4 123360      0t0  TCP 127.0.0.1:9333 (LISTEN)
# btcd      66463            thuy   12u  IPv4 121427      0t0  TCP 127.0.0.1:18556 (LISTEN)
# btcwallet 66847            thuy    6u  IPv4 126106      0t0  TCP 127.0.0.1:18554 (LISTEN)
```