# Docker compare history

## NAME:
```
   docker-history-compare - Docker image history compare
```

## Why did we built it
For some of our clients we are using 3rd party services like CircleCI, CodeShip.
These services has a caching mecanisem that allows us to save a Docker image and next build to load it from cache to Docker.
We needed a tool to tell us if layers the in the image changed and only then to save it to cache.
This saved us ~20 sec for every build since we didn't had to save to cache already exsits images.

## USAGE:
```
   docker-history-compare [global options] command [command options] [arguments...]
```

## VERSION:
```
   0.0.1
```

## COMMANDS:
```
	help, h  Shows a list of commands or help for one command
```

## GLOBAL OPTIONS:
```
   --image, -i          Image ids seperated by ,
   --endpoint, -e "unix:///var/run/docker.sock" Docker endpoint. example using Boot2Docker tcp://[ip]:[port]
   --match, -m "0"        Number of matching layers
   --boot2docker, -b        Using Boot2Docker?
   --help, -h         show help
   --version, -v        print the version
```

## Release Binaries

To try the latest release of the docker-history-compare command-line interface without installing go, download the statically-linked binary for your architecture from **[Github Releases](https://github.com/devops-israel/docker-compare-history/releases).**

## Credits and License

Developed by **[Josh Dvir](https://github.com/shukydvir/)** for **[DevOps Israel](http://devops.co.il)**

This software is released under the MIT License.