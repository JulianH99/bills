# Bills

- `bills remove [name]`
- `bills set [name] [day-of-month]`
- `bills calcurse` : this will set a recurring event in calcurse
    -  ! seems like it has to be the other way, which is not the best approach
- `bills list`
    - this will list all the bills available
    - `--json`
        - this will list available bills in json format
- `bills notify`: can be combined with `bills`
    - this will notify as early as three days using `notify-send` for every bill until they're marked as paid on the current month
- `bills paid [name]`
- `bills startup`:  can be `bills` for short
    - will start the service to check every day if the bills need to be reset or not (due to a start of month) 
