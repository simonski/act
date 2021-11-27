# Developer Guide

## Building

https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.4.html

To build locally

```bash
make
```

To see the make targets

```bash
make help
```

I use some workflows in github actions (./github/workflows) to compile and test.  

## GitHub/OSS

Use github PRs and vefify using "Compile and Test" actions that the PR is good.

## IMPORTANT

The simplicity is paramount. This means the software should

    1 install easily
    2 run quickly and simply
    3 take the pain away

The objective is to be the best task management tool for developers. This means we need to

	1. work for developers
    2. and teams
    3. be useful
    4. be used
	5. integrate with version control
    
Reality means there are lots of various competing pieces of software.  For now we get the basics right THEN we move to scanning git, multiplayer games, etc.

