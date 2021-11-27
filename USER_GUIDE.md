# act User Guide

## Install

    go get github.com/simonski/act

## Setup

    act init
    > act.db created

## Create a single task

    act add "do a thing"

## List all tasks

    act ls

## Create a task with a headline and description

    act add "do a thing" "here is the description"

All tasks are part of a project; if unspecified, the project 'default'.

## List more detail on the task

    act ls t1

## Delete a task

    act delete t1

## Create a project

    act create project

## Start a task

    act start t1

## Start a new task

    act start "I am a new task"

act will see this is a new task, create and start it.

## What have I been working on

    act status

## Export my activity for the last 3 days

    act activity -days 3

## Export my activity since January
    
    act activity -since january

## In json

    act activity -since january -format json

## Run act as a server

Multiplayer (> 1 person) works wiht act, your ACT_USERNAME, or whoami will be retained.  in this case 
it is easy to spoof another user but that's going to come later, via PKI if ever.

    act server -p 5989
    > act is running on 0.0.0.0:5989

## Connect some git reposititories as default backends to monitor for tasks

Note these all only really make sense when running taksy in server mode. You can run as client-only.
you run act git sync

    act git add -repo git@server/reponame -project 'the name` -description `the description'
    > Added repo

    act git add -repo git@server/otherrepo 
    > Added repo

    > Tell act to scan git every 1m
    act git scan 1m

    act git ls
    act git disable git@server/otherepo

    > Tell act to force-sync to gits
    act git sync-all


## Connect to act from another comptuer

    export ACT_URL=server:port
    act status
    > actDB: running on 10.2.0.140:5989
    > There are 3 projects, 180 tasks, 5 users.
    export ACT_USERNAME=me_laptop
    
    act start "foo bar"
    > Created task "foo bar", ID=T104
    
    act list projects
    > P1 - Billabong bing bang
    > P2 - The zed of fungible
    > P3 - Crypto-yo!

    act default projet P2
    > Set default/active projec to 'P2 - The zed of fungible`
    
    act reproject T104 
    > Reassigned task T104 to Project `P2 - The zed of fungible`

    act activity
    > You 
    >    T102 `Bloo gar`, Created > Active > Completed
    >    T104 'foo bar', Created
    > 
    > Jane
    >    T101, 'zingo', Created
      


