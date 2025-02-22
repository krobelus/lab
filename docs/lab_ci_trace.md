## lab ci trace

Trace the output of a ci job

### Synopsis

Download the CI pipeline job artifacts for the given or current branch if
none provided. If a job is not specified the latest running job or last
job in the pipeline is used

The branch name, when using with the --merge-request option, can be the
merge request number, which matches the branch name internally.	The "job"
portion is the given job name, which may contain whitespace characters
and which, for this specific case, must be quoted.

```
lab ci trace [remote] [branch[:job]] [flags]
```

### Examples

```
lab ci trace upstream feature_branch
lab ci trace upstream 18 --merge-request
lab ci trace upstream 18:'my custom stage' --merge-request
lab ci trace upstream 18:'my custom stage' --merge-request --bridge 'security-tests'
```

### Options

```
  -h, --help            help for trace
      --merge-request   use merge request pipeline if enabled
```

### Options inherited from parent commands

```
      --bridge string   Bridge job (downstream pipeline) name
      --debug           Enable debug logging level
      --follow          Follow bridge jobs (downstream pipelines) in a multi-projects setup
      --no-pager        Do not pipe output into a pager
      --quiet           Turn off any sort of logging. Only command output is printed
```

### SEE ALSO

* [lab ci](lab_ci.md)	 - Work with GitLab CI pipelines and jobs

###### Auto generated by spf13/cobra on 30-Sep-2021
