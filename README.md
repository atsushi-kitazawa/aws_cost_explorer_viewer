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
Set apikey and secretkey in setting.yaml.

The value of name can be anything.

Specify the start time and end time for checking the cost.
```sh
% ls
aws_cost_explorer_viewer   setting.yaml

% ./aws_cost_explorer_viewer 2021-02-01 2021-02-28
```
