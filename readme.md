# Bills
Cli to remind you to pay your bills using notify-send

## Usage
- `bills add [name] [day-of-month]`: this will add a new bill
- `bills set [name] [day-of-month]`
- `bills remove [name]`
- `bills list`
    - this will list all the bills available
    - `--json`
        - this will list available bills in json format
- `bills notify`: can be combined with `bills`
    - this will notify as early as three days using `notify-send` for every bill until they're marked as paid on the current month
- `bills paid [name]`
- `bills reset`: this should be run every day. It will check if the current day
  is the first of the month and then reset the bills and mark them as unpaid

Now, to make it work you will need tell your system to execute `bills reset` and
`bills notify` on each startup. Depending on your desktop environment, window
manager or operating system, you can do that in different ways.

## Installation
Installation is still pending. For now, you can clone the repo and install it
using `go install`.


## TODO:
- Better ui?
- Testing
- Installation from AUR
- Make it work with systemd? (prehaps to avoid the need for users to add two
commands to their startup scripts)
