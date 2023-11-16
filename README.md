<div align="center">
<pre>

░██████╗░░█████╗░██████╗░███████╗░█████╗░░█████╗░██████╗░██████╗░███████╗██████╗░
██╔════╝░██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗
██║░░██╗░██║░░██║██████╔╝█████╗░░██║░░╚═╝██║░░██║██████╔╝██║░░██║█████╗░░██████╔╝
██║░░╚██╗██║░░██║██╔══██╗██╔══╝░░██║░░██╗██║░░██║██╔══██╗██║░░██║██╔══╝░░██╔══██╗
╚██████╔╝╚█████╔╝██║░░██║███████╗╚█████╔╝╚█████╔╝██║░░██║██████╔╝███████╗██║░░██║
░╚═════╝░░╚════╝░╚═╝░░╚═╝╚══════╝░╚════╝░░╚════╝░╚═╝░░╚═╝╚═════╝░╚══════╝╚═╝░░╚═╝
---------------------------------------------------
go cli program to make screenshots
</pre>
</div>

Do you use a distribution without a pre-installed screenshot? use this program

## Installation

use the go installer to get the program

```
go install https://github.com/fzbian/gorecorder
```

or use the releases

[Releases](https://github.com/fzbian/gorecorder/releases)

## Demostration

video

## Usage example

To get help with commandline arguments

```
gorecorder -h
```

List avaliable screens
```
gorecorder -l
```

Simple screenshot with screen argument

```
gorecorder <screenId>
```

Screenshot with screen argument and output file name

```
gorecorder <screenId> <outputFileName>
```

Screenshot with screen argument, output file name and delay in seconds

```
gorecorder <screenId> <outputFileName> <delaySeconds>
```

## Contributing

Feel free to contribute to this project.

## License

This project is licensed under the [MIT license] - see the [LICENSE FILE](LICENSE) file for details.