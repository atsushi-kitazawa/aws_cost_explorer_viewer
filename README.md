# aws cost explorer viewer
This is CLI tool for checking the status of AWS cost.

This tool is written by Go.

# install
```sh
% git clone https://github.com/atsushi-kitazawa/aws_cost_explorer_viewer.git
% go install
```

If your environments is Linux or Windows, Execute file are also available.

[Release 1.0 · atsushi-kitazawa/aws_cost_explorer_viewer](https://github.com/atsushi-kitazawa/aws_cost_explorer_viewer/releases/tag/1.0)

# usege
Specify target region, username, apikey and secretkey in setting.yaml.

The value of username can be anything.

rename setting.yaml.template to setting.yaml and place in the same directory as the command.
```sh
% ls
aws_cost_explorer_viewer   setting.yaml

% ./aws_cost_explorer_viewer
```

# target period
The target period is the month of execution.

You can specify the target period with command line arguments.
```sh
% ./aws_cost_explorer_viewer 2021-02-01 2021-02-28
```
