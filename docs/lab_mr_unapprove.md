## lab mr unapprove

Unapprove merge request

```
lab mr unapprove [remote] <id> [flags]
```

### Examples

```
lab mr unapprove origin
lab mr unapprove upstream -F test_file
lab mr unapprove upstream -m "A helpfull comment"
lab mr unapprove upstream --with-comment
lab mr unapprove upstream -m "A helpfull\nComment" --force-linebreak
```

### Options

```
  -F, --file string           use the given file as the message
      --force-linebreak       append 2 spaces to the end of each line to force markdown linebreaks
  -h, --help                  help for unapprove
  -m, --message stringArray   use the given <msg>; multiple -m are concatenated as separate paragraphs
      --with-comment          Add a comment with the approval
```

### Options inherited from parent commands

```
      --debug      Enable debug logging level
      --no-pager   Do not pipe output into a pager
      --quiet      Turn off any sort of logging. Only command output is printed
```

### SEE ALSO

* [lab mr](lab_mr.md)	 - Describe, list, and create merge requests

###### Auto generated by spf13/cobra on 30-Sep-2021
