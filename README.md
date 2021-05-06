# Go backup script
Simple backup script written in go.

```
USAGE: BackupFrom BackupTo Duration

EXAMPLES:

Backup with removing old backups that were created more than 24h ago:
backupscript /home/washingmachine/desktop/new /backup -24h
backupscript.exe C:\Users\washingmachine\Desktop\new C:\Users\washingmachine\Desktop\ -24h

Only backup:
backupscript /home/washingmachine/desktop/new /backup
backupscript.exe C:\Users\washingmachine\Desktop\new C:\Users\washingmachine\Desktop\
```

BackupFrom: Path from where it will copy the files

BackupTo: Path where it will put the .tar.gz backup

Duration (optional): For example, if you enter -24h, the program will delete all backups that were created more than 24 hours ago

### More info about duration

100% working duration units are "m" and "h"

Theoretically you can use "ns", "us" (or "µs"), "ms", "s" too, but I didn't check it if it works.
