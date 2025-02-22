## lab mr merge

Merge an open merge request

### Synopsis

Merges an open merge request. If the pipeline in the project is
enabled and is still running for that specific MR, by default,
this command will sets the merge to only happen when the pipeline
succeeds.

```
lab mr merge [remote] <id> [flags]
```

### Examples

```
lab mr merge origin 10
lab mr merge upstream 11 -i
```

### Options

```
  -h, --help        help for merge
  -i, --immediate   merge immediately, regardless pipeline results
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
