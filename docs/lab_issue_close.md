## lab issue close

Close issue by ID

```
lab issue close [remote] <id> [flags]
```

### Examples

```
lab issue close 1234
lab issue close origin --duplicate 123 1234
lab issue close --duplicate other-project#123 1234
```

### Options

```
      --duplicate string   mark as duplicate of another issue
  -h, --help               help for close
```

### Options inherited from parent commands

```
      --debug      Enable debug logging level
      --no-pager   Do not pipe output into a pager
      --quiet      Turn off any sort of logging. Only command output is printed
```

### SEE ALSO

* [lab issue](lab_issue.md)	 - Describe, list, and create issues

###### Auto generated by spf13/cobra on 30-Sep-2021
