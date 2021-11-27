# 0.0.1

# create either $ACT_FILE or $PWD/act.db
act init

# create a new task on the current project
act add foo

# list tasks
task ls

# update a task
task update -id 1 -name "fred" -description "bar bar"

# set a description on foo
act update -id 1 -descripion "bar bar"

# delete a task
act rm -id 1

# set the project to a new project
act set-project fred

# list projects
act list-projects



act add "name" "details"
>T1 added.
act update t1 -name "fred" -description "bar blah"


act project ls                    lists all projects
act project 'name'                sets the project to be 'name'
act project create                creates a new project
act project update                updates something (name, location0
act project rm                    deletes a project (marks as deleted)
act project merge                 overlays one project onto anotehr
act project archive               archived are invisible to normal
act project unarchive             unachived come back to life

    walks up until we find a .git or uses current directory as the name of a project

act 

    look for a .db file
    if find one, maybe make a new project based on current location or root of a .git directory and use that name?
    act create project -i


- bare bones so it all goes into a act db
    TESTS
    USER GUIDE
    BARE BONES FUNCTIONALITY
    init        pass/fail
    add         - create a task
    list        - list all tasks
    update      - update name, description of a task
    complete    - mark a task as complete/inactive/archived
    group       - collect tasks together into a group
    tag         - add  a tag to something


Multiplayer
    think how to merge separate taskdbs via  simple sync call
    find a tree/data structure that describes change efficiently to indicate what is different

task ids should be Tnumber
project ids should be Pnumber
group ids should be Gnumber
tag

task add "name" "description"
task ls

task group add "mygroup"

act init
    create database containing schema
        FAIL IF EXISTS

    config
        version: 1.0
        created: Time().String()

# set default project id - all actions will carry out here
act default-project p1

# set default group, all new tasks will be added here
act default-group g1

    project
        groups
            task1
            task2
            task3
            task4
        group
            task1
            task2
            task3
            task4
        group
            task1
            task2
            task3
            task4
