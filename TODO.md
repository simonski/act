# 0.0.1

rename todo to todo

User Guide
Tools Integration / API integration with codebase &/or editing text files.
http/s server for a simple task management
multiplayer user style via the api an a basic auth layer
server builds from a possibly different repo / module)
	todo-api
	todo		(becomes server)
	todo-rpc

milestone: use the the tool to capture tasks, not notes
	context of location & time and machine is important as part of the note/task
	contxt of the capture - so a user does not say this is a task, they use certain word we can then prune/capture analyse offline to decide this IS a task / todoion / reminder"

db should version all things adn perhaps use a latest version number in the queries when modifying anything?
db could then prune over time anything old-old

users could go to diff db and/or someting else
prepare a basic login/register/signup separate API for posrtgres and sqlite (review java cloud project)
prepare a basic db layer module to allow me to integrate different models (e.g. user model, task model)


# create either $TODO_FILE or $PWD/todo.db
todo init

# create a new task on the current project
todo add foo

# list tasks
task ls

# update a task
task update -id 1 -name "fred" -description "bar bar"

# set a description on foo
todo update -id 1 -descripion "bar bar"

# delete a task
todo rm -id 1

# set the project to a new project
todo set-project fred

# list projects
todo list-projects



todo add "name" "details"
>T1 added.
todo update t1 -name "fred" -description "bar blah"


todo project ls                    lists all projects
todo project 'name'                sets the project to be 'name'
todo project create                creates a new project
todo project update                updates something (name, location0
todo project rm                    deletes a project (marks as deleted)
todo project merge                 overlays one project onto anotehr
todo project archive               archived are invisible to normal
todo project unarchive             unachived come back to life

    walks up until we find a .git or uses current directory as the name of a project

todo 

    look for a .db file
    if find one, maybe make a new project based on current location or root of a .git directory and use that name?
    todo create project -i


- bare bones so it all goes into a todo db
    TESTS
    USER GUIDE
    BARE BONES FUNCTIONALITY
    init        pass/fail
    add         - create a task
    list        - list all tasks
    update      - update name, description of a task
    complete    - mark a task as complete/intodoive/archived
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

todo init
    create database containing schema
        FAIL IF EXISTS

    config
        version: 1.0
        created: Time().String()

# set default project id - all actions will carry out here
todo default-project p1

# set default group, all new tasks will be added here
todo default-group g1

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

