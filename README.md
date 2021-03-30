# aws cost explorer viewer
This is CLI tool for checking the status of AWS cost.

This tool is written by Go.

![](https://user-images.githubusercontent.com/47269784/112724853-92ef9180-8f58-11eb-9e6d-26acd851a04a.gif)

# install
```sh
% git clone https://github.com/atsushi-kitazawa/aws_cost_explorer_viewer.git
% go install
```

If your environments is Linux or Windows, Execute file are also available.

[Release 1.1](https://github.com/atsushi-kitazawa/aws_cost_explorer_viewer/releases/tag/1.1)

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
The target period is the month of this tool execution.

You can specify the target period with command line arguments.
```sh
% ./aws_cost_explorer_viewer -start 2021-02-01 -end 2021-03-01
```
