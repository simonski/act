# todo User Guide

## Install

    go get github.com/simonski/todo

## Setup

    todo init
    > todo.db created

## Create a single task

    todo add "do a thing"

## List all tasks

    todo ls

## Create a task with a headline and description

    todo add "do a thing" "here is the description"

All tasks are part of a project; if unspecified, the project 'default'.

## List more detail on the task

    todo ls t1

## Delete a task

    todo delete t1

## Bring that back

    todo undo

## What have I done?

    todo history

## Create a project

    todo create-project "foo"

## Make this teh default project

    todo set-project "foo"

## Start a task

    todo start t1

## Start a new task

    todo start "I am a new task"

todo will see this is a new task, create and start it.

## What have I been working on

    todo status

## Export my todoivity for the last 3 days

    todo todoivity -days 3

## Export my todoivity since January
    
    todo todoivity -since january

## In json

    todo todoivity -since january -format json

## Run todo as a server

Multiplayer (> 1 person) works wiht todo, your ACT_USERNAME, or whoami will be retained. In this case 
it is easy to spoof another user but that's going to come later, via PKI if ever.

    todo server -p 5989
    > todo is running on 0.0.0.0:5989

## Connect some git reposititories as default backends to monitor for tasks

Note these all only really make sense when running todo in server mode. You can run as client-only.
you run todo git sync

    todo git add -repo git@server/reponame -project 'the name` -description `the description'
    > Added repo

    todo git add -repo git@server/otherrepo 
    > Added repo

    > Tell todo to scan git every 1m
    todo git scan 1m

    todo git ls
    todo git disable git@server/otherepo

    > Tell todo to force-sync to gits
    todo git sync-all


## Connect to todo from another comptuer

    export ACT_URL=server:port
    todo status
    > todoDB: running on 10.2.0.140:5989
    > There are 3 projects, 180 tasks, 5 users.
    export ACT_USERNAME=me_laptop
    
    todo start "foo bar"
    > Created task "foo bar", ID=T104
    
    todo list projects
    > P1 - Billabong bing bang
    > P2 - The zed of fungible
    > P3 - Crypto-yo!

    todo default projet P2
    > Set default/todoive projec to 'P2 - The zed of fungible`
    
    todo reproject T104 
    > Reassigned task T104 to Project `P2 - The zed of fungible`

    todo todoivity
    > You 
    >    T102 `Bloo gar`, Created > Active > Completed
    >    T104 'foo bar', Created
    > 
    > Jane
    >    T101, 'zingo', Created
      


