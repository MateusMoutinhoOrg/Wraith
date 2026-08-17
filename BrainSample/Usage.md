
### Brain Workflow:

## Process:

### Show Error 
- creates a Error.md containing informations about the error.


### Stop Execution
- verify if <Project>/Task.yaml exists:
  - if not exists: creates it with default values.
  - edit the <Project>/Task.yaml.aplly  to false.


## Glossary:

## Tick Worflow:
- User runs: 
~~~bash
./brain tick
~~~
- For each project:
  - brain trys to read: <Project>/Task.yaml
  - if the file does not exist:
      - [Stop Execution](#stop-execution)
  - if   <Project>/Task.yaml.name not contains a valida action:
      - [Show a error](#show-error) 
      - [Stop Execution](#stop-execution)

  - if <Project>/Task.yaml.aplly == false:
    - comment: it not shows a error because its not a error 
    - [Stop Execution](#stop-execution) 
 
  - executes  <Project>/Task.yaml.name action.
  - if <Project>/Task.yaml.name action fails:
    - [Show a error](#show-error) 
    - [Stop Execution](#stop-execution) 

'  - renders all markdowns of <Project>/Dasboard


