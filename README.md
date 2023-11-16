<div align="center">
<pre>

███████╗ ██████╗██████╗ ███████╗███╗   ██╗ ██████╗  ██████╗ 
██╔════╝██╔════╝██╔══██╗██╔════╝████╗  ██║██╔════╝ ██╔═══██╗
███████╗██║     ██████╔╝█████╗  ██╔██╗ ██║██║  ███╗██║   ██║
╚════██║██║     ██╔══██╗██╔══╝  ██║╚██╗██║██║   ██║██║   ██║
███████║╚██████╗██║  ██║███████╗██║ ╚████║╚██████╔╝╚██████╔╝
╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝  ╚═════╝ 
                                                            
---------------------------------------------------
go cli program to make screenshots
</pre>
</div>

Do you use a distribution without a pre-installed screenshot? use this program

## Installation

use the go installer to get the program

```
go install https://github.com/fzbian/screengo
```

or use the releases

[Releases](https://github.com/fzbian/screengo/releases)

## Demostration

https://github.com/fzbian/screengo/assets/66271721/53d94f96-ea84-4ca7-abd7-015bd968595c

## Usage example

To get help with commandline arguments
```
screengo -h
```

List avaliable screens
```
screengo -l
```

Simple screenshot with screen argument
```
screengo <screenId>
```

Screenshot with screen argument and output file name

```
screengo <screenId> <outputFileName>
```

Screenshot with screen argument, output file name and delay in seconds

```
screengo <screenId> <outputFileName> <delaySeconds>
```

## Contributing

Feel free to contribute to this project.

## License

This project is licensed under the [MIT license] - see the [LICENSE FILE](LICENSE) file for details.
