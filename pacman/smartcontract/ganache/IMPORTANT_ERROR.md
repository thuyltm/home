# Install Ganache
On Ubuntu 22.04 and later, installing the older fuse package might conflict with fuse3, which is used by the system for desktop functionality, potentially removing critical packages

When the gdm.service fails with a core-dump result, it means the process encountered a segmentation fault. This usually happens due to driver conflict, corrupted configuration file or missing system dependencies

To see exactly why it crashed, use journalctl to view the specific error leading to the dump
```sh
journalctl -u gdm.service -b
```

Reinstall GDM and GNOE shell
```sh
sudo apt install --reinstall gdm3 gnome-shell ubuntu-desktop
```